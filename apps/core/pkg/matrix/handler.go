package matrix

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	assistant_public "github.com/piksar-eu/webapp/apps/core/pkg/assistant/public"
	"github.com/piksar-eu/webapp/apps/core/pkg/command"
	"go.mau.fi/util/dbutil"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/crypto/cryptohelper"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
)

type Matrix struct {
	client     *mautrix.Client
	commandBus *command.CommandBus
}

func NewMatrix(db *sql.DB, commandBus *command.CommandBus) (*Matrix, error) {
	ctx := context.Background()
	userID := id.UserID(os.Getenv("MATRIX_USER"))

	client, err := mautrix.NewClient(
		os.Getenv("MATRIX_SERVER"),
		userID,
		"",
	)
	if err != nil {
		return nil, fmt.Errorf("błąd tworzenia klienta: %v", err)
	}
	cryptoDb, err := dbutil.NewWithDB(db, "postgres")

	cryptoHelper, err := cryptohelper.NewCryptoHelper(client, []byte("meow"), cryptoDb)
	if err != nil {
		return nil, fmt.Errorf("błąd inicjalizacji crypto helpera: %v", err)
	}

	password := os.Getenv("MATRIX_PASSWORD")
	cryptoHelper.LoginAs = &mautrix.ReqLogin{
		Type:       mautrix.AuthTypePassword,
		Identifier: mautrix.UserIdentifier{Type: mautrix.IdentifierTypeUser, User: string(userID)},
		Password:   password,
	}
	if err := cryptoHelper.Init(ctx); err != nil {
		return nil, fmt.Errorf("crypto init error: %v", err)
	}
	client.Crypto = cryptoHelper

	return &Matrix{
		client:     client,
		commandBus: commandBus,
	}, nil
}

func (m *Matrix) sendTextMsg(roomId, message string) error {
	ctx := context.Background()
	rId := id.RoomID(roomId)
	_, err := m.client.SendText(ctx, rId, message)
	if err != nil {
		return fmt.Errorf("can not send matrix message")
	}

	return nil
}

func (m *Matrix) syncMatrix() {
	syncer := m.client.Syncer.(*mautrix.DefaultSyncer)
	syncer.OnEventType(event.StateMember, func(ctx context.Context, evt *event.Event) {
		if evt.GetStateKey() == m.client.UserID.String() && evt.Content.AsMember().Membership == event.MembershipInvite {
			_, err := m.client.JoinRoomByID(ctx, evt.RoomID)
			if err != nil {
				log.Printf("Nie można dołączyć do pokoju %s: %v", evt.RoomID, err)
			} else {
				log.Printf("Dołączono do pokoju: %s (zaproszenie od %s)", evt.RoomID, evt.Sender)
			}
		}
	})

	syncer.OnEventType(event.EventMessage, func(ctx context.Context, evt *event.Event) {
		if evt.Sender == m.client.UserID {
			return // ignoruj siebie
		}

		msg := evt.Content.AsMessage()
		switch msg.MsgType {
		case event.MsgText:
			log.Printf("[tekst] %s: %s", evt.Sender, msg.Body)
			m.processMsg(ctx, evt, msg.Body)

		case event.MsgAudio:
			log.Printf("[audio] %s wysłał audio", evt.Sender)

			if msg.File != nil && msg.File.URL != "" {
				data, _ := m.client.DownloadBytes(ctx, msg.File.URL.ParseOrIgnore())

				reader := bytes.NewReader(data)

				ef := msg.File
				ef.PrepareForDecryption()
				plainReader := ef.DecryptStream(reader)

				res, err := transcript(plainReader)
				if err != nil {
					log.Printf("Nie udało się wykonać transkrypcji audio")
				}

				m.processMsg(ctx, evt, res)

			} else {
				log.Printf("Wiadomość audio od %s nie zawiera URL pliku.", evt.Sender)
			}

		default:
			log.Printf("[inne] Typ wiadomości: %s, od: %s", msg.MsgType, evt.Sender)
		}
	})

	log.Println("Bot działa...")

	if err := m.client.Sync(); err != nil {
		log.Fatalf("Błąd synchronizacji: %v", err)
	}
}

func (m *Matrix) processMsg(ctx context.Context, evt *event.Event, msg string) error {

	ec := evt.Content.Parsed.(*event.MessageEventContent)

	rt := ec.RelatesTo
	if rt == nil {
		rt = &event.RelatesTo{
			Type:    "io.element.thread",
			EventID: evt.ID,
		}
	}

	reaction, err := m.client.SendReaction(ctx, evt.RoomID, rt.EventID, "⏳")
	if err != nil {
		return fmt.Errorf("Błąd przy wysyłaniu reakcji: %v", err)
	}

	res, err := m.commandBus.Dispatch(ctx, assistant_public.NewCallAssistantCommand(msg))

	if err != nil {
		return fmt.Errorf("Błąd przy wywołaniu asystenta: %v", err)
	}
	assistRes := res.(*assistant_public.CallAssistantRes)

	if assistRes.Content != "" {
		replyContent := &event.MessageEventContent{
			Body:      assistRes.Content,
			MsgType:   event.MsgText,
			RelatesTo: rt,
		}

		r, err := m.client.SendMessageEvent(ctx, evt.RoomID, event.EventMessage, replyContent)
		if err != nil {
			log.Printf("Błąd przy wysyłaniu odpowiedzi: %v", err)
		}
		log.Println(r.EventID.String())
	}

	m.client.RedactEvent(ctx, evt.RoomID, reaction.EventID) // This removes the reaction

	return nil
}

func transcript(r io.Reader) (string, error) {
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	go func() {
		part, err := writer.CreateFormFile("audio_file", "audio.ogg")
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		_, err = io.Copy(part, r)
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		err = writer.Close()
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		pw.Close()
	}()

	req, _ := http.NewRequest("POST", "http://localhost:9000/asr?output=json&language=pl&vad_filter=true", pr)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// wykonaj request
	clientHTTP := &http.Client{}
	resp, err := clientHTTP.Do(req)
	if err != nil {
		return "", fmt.Errorf("błąd wykonania requestu do whispera: %v", err)
	}
	defer resp.Body.Close()

	// odczytaj odpowiedź
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("błąd odczytu odpowiedzi: %v", err)
	}

	var response struct {
		Text string `json:"text"`
	}

	err = json.Unmarshal([]byte(string(body)), &response)
	if err != nil {
		return "", fmt.Errorf("błąd parsowania odpowiedzi whisper: %v", err)
	}

	return response.Text, nil
}

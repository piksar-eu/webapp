package assistant

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/hirochachacha/go-smb2"
	"github.com/piksar-eu/webapp/apps/core/pkg/assistant/public"
	"github.com/piksar-eu/webapp/apps/core/pkg/command"
	"github.com/sashabaranov/go-openai"
)

type CallAssistantHandler struct{}

func (h *CallAssistantHandler) Handle(ctx context.Context, cm command.Command) (any, error) {
	c, ok := cm.(public.CallAssistantCommand)
	if !ok {
		return nil, fmt.Errorf("invalid command type")
	}

	msg, err := openaiCall(c)

	if err != nil {
		msg = fmt.Sprintf("⛔️ Błąd podczas komunikacji z asystentem: %v", err)
	}

	return &public.CallAssistantRes{
		Content: msg,
	}, nil
}

func openaiCall(c public.CallAssistantCommand) (string, error) {

	cnf := openai.DefaultConfig(os.Getenv("OPENROUTER_KEY"))
	cnf.BaseURL = "https://openrouter.ai/api/v1"

	client := openai.NewClientWithConfig(cnf)

	now := time.Now()
	date := now.Format("Monday, 2 January 2006 at 15:04")

	messages := []openai.ChatCompletionMessage{
		{
			Role: "system",
			Content: fmt.Sprintf(`
				Jesteś inteligentnym asystentem odpowiedzialnym za rozdzielanie zadań do odpowiednich narzędzi.
				Staraj się wybrać narzędzie na podstawie opisu jego możliwości.
				Wybierz tylko jedno narzędzie.
				Dziś jest %s`, date),
		},
		{
			Role:    "user",
			Content: c.Content,
		},
	}

	tools := []openai.Tool{
		{
			Type: openai.ToolTypeFunction,
			Function: &openai.FunctionDefinition{
				Name:        "lead_monitor",
				Description: `Narzędzie do monitorowania liczby nowych leadów (subskrybentów newslettera).`,
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"since": map[string]any{
							"type":        "string",
							"description": "Data początkowa w formacie ISO 8601 YYYY-MM-DDTHH:MM:SS±HH:MM, strefa czasowa polska.",
						},
					},
					"required": []string{},
				},
			},
		},
		{
			Type: openai.ToolTypeFunction,
			Function: &openai.FunctionDefinition{
				Name:        "create_note",
				Description: "Zapisuje różnego rodzaju informacje w formie notatek. Wymaga wyłącznie podania treści.",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"content": map[string]any{
							"type":        "string",
							"description": "Sformatowana treść notatki. znajdź i popraw literówki, niedomówienia, błędy transkrypcji",
						},
					},
					"required": []string{"content"},
				},
			},
		},
	}

	req := openai.ChatCompletionRequest{
		// Model:    "qwen3:8b",
		Model:    "openai/gpt-3.5-turbo",
		Messages: messages,
		Tools:    tools,
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil || len(resp.Choices) == 0 {
		return "", err
	}

	res := &resp.Choices[0].Message

	re := regexp.MustCompile(`(?s)<think>.*?</think>`)
	resContent := strings.TrimSpace(re.ReplaceAllString(res.Content, ""))

	if resContent != "" {
		return resContent, nil
	}

	for _, tc := range res.ToolCalls {
		switch tc.Function.Name {
		case "lead_monitor":
			return leadMonitorTool(tc.Function.Arguments), nil
		case "create_note":
			return createNoteTool(tc.Function.Arguments), nil
		}
	}

	return "", nil
}

func leadMonitorTool(args string) string {
	type Args struct {
		Since string `json:"since"`
	}
	var a Args
	json.Unmarshal([]byte(args), &a)
	// Zahardkodowana odpowiedz
	return fmt.Sprintf("Od %s pozyskano %d nowych leadów.", a.Since, rand.Intn(100-0+1)+0)
}

func createNoteTool(args string) string {
	type Args struct {
		Content string `json:"content"`
	}
	var a Args
	json.Unmarshal([]byte(args), &a)

	fs := smbShare()

	fileName := time.Now().Format("2006.01") + ".md"
	filePath := fmt.Sprintf("_piotrek/obsidian piotr/journal/%s", fileName)

	// Check if file exists
	_, err := fs.Stat(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("error checking file existence: %w", err).Error()
		}

		// File doesn't exist - create all parent directories first
		dir := filepath.Dir(filePath)
		if dir != "." {
			err = fs.MkdirAll(dir, 0755)
			if err != nil {
				return fmt.Errorf("error creating directories: %w", err).Error()
			}
		}

		// Create the file
		file, err := fs.Create(filePath)
		if err != nil {
			return fmt.Errorf("error creating file: %w", err).Error()
		}
		file.Close()
	}

	// Open file in append mode
	file, err := fs.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err).Error()
	}
	defer file.Close()

	// Add timestamp and content
	fullContent := fmt.Sprintf("\n\n--- %s---\n%s", time.Now().Format("2006-01-02 15:04:05"), a.Content)

	// Write content
	_, err = file.WriteString(fullContent)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err).Error()
	}

	return fmt.Sprintf("Dopisałem Ci notatkę do obsidiana (journal/%s)", fileName)
}

func smbShare() *smb2.Share {

	// Parametry połączenia SMB
	server := os.Getenv("SMB_SERVER")
	share := os.Getenv("SMB_SHARE")
	username := os.Getenv("SMB_USERNAME")
	password := os.Getenv("SMB_PASSWORD")

	// TCP połączenie z serwerem SMB
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("Błąd połączenia:", err)
		return nil
	}
	// defer conn.Close()

	// Konfiguracja klienta SMB
	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     username,
			Password: password,
		},
	}

	// Tworzymy sesję SMB
	s, err := d.Dial(conn)
	if err != nil {
		fmt.Println("Błąd sesji SMB:", err)
		return nil
	}
	// defer s.Logoff()

	// Otwieramy udział (share)
	fs, err := s.Mount(share)
	if err != nil {
		fmt.Println("Błąd montowania udziału:", err)
		return nil
	}
	// defer fs.Umount()

	return fs
}

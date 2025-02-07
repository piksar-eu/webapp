package auth

import (
	"crypto"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/piksar-eu/webapp/apps/core/pkg/shared"
	"github.com/piksar-eu/webapp/apps/core/pkg/web"
	"github.com/posterity/srp"
)

type LoginInitReq struct {
	Email string `json:"email"`
}

type LoginInitRes struct {
	SRP LoginInitResSRP
}

type LoginInitResSRP struct {
	Salt string `json:"salt"`
	B    string `json:"B"`
}

type LoginSRPReq struct {
	A  string `json:"A"`
	M1 string `json:"M1"`
}

type LoginSRPRes struct {
	M2   string        `json:"M2"`
	User *LoginResUser `json:"user"`
}

type LoginResUser struct {
	Email string `json:"email"`
}

var srpParams = &srp.Params{
	Group: srp.RFC5054Group3072,
	Hash:  crypto.SHA256,
	KDF:   srp.RFC5054KDF,
}

type LoginHandler struct {
	userRepo UserRepository
	sessCtx  web.SessionContext
}

func (h *LoginHandler) HandleInit(credentials *LoginInitReq) (*LoginInitRes, error) {

	var (
		salt, verifier []byte
	)

	user, _ := h.userRepo.Get(credentials.Email)

	if user != nil {
		data := user.getAuthMethodData("srp").(SRPData)
		salt, _ = hex.DecodeString(data.Salt)
		verifier, _ = hex.DecodeString(data.Verifier)
	} else {
		verifier, salt = FakeSRP(credentials.Email)
	}

	server, err := srp.NewServer(srpParams, credentials.Email, salt, verifier)
	if err != nil {
		return nil, fmt.Errorf("failed to create srp server")
	}

	srpState, err := server.Save()
	if err != nil {
		return nil, fmt.Errorf("failed to save srp server state")
	}

	h.sessCtx.Add("srpUsername", credentials.Email)
	h.sessCtx.Add("srpState", hex.EncodeToString(srpState))

	return &LoginInitRes{
		SRP: LoginInitResSRP{
			Salt: hex.EncodeToString(salt),
			B:    hex.EncodeToString(server.B()),
		},
	}, nil
}

func (h *LoginHandler) HandleSRP(credentials *LoginSRPReq) (*LoginSRPRes, error) {
	srpState := h.sessCtx.Get("srpState")
	if srpState == nil {
		return nil, fmt.Errorf("failed to read srp server state from session")
	}

	serverState, err := hex.DecodeString(srpState.(string))
	if err != nil {
		return nil, fmt.Errorf("failed to decode srp server state")
	}

	server, err := srp.RestoreServer(srpParams, serverState)
	if err != nil {
		return nil, fmt.Errorf("failed to restore srp server state")
	}

	a, err := hex.DecodeString(credentials.A)
	if err != nil {
		return nil, fmt.Errorf("failed to decode A")
	}

	if err := server.SetA(a); err != nil {
		return nil, fmt.Errorf("failed to handle A")
	}

	M1, err := hex.DecodeString(credentials.M1)
	if err != nil {
		return nil, fmt.Errorf("failed to decode M1")
	}

	ok, err := server.CheckM1(M1)
	if err != nil || !ok {
		return nil, &AuthenticationError{"incorrect ClientProof"}
	}

	M2, err := server.ComputeM2()
	if err != nil {
		return nil, fmt.Errorf("failed to compute M2")
	}

	email := h.sessCtx.Get("srpUsername").(string)
	userData := h.authSuccess(email)

	h.sessCtx.Del("srpState", "srpUsername")

	return &LoginSRPRes{
		M2:   hex.EncodeToString(M2),
		User: userData,
	}, nil
}

func (h *LoginHandler) authSuccess(email string) *LoginResUser {
	h.sessCtx.Add("user", &shared.SessionUser{
		Email:    email,
		LoggedAt: time.Now(),
	})

	return &LoginResUser{
		Email: email,
	}
}

func FakeSRP(username string) ([]byte, []byte) {
	saltHash := sha256.Sum256([]byte(fmt.Sprintf("%s_xxx", username)))

	tp, err := srp.ComputeVerifier(srpParams, username, "p@$$w0rd", saltHash[:])
	if err != nil {
		log.Fatalf("failed to compute verifier: %v", err)
	}

	return tp.Verifier(), saltHash[:]
}

type AuthenticationError struct {
	msg string
}

func (e *AuthenticationError) Error() string {
	return e.msg
}

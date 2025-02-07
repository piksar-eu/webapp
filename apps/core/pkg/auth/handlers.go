package auth

import (
	"fmt"

	"github.com/piksar-eu/webapp/apps/core/pkg/shared"
)

type RegistrationHandler struct {
	userRepo UserRepository
}

func (h *RegistrationHandler) Handle(email string, salt string, verifier string) error {
	email, err := shared.SanitizeEmail(email)
	if err != nil {
		return fmt.Errorf("incorrect email")
	}

	user, err := h.userRepo.Get(email)
	if err != nil {
		return err
	}

	if user != nil {
		return fmt.Errorf("already registered")
	}

	user, err = createUser(email)

	if err != nil {
		return err
	}

	user.addAuthenticationMethod(AuthMethod{
		Method: "srp",
		Data: SRPData{
			Salt:     salt,
			Verifier: verifier,
		},
	})

	err = h.userRepo.Save(user)
	if err != nil {
		return err
	}

	return nil
}

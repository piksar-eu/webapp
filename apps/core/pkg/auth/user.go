package auth

import (
	"time"

	"github.com/piksar-eu/webapp/apps/core/pkg/shared"
)

type UserRepository interface {
	Get(email string) (*User, error)
	Save(*User) error
}

type User struct {
	Email       string
	AuthMethods []AuthMethod
	CreatedAt   time.Time
}

type AuthMethod struct {
	Method string
	Data   interface{}
}

type SRPData struct {
	Salt     string `json:"salt"`
	Verifier string `json:"verifier"`
}

func createUser(email string) (*User, error) {
	email, err := shared.SanitizeEmail(email)

	if err != nil {
		return nil, err
	}

	return &User{
		Email:     email,
		CreatedAt: time.Now(),
	}, nil
}

func (u *User) addAuthenticationMethod(method AuthMethod) {
	for k, m := range u.AuthMethods {
		if m.Method == method.Method {
			u.AuthMethods[k].Data = method.Data
			return
		}
	}

	u.AuthMethods = append(u.AuthMethods, method)
}

func (u *User) getAuthMethodData(method string) interface{} {
	for _, authMethod := range u.AuthMethods {
		if authMethod.Method == method {
			return authMethod.Data
		}
	}
	return nil
}

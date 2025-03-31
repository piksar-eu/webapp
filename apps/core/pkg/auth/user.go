package auth

import (
	"slices"
	"time"

	"github.com/piksar-eu/webapp/apps/core/pkg/shared"
)

type UserRepository interface {
	GetById(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	List() (map[string]*User, error)
	Save(*User) error
	NewId() string
}

type User struct {
	Id          string
	Email       string
	Name        string
	AuthMethods []AuthMethod
	Roles       []string
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

func createUser(id string, email string) (*User, error) {
	email, err := shared.SanitizeEmail(email)

	if err != nil {
		return nil, err
	}

	return &User{
		Id:        id,
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

func (u *User) addRoles(ids ...string) {
	for _, id := range ids {
		if slices.Contains(u.Roles, id) {
			continue
		}

		u.Roles = append(u.Roles, id)
	}
}

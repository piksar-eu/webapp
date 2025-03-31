package auth

import (
	"fmt"
)

type UserRes struct {
	Id    string   `json:"id"`
	Email string   `json:"email"`
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}

type UsersHandler struct {
	userRepo  UserRepository
	authStore AuthorizationStore
}

func (h *UsersHandler) List() ([]*UserRes, error) {
	users, err := h.userRepo.List()
	usersRes := make([]*UserRes, 0)

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		usersRes = append(usersRes, &UserRes{
			Id:    user.Id,
			Email: user.Email,
			Name:  user.Name,
			Roles: user.Roles,
		})
	}

	return usersRes, nil
}

func (h *UsersHandler) Get(id string) (*UserRes, error) {
	user, err := h.userRepo.GetById(id)

	if err != nil {
		return nil, err
	}

	return &UserRes{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
		Roles: user.Roles,
	}, nil
}

func (h *UsersHandler) Update(id string, name string, roles []string) error {
	user, _ := h.userRepo.GetById(id)

	if user == nil {
		return fmt.Errorf("failed to find user")
	}

	user.Name = name
	user.Roles = []string{}

	for _, r := range roles {
		if !h.authStore.RoleExists(r) {
			continue
		}

		user.addRoles(r)
	}

	err := h.userRepo.Save(user)

	if err != nil {
		return err
	}

	return nil
}

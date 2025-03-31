package auth

import "slices"

type RoleRepository interface {
	Get(id string) (*Role, error)
	List() (map[string]*Role, error)
	Save(*Role) error
	Delete(id string) error
	NewId() string
}

type Role struct {
	Id          string
	Name        string
	Permissions []string
}

func createRole(id string, name string) (*Role, error) {
	return &Role{
		Id:   id,
		Name: name,
	}, nil
}

func (r *Role) addPermissions(id ...string) {
	for _, p := range id {
		if slices.Contains(r.Permissions, p) {
			continue
		}

		r.Permissions = append(r.Permissions, p)
	}
}

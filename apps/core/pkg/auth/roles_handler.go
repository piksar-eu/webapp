package auth

import "fmt"

type RolesHandler struct {
	roleRepo  RoleRepository
	authStore AuthorizationStore
}

type RoleRes struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

func (h *RolesHandler) List() ([]*RoleRes, error) {
	roles, err := h.roleRepo.List()
	rolesRes := make([]*RoleRes, 0)

	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		rolesRes = append(rolesRes, &RoleRes{
			Id:          role.Id,
			Name:        role.Name,
			Permissions: role.Permissions,
		})
	}

	return rolesRes, nil
}

func (h *RolesHandler) Get(id string) (*RoleRes, error) {
	role, err := h.roleRepo.Get(id)

	if err != nil {
		return nil, err
	}

	return &RoleRes{
		Id:          role.Id,
		Name:        role.Name,
		Permissions: role.Permissions,
	}, nil
}

func (h *RolesHandler) Create(name string, permissions []string) error {
	role, err := createRole(h.roleRepo.NewId(), name)

	if err != nil {
		return err
	}

	for _, p := range permissions {
		if !h.authStore.PermissionExists(p) {
			continue
		}

		role.addPermissions(p)
	}

	err = h.roleRepo.Save(role)

	if err != nil {
		return err
	}

	return nil
}

func (h *RolesHandler) Update(id string, name string, permissions []string) error {
	role, _ := h.roleRepo.Get(id)

	if role == nil {
		return fmt.Errorf("failed to find role")
	}

	role.Name = name
	role.Permissions = []string{}

	for _, p := range permissions {
		if !h.authStore.PermissionExists(p) {
			continue
		}

		role.addPermissions(p)
	}

	err := h.roleRepo.Save(role)

	if err != nil {
		return err
	}

	return nil
}

func (h *RolesHandler) Delete(id string) error {

	err := h.roleRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

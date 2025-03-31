package auth

type PermissionRes struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PermissionsHandler struct {
	authStore AuthorizationStore
}

func (h *PermissionsHandler) List() []*PermissionRes {
	permissionsRes := make([]*PermissionRes, 0)

	for _, p := range h.authStore.Permissions() {
		permissionsRes = append(permissionsRes, &PermissionRes{
			Id:   p.Id,
			Name: p.Name,
		})
	}

	return permissionsRes
}

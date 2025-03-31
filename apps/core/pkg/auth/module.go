package auth

const PermissionRolesView = "roles_view"
const PermissionRolesEdit = "roles_edit"
const PermissionUsersView = "users_view"
const PermissionUsersEdit = "users_edit"

func LoadModule(userRepo UserRepository, roleRepo RoleRepository, authStore AuthorizationStore) *Module {
	m := &Module{
		userRepo:  userRepo,
		roleRepo:  roleRepo,
		authStore: authStore,
	}

	m.load()

	return m
}

type Module struct {
	userRepo  UserRepository
	roleRepo  RoleRepository
	authStore AuthorizationStore
}

func (m *Module) load() {
	m.authStore.RegisterPermission(Permission{Id: PermissionRolesView, Name: "Can view roles"})
	m.authStore.RegisterPermission(Permission{Id: PermissionRolesEdit, Name: "Can edit roles"})
	m.authStore.RegisterPermission(Permission{Id: PermissionUsersView, Name: "Can view users"})
	m.authStore.RegisterPermission(Permission{Id: PermissionUsersEdit, Name: "Can edit users"})
}

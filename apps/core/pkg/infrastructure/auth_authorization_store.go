package infrastructure

import (
	"slices"
	"sync"
	"time"

	"github.com/piksar-eu/webapp/apps/core/pkg/auth"
)

func NewAuthorizationStore(userRepo auth.UserRepository, roleRepo auth.RoleRepository) auth.AuthorizationStore {
	as := &AuthorizationStoreImpl{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}

	as.ReloadRoles()

	return as
}

type AuthorizationStoreImpl struct {
	userRepo      auth.UserRepository
	roleRepo      auth.RoleRepository
	permissions   []auth.Permission
	appRoles      *sync.Map
	userRoleCache sync.Map
}

func (s *AuthorizationStoreImpl) RegisterPermission(p auth.Permission) {
	s.permissions = append(s.permissions, p)
}

func (s *AuthorizationStoreImpl) Permissions() []auth.Permission {
	return s.permissions
}

func (s *AuthorizationStoreImpl) PermissionExists(p string) bool {
	for _, permission := range s.permissions {
		if permission.Id == p {
			return true
		}
	}

	return false
}

func (s *AuthorizationStoreImpl) RoleExists(r string) bool {
	_, ok := s.appRoles.Load(r)

	return ok
}

func (s *AuthorizationStoreImpl) ReloadRoles() {
	rolesList, _ := s.roleRepo.List()

	newRoles := sync.Map{}

	for _, role := range rolesList {
		newRoles.Store(role.Id, role.Permissions)
	}

	s.appRoles = &newRoles
}

func (s *AuthorizationStoreImpl) userRoles(uId string) []string {
	if entry, ok := s.userRoleCache.Load(uId); ok {
		cached := entry.(userRoleCacheEntry)
		if time.Now().Before(cached.expiration) {
			return cached.roles
		}
		s.userRoleCache.Delete(uId)
	}

	user, _ := s.userRepo.GetById(uId)
	if user == nil {
		return []string{}
	}

	s.userRoleCache.Store(uId, userRoleCacheEntry{
		roles:      user.Roles,
		expiration: time.Now().Add(3 * time.Minute),
	})

	return user.Roles
}

func (s *AuthorizationStoreImpl) computePermissions(roles []string) []string {
	var permissions []string

	if slices.Contains(roles, auth.AdminRoleId) {
		for _, p := range s.permissions {
			if slices.Contains(permissions, p.Id) {
				continue
			}

			permissions = append(permissions, p.Id)
		}

		return permissions
	}

	for _, r := range roles {
		rolePermissions, _ := s.appRoles.Load(r)
		for _, p := range rolePermissions.([]string) {
			if slices.Contains(permissions, p) {
				continue
			}

			permissions = append(permissions, p)
		}
	}

	return permissions
}

func (s *AuthorizationStoreImpl) UserPermissions(uId string) []string {
	if uId == "" {
		return []string{}
	}

	return s.computePermissions(s.userRoles(uId))
}

func (s *AuthorizationStoreImpl) HasPermission(uId string, p string) bool {
	return slices.Contains(s.UserPermissions(uId), p)
}

type userRoleCacheEntry struct {
	roles      []string
	expiration time.Time
}

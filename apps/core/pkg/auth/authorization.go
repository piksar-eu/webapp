package auth

import (
	"net/http"

	"github.com/piksar-eu/webapp/apps/core/pkg/web"
)

const AdminRoleId = "roAdm"

type Permission struct {
	Id   string
	Name string
}

type AuthorizationStore interface {
	RegisterPermission(Permission)
	Permissions() []Permission
	PermissionExists(p string) bool
	RoleExists(r string) bool
	HasPermission(u string, p string) bool
	ReloadRoles()
	UserPermissions(uId string) []string
}

func AuthorizationMiddleware(s AuthorizationStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "OPTIONS" {
				next.ServeHTTP(w, r)
				return
			}

			ctx := web.SessionCtx(r)
			user := web.GetSessionUser(r)
			userId := ""

			if user != nil {
				userId = user.Id
			}

			ctx.Add("guard", &Guard{authStore: s, userId: userId})

			next.ServeHTTP(w, r)
		})
	}
}

type Guard struct {
	authStore AuthorizationStore
	userId    string
}

func (g *Guard) Can(p string) bool {
	return g.authStore.HasPermission(g.userId, p)
}

func Can(r *http.Request, p string) bool {
	guard := web.SessionCtx(r).Get("guard").(*Guard)

	return guard.Can(p)
}

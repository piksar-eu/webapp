package traefik_auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/piksar-eu/webapp/apps/core/pkg/auth"
	"github.com/piksar-eu/webapp/apps/core/pkg/web"
)

const PermissionPanelView = "external_panel_view"
const PermissionTraefikView = "external_traefik_view"
const PermissionMemosView = "external_memos_view"

func LoadModule(authStore auth.AuthorizationStore) *Module {
	m := &Module{
		authStore: authStore,
	}

	m.load()

	return m
}

type Module struct {
	authStore auth.AuthorizationStore
}

func (m *Module) load() {
	m.authStore.RegisterPermission(auth.Permission{Id: PermissionPanelView, Name: "Can access panel.pragmatyczny.dev"})
	m.authStore.RegisterPermission(auth.Permission{Id: PermissionTraefikView, Name: "Can access traefik.pragmatyczny.dev"})
	m.authStore.RegisterPermission(auth.Permission{Id: PermissionMemosView, Name: "Can access memos.pragmatyczny.dev"})
}

func (m *Module) ServeApi(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/traefik-auth", func(w http.ResponseWriter, r *http.Request) {
		host := r.Header.Get("X-Forwarded-Host")
		uri := r.Header.Get("X-Forwarded-Uri")

		if u := web.GetSessionUser(r); u == nil {
			redirect := fmt.Sprintf("%s?redirect=https://%s%s", os.Getenv("VITE_LOGIN_URL"), host, uri)

			http.Redirect(w, r, redirect, http.StatusFound)
			return
		}

		requiredPermission := hostPermission(host)

		if requiredPermission != "" && auth.Can(r, requiredPermission) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(""))
			return
		}

		http.Error(w, "Access denied", http.StatusForbidden)
	})
}

func hostPermission(host string) string {
	switch host {
	case "panel.pragmatyczny.dev":
		return PermissionPanelView
	case "traefik.pragmatyczny.dev":
		return PermissionTraefikView
	case "memos.pragmatyczny.dev":
		return PermissionMemosView
	default:
		return ""
	}
}

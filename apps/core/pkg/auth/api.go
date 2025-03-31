package auth

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/piksar-eu/webapp/apps/core/pkg/shared"
	"github.com/piksar-eu/webapp/apps/core/pkg/web"
)

func (m *Module) ServeApi(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/auth/register", func(w http.ResponseWriter, r *http.Request) {
		if u := web.GetSessionUser(r); u != nil {
			http.Error(w, "User already logged in", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			Email    string `json:"email"`
			Salt     string `json:"salt"`
			Verifier string `json:"verifier"`
		}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		registerHandler := RegistrationHandler{
			userRepo: m.userRepo,
		}

		err = registerHandler.Handle(req.Email, req.Salt, req.Verifier)

		if err != nil {
			http.Error(w, "Registration failed", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	})

	mux.HandleFunc("POST /api/auth/login", func(w http.ResponseWriter, r *http.Request) {
		if u := web.GetSessionUser(r); u != nil {
			http.Error(w, "User already logged in", http.StatusMethodNotAllowed)
			return
		}

		loginHandler := LoginHandler{
			userRepo: m.userRepo,
			sessCtx:  *web.SessionCtx(r),
		}

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		var res interface{}
		if req := shared.StrictUnmarshal[LoginSRPReq](bodyBytes); req != nil {
			res, err = loginHandler.HandleSRP(req)
		} else if req := shared.StrictUnmarshal[LoginInitReq](bodyBytes); req != nil {
			res, err = loginHandler.HandleInit(req)
		} else {
			http.Error(w, "Incorrect request data", http.StatusBadRequest)
			return
		}

		if err != nil {
			if _, ok := err.(*AuthenticationError); ok {
				http.Error(w, "Incorrect authentication data", http.StatusBadRequest)
			} else {
				http.Error(w, "Authentication failed", http.StatusInternalServerError)
			}

			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("/api/auth/logout", func(w http.ResponseWriter, r *http.Request) {
		if u := web.GetSessionUser(r); u != nil {
			sessCtx := web.SessionCtx(r)
			sessCtx.Del("user")
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	})

	mux.HandleFunc("GET /api/auth/roles", func(w http.ResponseWriter, r *http.Request) {
		if allowed := Can(r, PermissionRolesView); !allowed {
			http.Error(w, "Access danied", http.StatusForbidden)
			return
		}

		handler := RolesHandler{roleRepo: m.roleRepo}
		res, err := handler.List()

		if err != nil {
			http.Error(w, "Failed", http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("GET /api/auth/roles/{id}", func(w http.ResponseWriter, r *http.Request) {
		if allowed := Can(r, PermissionRolesView); !allowed {
			http.Error(w, "Access danied", http.StatusForbidden)
			return
		}

		id := r.PathValue("id")

		handler := RolesHandler{roleRepo: m.roleRepo}
		res, err := handler.Get(id)

		if err != nil {
			http.Error(w, "Failed", http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("POST /api/auth/roles", func(w http.ResponseWriter, r *http.Request) {
		if allowed := Can(r, PermissionRolesEdit); !allowed {
			http.Error(w, "Access danied", http.StatusForbidden)
			return
		}

		var req struct {
			Name        string   `json:"name"`
			Permissions []string `json:"permissions"`
		}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		handler := RolesHandler{roleRepo: m.roleRepo, authStore: m.authStore}
		err = handler.Create(req.Name, req.Permissions)

		if err != nil {
			http.Error(w, "Failed", http.StatusInternalServerError)

			return
		}

		defer m.authStore.ReloadRoles()

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	})

	mux.HandleFunc("POST /api/auth/roles/{id}", func(w http.ResponseWriter, r *http.Request) {
		if allowed := Can(r, PermissionRolesEdit); !allowed {
			http.Error(w, "Access danied", http.StatusForbidden)
			return
		}

		id := r.PathValue("id")

		var req struct {
			Name        string   `json:"name"`
			Permissions []string `json:"permissions"`
		}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		handler := RolesHandler{roleRepo: m.roleRepo, authStore: m.authStore}
		err = handler.Update(id, req.Name, req.Permissions)

		if err != nil {
			http.Error(w, "Failed", http.StatusInternalServerError)

			return
		}

		defer m.authStore.ReloadRoles()

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	})

	mux.HandleFunc("DELETE /api/auth/roles/{id}", func(w http.ResponseWriter, r *http.Request) {
		if allowed := Can(r, PermissionRolesEdit); !allowed {
			http.Error(w, "Access danied", http.StatusForbidden)
			return
		}

		id := r.PathValue("id")

		handler := RolesHandler{roleRepo: m.roleRepo, authStore: m.authStore}
		err := handler.Delete(id)

		if err != nil {
			http.Error(w, "Failed", http.StatusInternalServerError)

			return
		}

		defer m.authStore.ReloadRoles()

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	})

	mux.HandleFunc("GET /api/auth/permissions", func(w http.ResponseWriter, r *http.Request) {
		handler := PermissionsHandler{authStore: m.authStore}
		res := handler.List()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("GET /api/auth/users", func(w http.ResponseWriter, r *http.Request) {
		if allowed := Can(r, PermissionUsersView); !allowed {
			http.Error(w, "Access danied", http.StatusForbidden)
			return
		}

		handler := UsersHandler{userRepo: m.userRepo, authStore: m.authStore}
		res, err := handler.List()

		if err != nil {
			http.Error(w, "Failed", http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("GET /api/auth/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		if allowed := Can(r, PermissionUsersView); !allowed {
			http.Error(w, "Access danied", http.StatusForbidden)
			return
		}

		id := r.PathValue("id")

		handler := UsersHandler{userRepo: m.userRepo, authStore: m.authStore}
		res, err := handler.Get(id)

		if err != nil {
			http.Error(w, "Failed", http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("POST /api/auth/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		if allowed := Can(r, PermissionUsersEdit); !allowed {
			http.Error(w, "Access danied", http.StatusForbidden)
			return
		}

		id := r.PathValue("id")

		var req struct {
			Name  string   `json:"name"`
			Roles []string `json:"roles"`
		}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		handler := UsersHandler{userRepo: m.userRepo, authStore: m.authStore}
		err = handler.Update(id, req.Name, req.Roles)

		if err != nil {
			http.Error(w, "Failed", http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	})
}

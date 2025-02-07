package auth

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/piksar-eu/webapp/apps/core/pkg/shared"
	"github.com/piksar-eu/webapp/apps/core/pkg/web"
)

func ServeApi(mux *http.ServeMux, userRepo UserRepository) {
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
			userRepo: userRepo,
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
			userRepo: userRepo,
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
}

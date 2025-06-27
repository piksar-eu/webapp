package easyconnect

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/piksar-eu/webapp/apps/core/pkg/easyconnect/public"
)

func (m *Module) ServeApi(mux *http.ServeMux) {
	type subscribeRequest struct {
		Email string `json:"email"`
	}

	mux.HandleFunc("POST /api/easyconnect/subscribe", func(w http.ResponseWriter, r *http.Request) {
		var req subscribeRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = m.commandBus.Dispatch(context.Background(), public.NewCreateLeadCommand(req.Email, "newsletter", true))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	})
}

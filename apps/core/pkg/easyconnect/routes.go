package easyconnect

import (
	"encoding/json"
	"net/http"
)

func ServeApi(mux *http.ServeMux, leadRepo LeadRepository) {

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

		subscribe := SubscribeFn(leadRepo)
		err = subscribe(req.Email)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	})
}

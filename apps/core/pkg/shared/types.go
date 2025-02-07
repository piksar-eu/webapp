package shared

import "time"

type SessionUser struct {
	Email    string    `json:"email"`
	LoggedAt time.Time `json:"loggedAt"`
}

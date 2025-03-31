package shared

import "time"

type SessionUser struct {
	Id       string    `json:"id"`
	Email    string    `json:"email"`
	LoggedAt time.Time `json:"loggedAt"`
}

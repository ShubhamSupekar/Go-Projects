package user

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type UserFields struct {
	Username  string
	Firstname string
	Lastname  string
	Password  string
}

func AlreadyLoggedIn(r *http.Request) bool {
	_, err := r.Cookie("session")
	if err == nil {
		return true
	}
	return false
}
func Cookie(w http.ResponseWriter) string {
	sID := uuid.NewV4()
	co := &http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}
	http.SetCookie(w, co)
	return co.Value
}

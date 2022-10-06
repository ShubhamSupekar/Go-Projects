package handlers

import (
	"net/http"

	"user.com/config"
	"user.com/user"
)

// logout
func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	if user.AlreadyLoggedIn(r) == false {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	c, _ := r.Cookie("session")
	val := c.Value
	if val == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	_, err := config.Db.Exec("DELETE FROM sessions WHERE uid=$1;", val)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

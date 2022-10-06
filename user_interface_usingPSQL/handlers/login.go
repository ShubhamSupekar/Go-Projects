package handlers

import (
	"net/http"
	"user.com/user"
	"golang.org/x/crypto/bcrypt"
	"user.com/config"
)

// login
func Login(w http.ResponseWriter, r *http.Request) {
	if user.AlreadyLoggedIn(r) == true {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	config.Tpl.ExecuteTemplate(w, "login.html", nil)
}
func LoginAuth(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("username")
	p := r.FormValue("password")
	var hash string
	row := config.Db.QueryRow("SELECT  pass FROM users WHERE username = $1", u)
	err := row.Scan(&hash)
	if err != nil {
		config.Tpl.ExecuteTemplate(w, "login.html", "Username Not Registred")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	if err != nil {
		config.Tpl.ExecuteTemplate(w, "login.html", "Incorrect Password")
		return
	}
	co := user.Cookie(w)
	_, err = config.Db.Exec("INSERT INTO sessions (uname, uid) VALUES ($1, $2)", u, co)
	if err != nil {
		config.Tpl.ExecuteTemplate(w, "login.html", "Internal server error")
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

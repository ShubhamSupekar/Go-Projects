package handlers

import (
	"database/sql"
	"net/http"

	"user.com/config"
	"user.com/user"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	if user.AlreadyLoggedIn(r) == false {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	c, _ := r.Cookie("session")
	// var val string
	var U user.UserFields
	row := config.Db.QueryRow("SELECT  uname FROM sessions WHERE uid = $1", c.Value)
	err := row.Scan(&U.Username)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	// row1 := config.Db.QueryRow("SELECT  fname FROM users WHERE username = $1", val)
	row1 := config.Db.QueryRow("SELECT fname, lname FROM users WHERE username = $1", U.Username)
	err = row1.Scan(&U.Firstname, &U.Lastname)
	if err == sql.ErrNoRows {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	config.Tpl.ExecuteTemplate(w, "index.html", U)
}

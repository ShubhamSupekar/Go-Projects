package handlers

import (
	"database/sql"
	"net/http"
	"unicode"

	"golang.org/x/crypto/bcrypt"
	"user.com/config"
	"user.com/user"
)

// signup
func Signup(w http.ResponseWriter, r *http.Request) {
	config.Tpl.ExecuteTemplate(w, "signup.html", nil)
}

func SignupAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	if user.AlreadyLoggedIn(r) == true {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//get form values
	us := user.UserFields{}
	u := r.FormValue("username")
	//checking username criteria
	var nameLength bool
	if 5 <= len(u) && len(u) <= 50 {
		nameLength = true
	}
	row := config.Db.QueryRow("SELECT * FROM users WHERE username = $1", u)
	error := row.Scan(&u)
	if error != sql.ErrNoRows {
		config.Tpl.ExecuteTemplate(w, "signup.html", "Username has been taken choose another one")
		return
	}

	f := r.FormValue("firstname")
	var nameAlphanumericFirst = true
	var nameAlphanumericLast = true
	for _, char := range f {
		if !unicode.IsLetter(char) == true {
			nameAlphanumericFirst = false
		}
	}
	l := r.FormValue("lastname")
	for _, char := range l {
		if !unicode.IsLetter(char) == true {
			nameAlphanumericLast = false
		}
	}
	p := r.FormValue("password")
	//checking password criteria
	var Lower, Upper, Number, Symbol, Length, Nospace bool
	Nospace = true
	for _, char := range p {
		switch {
		case unicode.IsLower(char):
			Lower = true
		case unicode.IsUpper(char):
			Upper = true
		case unicode.IsNumber(char):
			Number = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			Symbol = true
		case unicode.IsSpace(int32(char)):
			Nospace = false
		}
	}
	if 11 < len(p) && len(p) < 60 {
		Length = true
	}
	if !Length || !Upper || !Lower || !nameAlphanumericFirst || !nameAlphanumericLast || !nameLength || !Symbol || !Nospace || !Number {
		config.Tpl.ExecuteTemplate(w, "signup.html", "Check username and password criteria")
		return
	}
	us.Firstname = f
	us.Lastname = l
	us.Username = u
	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	us.Password = string(bs)
	//validate from values
	if us.Firstname == "" || us.Lastname == "" || us.Username == "" || us.Password == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	co := user.Cookie(w)
	_, err1 := config.Db.Exec("INSERT INTO users (username, fname, lname, pass) VALUES ($1, $2, $3, $4)", us.Username, us.Firstname, us.Lastname, us.Password)
	_, err2 := config.Db.Exec("INSERT INTO sessions (uname, uid) VALUES ($1, $2)", us.Username, co)
	if err1 != nil && err2 != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

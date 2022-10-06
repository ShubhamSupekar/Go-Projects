package main

import (
	"net/http"

	"user.com/handlers"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/loginAuth", handlers.LoginAuth)
	http.HandleFunc("/logout", handlers.Logout)
	http.HandleFunc("/signup", handlers.Signup)
	http.HandleFunc("/signupAuth", handlers.SignupAuth)
	http.ListenAndServe(":8080", nil)
}

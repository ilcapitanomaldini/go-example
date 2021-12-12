package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func main() {

	router.HandleFunc("/", loginPageHandler)
	router.HandleFunc("/userPage", userPageHandler)

	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")
	router.HandleFunc("/forgotPassword", forgotPasswordHandler).Methods("POST")
	router.HandleFunc("/updatePassword", updatePasswordHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":9000", nil)
}

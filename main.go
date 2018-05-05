package main

import (
	"whatsapp/servicehandlers"
	"log"
	"net/http"
)

func main() {

	p := servicehandlers.PingHandler{}
	a := servicehandlers.AuthenticateHandler{}
	s := servicehandlers.SignupHandler{}
	//c := servicehandlers.ChatHandler{}
	http.Handle("/ping", p)
	http.Handle("/whatsapp/signup", s)
	http.Handle("/whatsapp/authenticate", a)
	//http.Handle("/whatsapp/Chat",)

	x := http.ListenAndServe(":8080", nil)
	log.Fatal(x)
}

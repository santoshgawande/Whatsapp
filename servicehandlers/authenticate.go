package servicehandlers

import (
	"whatsapp/dao"
	"encoding/json"
	"net/http"
)

type AuthenticateHandler struct {
}

func (p AuthenticateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

//Authenciate  Method
func (p AuthenticateHandler) Get(r *http.Request) (string, int) {
	return "Invalid Request", 200
}

func (p AuthenticateHandler) Put(r *http.Request) (string, int) {
	return "Invalid Request", 200
}


func (p AuthenticateHandler) Post(r *http.Request) (string, int) {

	var response dao.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&response)

	if err != nil {
		panic(err)
	}

	if !dao.AuthenticateUser(response.Email, response.Password) {
		return "Invalid Email or Password", 400
	}

	return "User Authentication Done", 200
}


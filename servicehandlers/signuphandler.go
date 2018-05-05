package servicehandlers

import (
	"net/http"
	"whatsapp/dao"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

type SignupHandler struct {
}

func (p SignupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p SignupHandler) Get(r *http.Request) (string, int) {
	return "Sign Up GET Called", 200
}

func (p SignupHandler) Put(r *http.Request) (string, int) {
	return "Sign Up PUT Called", 200
}

func (p SignupHandler) Post(r *http.Request) (string, int) {
	//return "Sign Up POST Called", 200

	var signup dao.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&signup)
	if err != nil {
		// return "BAD Request", 400
		panic(err)
	}

	id := bson.NewObjectId()
	success := dao.CreateUser(id, signup.Email, signup.Password)

	if success {
		return id, 200
	}
	return "Internal Server Error", 500


}

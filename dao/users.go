package dao

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User Collection Struct
type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
}

var userCollection = GetUserCollection(session)

// returns all users from users collection
func GetAllUsers() []User {
	var users []User

	err := userCollection.Find(nil).All(&users)
	if err != nil {
		panic(err)
	}

	return users
}

// GetUserByEmail return the user if exists
func GetUserByEmail(email string) User {
	var user User

	err := userCollection.Find(bson.M{"email": email}).One(&user)

	if err == mgo.ErrNotFound {
		return user
	} else if err != nil {
		panic(err)
	}

	return user
}
// CreateUser : Create user in Users Collection
func CreateUser(id bson.ObjectId, email string, password string) bool {
	var user User
	user.Id = id
	user.Email = email
	user.Password = password
	
//get email from User Collections
	u := GetUserByEmail(email)
	//fmt.Println("User Data",u.Email)
	if (u.Email != user.Email){
		fmt.Println("User Created:",id)
	err := userCollection.Insert(&user)
	if err != nil {
		return false
	}
	return true
	}else{
		return false
	
}
}
//Authenticate User 
func AuthenticateUser(email string, password string) bool {
	user := GetUserByEmail(email)
	pass := password

	if pass == user.Password {
		return true
	}
	return false
}

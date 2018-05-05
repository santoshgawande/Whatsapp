package dao
import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ChatDetail struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Sender    string        `bson:"email"`
	Receiver string        `bson:"email"`
	Message  string         `bson:"email"
}


package dao

import mgo "gopkg.in/mgo.v2"

const dbName = "Whatsapp"
const hostname = "127.0.0.1"

const (
	userdb         = "users"
	//chatdetailsdb   ="chatdetails"
)

var session = getDatabaseSession()

// defer closeDatabaseSession(session)

func getDatabaseSession() *mgo.Session {
	s, err := mgo.Dial(hostname)

	if err != nil {
		panic(err)
	}
	return s
}

func closeDatabaseSession(s *mgo.Session) {
	s.Close()
}

// Returns User Collection handler
func GetUserCollection(s *mgo.Session) *mgo.Collection {
	return s.DB(dbName).C(userdb)
}


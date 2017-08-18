package db

import (
	"gopkg.in/mgo.v2"
)

func CreateSession(url string) *mgo.Session {
	session, err := mgo.Dial(url)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}

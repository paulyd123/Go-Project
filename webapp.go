package main

import (
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	username string
	email    string
}

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Get("/test", connect)
	m.Run()
	m.Run()
}

func connect() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// Collection User test
	c := session.DB("College").C("User")

	// Insert
	err = c.Insert(&User{username: "Gareth", email: "test@test.test"})
	if err != nil {
		panic(err)
	}
}

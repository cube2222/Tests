package main

import (
	"log"

	"github.com/dancannon/gorethink"
)

type User struct {
	ID   string `gorethink:"id"`
	Name string `gorethink:"name"`
	Age  int    `gorethink:"age"`
}

func main() {
	session, err := gorethink.Connect(gorethink.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	_, err = gorethink.DB("application").Table("users").Get("xyzbaac").Update(map[string]interface{}{
		"age": 22,
	}).RunWrite(session)
	if err != nil {
		log.Fatal(err)
	}
}

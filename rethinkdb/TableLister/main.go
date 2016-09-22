package main

import (
	"github.com/dancannon/gorethink"
	"log"
	"fmt"
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

	users := []User{}
	cursor, err := gorethink.DB("application").Table("users").Run(session)
	if err != nil {
		log.Fatal(err)
	}

	err = cursor.All(&users)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Printf("ID: %v Name: %v Age: %v\n", user.ID, user.Name, user.Age)
	}
}

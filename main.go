package main

import (
	db "appointerServer/db"
	"log"
)

/*
	TODO
	01 Initialize server
	03 Call correct procedure per req

	TO CHECK
	02 Initialize connection with db and close it
*/

func main() {

	client, ctx, cancel, err := db.ConnectToDB()
	if err != nil {
		log.Panic(err)
	}

	defer db.CloseDBClient(client, ctx, cancel)

}

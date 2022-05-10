package main

import (
	"log"

	"github.com/nyan2d/note/app"
	"github.com/nyan2d/smartbolt"
)

func main() {
	databasePath := "database.db"
	db, err := smartbolt.Open(databasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	a := app.NewApp(db)
	a.Run()
}

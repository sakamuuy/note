package db

import (
	"fmt"
	"log"
)

func AddTag(name string) {
	Open()
	defer Close()

	if !IsInitialized() {
		Initialize()
	}

	tx := BeginTransaction()
	stmt, err := tx.Prepare("insert into tags(name, created_at, updated_at) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	now := GetNowFormattedStr()
	_, err = stmt.Exec(name, now, now)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	fmt.Printf("Create tag \"%v\" 🚀 \n", name)
	return
}

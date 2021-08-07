package db

import (
	"fmt"
	"log"
)

func AddFile(name string) {
	Open()
	defer Close()

	if !IsInitialized() {
		Initialize()
	}

	tx := BeginTransaction()
	stmt, err := tx.Prepare("insert into files(name, created_at, updated_at) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	now := GetNowFormattedStr()
	_, err = stmt.Exec(name, now, now)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	fmt.Printf("Create folder %v 🚀 \n", name)
	return
}

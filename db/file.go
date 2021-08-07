package db

import (
	"fmt"
	"log"
)

func AddFile(name string, folderName string) {
	Open()
	defer Close()

	if !IsInitialized() {
		Initialize()
	}

	tx := BeginTransaction()

	rows, err := db.Query("select id from folders where name=?", folderName)
	if err != nil {
		log.Fatal(err)
	}
	if !rows.Next() {
		log.Fatalf("There's no such folder. %v \n", folderName)
	}

	var folderId int
	err = rows.Scan(&folderId)
	if err != nil {
		log.Fatal(err)
	}
	rows.Close()

	stmt, err := tx.Prepare("insert into files(name, created_at, updated_at, folder_id) values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	now := GetNowFormattedStr()
	fmt.Printf("folderId: %v \n", folderId)
	_, err = stmt.Exec(name, now, now, folderId)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Create file %v 🚀 \n", name)
	return
}

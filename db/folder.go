package db

import (
	"fmt"
	"log"
)

func AddFolder(name string) {
	Open()
	defer Close()

	if !IsInitialized() {
		Initialize()
	}

	tx := BeginTransaction()
	stmt, err := tx.Prepare("insert into folders(name, created_at, updated_at) values(?, ?, ?)")
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

	fmt.Printf("Create folder \"%v\" 🚀 \n", name)
	return
}

func GetFilesFolderHas(folderName string) (fileNames []string) {
	Open()
	defer Close()

	rows, err := db.Query(`
		select 
			files.name 
		from folders 
		inner join files 
			on folders.id = files.folder_id 
		where folders.name=?
	`, folderName)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var fileName string
		rows.Scan(&fileName)
		fileNames = append(fileNames, fileName)
	}

	return
}

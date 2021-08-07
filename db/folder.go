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

func GetFolderByName(name string) string {
	Open()
	defer Close()

	rows, err := db.Query(`SELECT id FROM folders WHERE name=?`, name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var id string
	for rows.Next() {
		rows.Scan(&id)
	}

	return id
}

func PatchNewNameFolder(id string, newName string) {
	Open()
	defer Close()

	tx := BeginTransaction()
	stmt, err := tx.Prepare("update folders set name = ? where id = ?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(newName, id)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()
	fmt.Printf("Updated folder %v", newName)
	return
}

func GetAllFolderName() (folderNames []string) {
	Open()
	defer Close()

	rows, err := db.Query(`select name from folders`)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var folderName string
		rows.Scan(&folderName)
		folderNames = append(folderNames, folderName)
	}

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

func DeleteFolder(folderName string) {
	Open()
	defer Close()

	tx := BeginTransaction()
	stmt, err := tx.Prepare("delete from folders where name=?")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(folderName)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()
}

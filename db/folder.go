package db

import "log"

func AddFolder(name string) {
	Open()
	defer Close()

	if !IsInitialized() {
		Initialize()
	}

	tx := BeginTransaction()
	stmt, err := tx.Prepare("insert into folders(name) values(?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(name)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	return
}

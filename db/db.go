package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const dbName = "./db/noteapp.sql"

func Open() {
	tmp_db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	db = tmp_db
}

func Close() {
	db.Close()
}

/*
Return true if local db is *initialized.
initialized is db has metatable and is_initialized is true
*/
func IsInitialized() bool {
	metaRows, err := db.Query("select is_initialized from meta")
	if err != nil {
		fmt.Println("Run `noteapp init` before run any commands")
		log.Fatal(err)
	}
	defer metaRows.Close()

	if !metaRows.Next() {
		return false
	}

	var isInitialized string
	err = metaRows.Scan(&isInitialized)

	if err != nil {
		log.Fatal(err)
	}

	return isInitialized == "1"
}

func Initialize() {
	tx := BeginTransaction()
	var err error

	statement := `
		create table folders (id integer primary key autoincrement, name text, created_at text, updated_at text);
		create table files (id integer primary key autoincrement, name text, created_at text, updated_at text);
		create table tags (id integer primary key autoincrement, name text, created_at text, updated_at text);
		create table meta (is_initialized integer);
		insert into meta(is_initialized) values(1);
	`
	_, err = db.Exec(statement)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()

	fmt.Printf("Initialized✨\n")
}

func BeginTransaction() *sql.Tx {
	transaction, err := db.Begin()
	if err != nil || transaction == nil {
		log.Fatal(err)
	}

	return transaction
}

// func _sinmple() {
// 	sqlStmt := `
// 	create table foo (id integer not null primary key, name text);
// 	delete from foo
// 	`

// 	_, err = db.Exec(sqlStmt)
// 	if err != nil {
// 		log.Printf("%q: %s\n", err, sqlStmt)
// 		return
// 	}

// 	transaction, err := db.Begin()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	statement, err := transaction.Prepare("insert into foo(id, name) values(?, ?)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer statement.Close()

// 	for i := 0; i < 100; i++ {
// 		_, err = statement.Exec(i, fmt.Sprintf("Hello world!%03d", i))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	transaction.Commit()

// 	rows, err := db.Query("select id, name from foo")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() { // Next()調べたらboolean返す、goにはwhileがなくて for bool {}でかける
// 		var id int
// 		var name string
// 		err = rows.Scan(&id, &name)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(id, name)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	statement, err = db.Prepare("select name from foo where id =?")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer statement.Close()
// 	var name string
// 	err = statement.QueryRow("3").Scan(&name)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(name)

// 	_, err = db.Exec("delete from foo")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	rows, err = db.Query("select id, name from foo")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var id int
// 		var name string
// 		err = rows.Scan(&id, &name)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

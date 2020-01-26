package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

func main() {
	db, err := sql.Open("sqlite3", "./db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("create table this_database (id int primary key, value text);")
	if err != nil {
		if !strings.Contains(err.Error(), "this_database already exists") {
			log.Fatal(err)
		}
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	state, err := tx.Prepare("insert into this_database (id, value) values (?, ?);")
	if err != nil {
		log.Fatal(err)
	}

	count := 100
	for i := 0; i < count; i++ { // データ保存
		_, err := state.Exec(i, fmt.Sprintf("データ:%d", i))
		if err != nil {
			continue
		}
	}
	state.Close()

	state, err = db.Prepare("select value from this_database where id = ?;")
	if err != nil {
		log.Fatal(err)
	}

	var data string
	for i := 0; i < count; i++ {
		if err := state.QueryRow(i).Scan(&data); err != nil {
			continue
		}

		fmt.Printf("%d: %+v\n", i, data)
	}
	state.Close()

	tx.Commit()

	/*
		# Outoput:
		92: データ:92
		93: データ:93
		94: データ:94
		95: データ:95
		96: データ:96
		97: データ
	*/
}

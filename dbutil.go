package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "os"
)

var db *sql.DB

func Setup() {
    if !Exists("data.db") {
      	log.Println("Table not found - creating it.")
        defer CreateTable()
    }
    log.Println("Opening connection to the database.")
    database, err := sql.Open("sqlite3", "data.db")
    if err != nil {
        log.Fatal(err)
    }
    db = database
}

func Exists(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }
    if os.IsNotExist(err) {
        return false
    } else {
        return true
    }
}

func CreateTable() {
    log.Println("Creating new table.")
    _, err := db.Exec("CREATE TABLE webpages (id INT, url VARCHAR(20));")
    if err != nil {
        log.Fatal(err)
    }
    _, err = db.Exec("CREATE TABLE webLinks (id INT, url VARCHAR(20));")
    if err != nil {
        log.Fatal(err)
    }
}

func Execute(stmt string) {
    _, err := db.Exec(stmt)
    if err != nil {
        log.Fatal(err)
    }
}

func Database() *sql.DB {
    return db
}

func Close() {
    if db != nil {
        db.Close()
    }
}

package main

import(
     //"github.com/jinzhu/gorm/dialects/postgres"
      //"github.com/jinzhu/gorm"
      _ "github.com/lib/pq"
      "database/sql"
      "log"
      "fmt"
)

const (
    DB_HOST     = "localhost:5436"
    DB_USER     = "postgres"
    DB_PASSWORD = "Roman"
    DB_NAME     = "postgres"
)


//DB_Migration Миграция БД
func DB_Migration() {
    dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
        DB_HOST,DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)

    var lastInsertId int
    err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
    checkErr(err)
    fmt.Println("last inserted id =", lastInsertId)

    defer db.Close()
}


func checkErr(err error) {
    if err != nil {
        log.Println(err)
    }
}
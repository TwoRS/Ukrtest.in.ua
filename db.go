package main

import(
    	//"github.com/jinzhu/gorm/dialects/postgres"
      "github.com/jinzhu/gorm"
      _ "github.com/lib/pq"
      "log"
)

//DbConnect Подключение к базе
func DbConnect() *gorm.DB{
    db, err := gorm.Open("postgres", "user=postgres password=RomanRakzin dbname=postgres sslmode=disable")  
    if err != nil {
       log.Println("Ошибка подключения к базе данных") 
    }
    db.LogMode(false)
    err = db.DB().Ping()
    if err != nil {log.Println("База данных отсутствует или нет доступа") }
    defer db.Close()
    return db
}

//DB_Migration Миграция БД
func DB_Migration() {
    db:=DbConnect()
    defer db.Close()
    //db.AutoMigrate(USER{},FOLDER{}/*,ZABOLEVANIE{},SYMPTOM{},ISSLEDOVANIE{},ISSLEDOVANIE{}*/)
    //db.Model(&FOLDER{}).AddForeignKey("user_id", "users (id)", "CASCADE", "CASCADE")  
}




//  defer db.Close()
 //   db.LogMode(true)
 //   db.AutoMigrate(X{}, Y{})
  //  db.Model(&Y{}).AddForeignKey("x_id", "xes (id)", "CASCADE", "CASCADE")


type X struct {
    gorm.Model
}
 
//Y belongs_to X
type Y struct {
    gorm.Model
    XID uint
}






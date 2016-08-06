package main

import(
    "log"
)

func main() {
    User := UserType{1,"1","1","Roman Rakzin"}
    UsersList[1] = User 
    log.Println(UsersList)
   // DbConnect() 
    DB_Migration()
    RouterStart()
    


    
}


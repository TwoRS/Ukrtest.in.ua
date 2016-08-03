package main
 
import(
    	"net/http"
    	"fmt"
        //"github.com/gorilla/mux"
)

//Authorize Авторизация
func Authorize(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Yes")
        var Login= r.PostFormValue("login")
        var Password= r.PostFormValue("password")
        fmt.Fprintln(w, Login+Password) 
          /*for _, UserInList:= range UsersList {
             if (UserInList.Login== Login && UserInList.Password== Password) {
                Make_Session(w,r,UserInList)                             
             }
          } */
} 

//Authorize Авторизация
func Ping(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "1") 
} 

//Make_Session Создать сессию
func Make_Session(){
   
}

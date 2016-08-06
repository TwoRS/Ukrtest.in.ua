package main
 
import(
    	"net/http"
    	"fmt"
        "time"
        "crypto/sha1"
        //"log"
        //"github.com/gorilla/mux"
)

//Authorize Авторизация
func Authorize(w http.ResponseWriter, r *http.Request) {
        var Login= r.PostFormValue("login")
        var Password= r.PostFormValue("password")
          for _, UserInList:= range UsersList {
             if (UserInList.Login== Login && UserInList.Password== Password) {
                Make_Session(w,r,UserInList)                             
             }
          }
} 

//Authorize Авторизация
func Ping(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "1") 
} 

//Logout Выход
func Logout(w http.ResponseWriter, r *http.Request) {
    var sess Session
    sess=getSession(r)    
    for _, session:= range SessionList {
        if session.UserId==sess.UserId {
            delete(SessionList,session.SessionId) 
        }
    } 
    fmt.Fprintln(w, "<script>var date = new Date(0);document.cookie = 'SessionId=; path=/; expires=' + date.toUTCString();window.location.replace('/login');</script>")
} 

//Make_Session Создать сессию
func Make_Session(w http.ResponseWriter, r *http.Request,User UserType){
  t := time.Now()
  hash_string := fmt.Sprintf("hash %s %s", t,User.Login)
  h := sha1.New()
  h.Write([]byte(hash_string))
  bs := h.Sum(nil)      
  SessionId:=fmt.Sprintf("%x", bs)

   var time string
   time= fmt.Sprintf("%d.%02d.%02d %02d:%02d:%02d",t.Year(),t.Month(),t.Day(),t.Hour(), t.Minute() ,t.Second())
   Create := string(time)

  sess := Session{SessionId, User.Id,User.FIO,User.Login,User.Password,"0",Create}
  SessionList[SessionId] = sess
  fmt.Fprintln(w, SessionId)

  //query:=fmt.Sprintf("insert into session (Session,UserId,Fio,BeginDate) values('%s',%d,'%s','%s')  ",SessionId,User.Id,User.FIO,time)
  //if new_query_exec(query){}else{} 
}

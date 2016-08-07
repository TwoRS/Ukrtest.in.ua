package main
 
import(
     "sync"
     "github.com/gorilla/websocket"
     "time"
     "net/http"
)

type SessionType struct{
  SessionId string
  UserId UserType
  BeginTime string //time !!!!!!!!
}

type UserType struct {
    Id  int 
    Login string  
    Password string       
    FIO string 
    //RegistrationDate string
    //Role string
    //RoditelId int
}

type ClientConnType struct {
    websocket *websocket.Conn
    Socket_Id uint64
}

type NumberConnectionType struct {
    counters uint64
    sync.Mutex
}

type Session struct{
  SessionId string
  UserId int
  UserName string
  UserLogin string
  UserPassword string
  UserRole string
  Created string
}

type Cookie struct {
    Name       string
    Value      string
    Path       string
    Domain     string
    Expires    time.Time
    RawExpires string
    MaxAge   int
    Secure   bool
    HttpOnly bool
    Raw      string
    Unparsed []string // Raw text of unparsed attribute-value pairs
}

type SecureAuthorize struct {
    http.Handler
} 

//----------------- УДАЛИТЬ ПОМЛЕ ИСПЫТАНИЯ
type USER struct {
    ID uint `gorm:"primary_key"`
    Name string
    Login string
    Password string
    Folders []FOLDER
}

type FOLDER struct{
    ID uint `gorm:"primary_key"`
    UserId uint
    HashId uint
    Name string
}


//getSession Functions
/*func getSession(r *http.Request) Session{
    cookie, err := r.Cookie("SessionId")
    if err != nil {
      session:=Session{}
      return session
    }else{
        //log.Println("Кука считана", cookie.Value)
        session:=SessionList[cookie.Value]
        return session
    }
}*/
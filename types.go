package main
 
import(
     "sync"
     "github.com/gorilla/websocket"
)

type SessionType struct{
  SessionId string
  UserId UserType
}

type UserType struct {
    Id  int 
    Login string  
    Password string       
    FIO string 
    RegistrationDate string
    Role string
    RoditelId int
}

type ClientConnType struct {
    websocket *websocket.Conn
    Socket_Id uint64
}

type NumberConnectionType struct {
    counters uint64
    sync.Mutex
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
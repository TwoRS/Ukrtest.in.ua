package main

import(
    	"net/http"
    	"github.com/gorilla/mux"
        "log"
        //"fmt"
)

 

//RouterStart роутер
func RouterStart() {
    var router = mux.NewRouter() 
    static := http.StripPrefix("/static/", http.FileServer(http.Dir("./Files/Static/")))
    router.PathPrefix("/static/").Handler(static).Methods("GET")
   
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {MainHandler(w, r,"Main")}).Methods("GET")
    router.HandleFunc("/forum", func(w http.ResponseWriter, r *http.Request) {MainHandler(w, r,"Main")}).Methods("GET")
    router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {MainHandler(w, r,"Login")}).Methods("GET")
    router.HandleFunc("/api/{method}", func(w http.ResponseWriter, r *http.Request) {ApiHandler(w, r,"Group1")})
    router.HandleFunc("/websocket", ClientSockServer) 

    log.Println("Server started") 
    http.ListenAndServe(":80", router)
}

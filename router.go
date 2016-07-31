package main

import(
    	"net/http"
    	"github.com/gorilla/mux"
        "log"
)

//RouterStart роутер
func RouterStart() {
    var router = mux.NewRouter() 
    static := http.StripPrefix("/static/", http.FileServer(http.Dir("./files/")))
    router.PathPrefix("/static/").Handler(static)
   
    router.HandleFunc("/", IndexHandler)
    router.HandleFunc("/chat", IndexHandler)
    router.HandleFunc("/forum", IndexHandler)

    log.Println("Server started") 
    http.ListenAndServe(":80", router)
}

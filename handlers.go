package main
 
import(
    	"net/http"
    	"fmt"
        "bytes"
        "html/template"
	    "github.com/tdewolff/minify"
	    "github.com/tdewolff/minify/css"
	    "github.com/tdewolff/minify/html"
	    "github.com/tdewolff/minify/js"
        "github.com/gorilla/mux"
)


//MainHandler Хандлер для загрузки шаблонов
func MainHandler(w http.ResponseWriter, r *http.Request,Path string) {
    w.Header().Set("Server", "Go Web Server by Rakzin Roman")
    w.Header().Set("Content-type", "text/html; charset=utf-8")
    templ  := template.New("templ")
    templ .Delims("<%", "%>")
    var templates = template.Must(templ.ParseGlob("Files/Templates/"+Path+"/*/*"))  
    var doc bytes.Buffer 
    var docWrite bytes.Buffer
    m := minify.New()
    m.AddFunc("text/css", css.Minify)
    m.AddFunc("text/html", html.Minify)
    m.AddFunc("text/javascript", js.Minify)
    err := templates.ExecuteTemplate(&doc, "IndexPage", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }  
    if err := m.Minify("text/html", &docWrite, &doc); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    //FilesCache.Set("IndexPage", &docWrite, cache.DefaultExpiration)
    fmt.Fprintln(w, &docWrite)
}

//ApiHandler Хандлер для API
func ApiHandler(w http.ResponseWriter, r *http.Request,Param string) {
    w.Header().Set("Server", "Go Web Server by Rakzin Roman")
    w.Header().Set("Content-type", "text/html; charset=utf-8")
    vars:=mux.Vars(r)
    method:=vars["method"]
    switch method {
        case "authorize":
            Authorize(w,r)
        case "ping":
            Ping(w,r)
        case "Logout":
            Logout(w,r)

        default: 
            fmt.Fprintln(w, "Метод не распознан") 
    }
}

//
func StartHandler(w http.ResponseWriter, r *http.Request,Handler string) {
    w.Header().Set("Server", "Go Web Server by Rakzin Roman")
    w.Header().Set("Content-type", "text/html; charset=utf-8")
    switch Handler {
        case "Logout":
            Logout(w,r)

        default: 
            fmt.Fprintln(w, "Ошибка. Хандлер не найден") 
    }
}


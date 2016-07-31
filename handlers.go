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
)


//IndexHandler Основной хандлер
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Server", "Go Web Server by Rakzin Roman")
    w.Header().Set("Content-type", "text/html; charset=utf-8")
    templ  := template.New("templ")
    templ .Delims("<%", "%>")
    var templates = template.Must(templ.ParseGlob("Templates/*/*/*"))  
    var doc bytes.Buffer 
    var docWrite bytes.Buffer
    m := minify.New()
    m.AddFunc("text/css", css.Minify)
    m.AddFunc("text/html", html.Minify)
    m.AddFunc("text/javascript", js.Minify)
    err := templates.ExecuteTemplate(&doc, "indexPage", nil)
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
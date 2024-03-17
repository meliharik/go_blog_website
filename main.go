package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"net/http"
	"os"
)

func main() {
	r := httprouter.New()
	r.GET("/", Anasayfa)
	r.POST("/deneme", Deneme)
	http.ListenAndServe(":8080", r)
}

func Anasayfa(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, _ := template.ParseFiles("index.html")
	view.Execute(w, nil)
}

func Deneme(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := r.FormValue("username")
	selected := r.FormValue("select_box")
	check := r.FormValue("checkbox")

	r.ParseMultipartForm(10 << 20)
	file, header, _ := r.FormFile("file")
	f, _ := os.OpenFile(header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	io.Copy(f, file)
	fmt.Println("Username: ", username)
	fmt.Println("Select Box: ", selected)
	fmt.Println("Checkbox: ", check)
}

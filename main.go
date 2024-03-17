package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Anasayfa)
	http.HandleFunc("/detay", Detay)
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Anasayfa(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("index.html", "navbar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title": "Anasayfa",
		"Name":  "Golang",
	}
	err = view.ExecuteTemplate(w, "anasayfa", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Detay(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("detay.html", "navbar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title": "Detay",
		"Name":  "Golang",
	}
	err = view.ExecuteTemplate(w, "detay", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

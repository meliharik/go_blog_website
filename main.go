package main

import (
	"goweb/config"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", config.Routes())
}

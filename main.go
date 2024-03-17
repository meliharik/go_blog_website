package main

import (
	admin_models "goweb/admin/models"
	"goweb/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()
	http.ListenAndServe(":8080", config.Routes())
}

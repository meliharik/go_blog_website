package main

import (
	"fmt"
	admin_models "goweb/admin/models"
	"goweb/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()
	post := admin_models.Post{}.Get(1)
	fmt.Println(post)

	post.Delete(1)
	fmt.Println(post)
	http.ListenAndServe(":8080", config.Routes())
}

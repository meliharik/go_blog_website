package config

import (
	"github.com/julienschmidt/httprouter"
	admin "goweb/admin/controllers"
	"net/http"
)

func Routes() *httprouter.Router {
	router := httprouter.New()
	// ADMIN
	router.GET("/admin", admin.Dashboard{}.Index)
	router.GET("/admin/yeni-ekle", admin.Dashboard{}.NewItem)

	// Serve static files
	router.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	return router
}

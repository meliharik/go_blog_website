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
	router.POST("/admin/add", admin.Dashboard{}.Add)
	router.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	router.GET("/admin/edit/:id", admin.Dashboard{}.Edit)

	// Serve static files
	router.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	router.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	return router
}

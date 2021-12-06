package routes

import (
	"github.com/tonnytg/viewerlight/pkg/web/controllers"
	"net/http"
)

// CallRoutes is the main function that handles all the routes
func CallRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/edit", controllers.Edit)

	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)
}

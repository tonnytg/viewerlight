package routes

import (
	"github.com/tonnytg/viewerlight/entity/products"
	"html/template"
	"net/http"
)

// tmpl is a map of all the templates used in the application.
var tmpl = template.Must(template.ParseGlob("templates/*.html"))

// CallRoutes is the main function that handles all the routes
func CallRoutes() {
	http.HandleFunc("/", Index)
}

// Index is the main page
func Index(w http.ResponseWriter, r *http.Request) {
	product := products.Product{ID: 1, Name: "test", Description: "test", Price: 1}
	p, err := products.Create(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
	tmpl.ExecuteTemplate(w, "index.html", p)
}

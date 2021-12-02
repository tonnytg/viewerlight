package routes

import (
	"github.com/tonnytg/viewerlight/entity/products"
	"html/template"
	"log"
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

	product1 := products.Product{ID: 1, Name: "test1", Description: "test1", Price: 1, Actions: []string{"Buy", "Sell"}}
	product2 := products.Product{ID: 2, Name: "test2", Description: "test2", Price: 2, Actions: []string{"Buy", "Sell"}}
	product3 := products.Product{ID: 3, Name: "test3", Description: "test3", Price: 3, Actions: []string{"Buy", "Sell"}}

	sliceProducts := []products.Product{product1, product2, product3}

	err := tmpl.ExecuteTemplate(w, "index.html", sliceProducts)
	if err != nil {
		log.Printf("Error executing t: %v", err)
        return
    }
}

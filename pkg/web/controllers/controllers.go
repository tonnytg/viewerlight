package controllers

import (
	"github.com/tonnytg/viewerlight/entity/entity"
	"github.com/tonnytg/viewerlight/pkg/database"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// tmpl is a map of all the templates used in the application.
var tmpl = template.Must(template.ParseGlob("pkg/web/templates/*.html"))

//var SliceProducts []entity.Product
//
//// TODO: migrate this to a database.
//func init() {
//	product1 := entity.Product{
//		ID:          1,
//		Name:        "Product 1",
//		Price:       10.00,
//		Description: "This is a product",
//		Actions:     []string{"buy", "sell"},
//	}
//	product2 := entity.Product{
//		ID:          2,
//		Name:        "Product 2",
//		Price:       20.00,
//		Description: "This is a product",
//		Actions:     []string{"buy", "sell"},
//	}
//	SliceProducts = append(SliceProducts, product1, product2)
//}

// Index is the main page
func Index(w http.ResponseWriter, r *http.Request) {

	products := database.GetProducts()

	err := tmpl.ExecuteTemplate(w, "index.html", products)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		return
	}
}

// New is form to new product
func New(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "New", nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		return
	}
}

// Insert get information from new page and insert
func Insert(w http.ResponseWriter, r *http.Request) {

	product := entity.Product{}

	if r.Method == "POST" {
		product.ID, _ = strconv.ParseInt(r.FormValue("id"), 10, 64)
		product.Name = r.FormValue("name")
		product.Description = r.FormValue("description")
		product.Price, _ = strconv.ParseFloat(r.FormValue("price"), 64)

		log.Println("Values:", product)
		database.SaveProduct(product)

	}
	http.Redirect(w, r, "/", 301)
}

//// CreateNewProduct save in slice
//func CreateNewProduct(dv ,id int64, name string, price float64, description string) {
//
//	products := database.DataValues{}
//
//	p := entity.Product{
//		ID:          id,
//		Name:        name,
//		Price:       price,
//		Description: description,
//		Actions:     []string{"save", "drop"},
//	}
//	products = append(products, p)
//}

// Edit New is form to new product
func Edit(w http.ResponseWriter, r *http.Request) {

	products := database.GetProducts()

	id := r.FormValue("id")
	var indice int
	idInt, _ := strconv.ParseInt(id, 10, 64)
	for i, v := range products {
		if v.ID == idInt {
			indice = i
			products[i] = entity.Product{
				ID:          v.ID,
				Name:        v.Name,
				Price:       v.Price,
				Description: v.Description,
				Actions:     []string{"save", "drop"},
			}
			break
		}
	}
	err := tmpl.ExecuteTemplate(w, "Edit", products[indice])
	if err != nil {
		log.Printf("Error executing template: %v", err)
		return
	}
}

// Update get information from edit page and update
func Update(w http.ResponseWriter, r *http.Request) {

	product := entity.Product{}

	if r.Method == "POST" {
		product.ID, _ = strconv.ParseInt(r.FormValue("id"), 10, 64)
		product.Name = r.FormValue("name")
		product.Description = r.FormValue("description")
		product.Price, _ = strconv.ParseFloat(r.FormValue("price"), 64)

		database.UpdateProduct(product)
	}
	http.Redirect(w, r, "/", 301)
}

// Delete remove a product
func Delete(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	database.DeleteProduct(id)

	http.Redirect(w, r, "/", http.StatusFound)
}

//// DeleteProductSlice create a func remove Product when the button is clicked
//// remove the product from the slice
//// redirect to index
//func DeleteProductSlice(id string) {
//
//	products := database.GetProducts()
//
//	for i, v := range products {
//		// convert srtring to int64
//		idInt, _ := strconv.ParseInt(id, 10, 64)
//		if v.ID == idInt {
//			log.Println("will be deleted", idInt)
//			products = append(products[:i], products[i+1:]...)
//		}
//	}
//}

package controllers

import (
	"fmt"
	"github.com/tonnytg/viewerlight/entity/products"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// tmpl is a map of all the templates used in the application.
var tmpl = template.Must(template.ParseGlob("pkg/web/templates/*.html"))

var SliceProducts []products.Product

func init() {
	product1 := products.Product{
		ID:          1,
		Name:        "Product 1",
		Price:       10.00,
		Description: "This is a product",
		Actions: []string{"buy", "sell"},
	}
	product2 := products.Product{
		ID:          2,
		Name:        "Product 2",
		Price:       20.00,
		Description: "This is a product",
		Actions: []string{"buy", "sell"},
	}
	SliceProducts = append(SliceProducts, product1, product2)
}

// Index is the main page
func Index(w http.ResponseWriter, r *http.Request) {

	log.Println("Index access")

	err := tmpl.ExecuteTemplate(w, "index.html", SliceProducts)
	if err != nil {
		log.Printf("Error executing t: %v", err)
		return
	}
}

// New is form to new product
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")

		priceConvertToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		idInt, _ := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Println("Error ID convert:", err)
		}

		CreateNewProduct(idInt, name, priceConvertToFloat, description)

		log.Println("Dados:", name, description, price, priceConvertToFloat)
	}
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")

		priceConvertToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error PRICE convert:", err)
		}

		idInt, _ := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Println("Error ID convert:", err)
		}
		for i, v := range SliceProducts {
			if v.ID == idInt {
				SliceProducts[i] = products.Product{
					ID:          idInt,
					Name:        name,
					Price:       priceConvertToFloat,
					Description: description,
					Actions:     []string{"save", "drop"},
				}
				break
			}
		}

		log.Println("Dados atualizados:", name, description, price, priceConvertToFloat)
	}
	http.Redirect(w, r, "/", 301)
}

// CreateNewProduct save
func CreateNewProduct(id int64, nome string, preco float64, descricao string) {

	p := products.Product{
		ID:          id,
		Name:        nome,
		Price:       preco,
		Description: descricao,
		Actions:     []string{"save", "drop"},
	}
	SliceProducts = append(SliceProducts, p)
}

// Edit New is form to new product
func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	var indice int
	idInt, _ := strconv.ParseInt(id, 10, 64)
	for i, v := range SliceProducts {
        if v.ID == idInt {
			indice = i
			SliceProducts[i] = products.Product{
				ID:          v.ID,
                Name:        v.Name,
                Price:       v.Price,
                Description: v.Description,
                Actions:     []string{"save", "drop"},
			}
			break
        }
    }
	tmpl.ExecuteTemplate(w, "Edit", SliceProducts[indice])
}

// Delete remove a product
func Delete(w http.ResponseWriter, r *http.Request) {

	log.Println("delete access")

	id := r.URL.Query().Get("id")
	if id == "" {
		fmt.Fprint(w, "No id provided")
		return
	}
	DeleteProductSlice(id)

	http.Redirect(w, r, "/", http.StatusFound)
}

// DeleteProductSlice create a func remove Product when the button is clicked
// remove the product from the slice
// redirect to index
func DeleteProductSlice(id string) {
	for i, v := range SliceProducts {
		// convert srtring to int64
		idInt, _ := strconv.ParseInt(id, 10, 64)
		if v.ID == idInt {
			SliceProducts = append(SliceProducts[:i], SliceProducts[i+1:]...)
		}
	}
}

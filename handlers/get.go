package handlers

import (
	"net/http"
	"microservices/data"
)

//swagger: route GET /products products listProducts
// Returns a list of products

// getProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("HANDLE GET PRODUCTS")

	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
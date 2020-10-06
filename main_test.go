package main

import (
	"coffe-api/sdk/client"
	"coffe-api/sdk/client/products"
	"fmt"
	"testing"
)

func TestOutClient(t *testing.T) {
	c := client.Default
	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", prod.GetPayload()[1])
	t.Fail()
}

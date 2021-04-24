package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/getProducts", getProducts)

	http.ListenAndServe(":8080", nil)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	productId := r.FormValue("productId")
	products := readData()
	var response []Product

	if productId != "" {
		for _, product := range products {
			searchKey, _ := strconv.Atoi(productId)
			if product.Id == searchKey {
				fmt.Println(product)
				response = append(response, product)
			}
		}
	} else {
		response = products
	}

	jData, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

}

func readData() []Product {
	jsonData, _ := ioutil.ReadFile("./products.json")
	var products []Product
	json.Unmarshal(jsonData, &products)
	return products
}

type Product struct {
	Id   int
	Name string
}

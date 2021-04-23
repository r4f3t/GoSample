package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var products [4]Product

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("[%s]\t%s:%s\n", request.Method, time.Now(), request.URL)
	//gelen istege karsilik dosya varsa gosterir. yoksa http 404 verir
	http.ServeFile(writer, request, request.URL.Path[1:])
}

func getProduct(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("[%s]\t%s:%s\n", request.Method, time.Now(), request.URL)
	var url string = request.URL.Path
	var result []Product
	parts := strings.Split(url, "/")

	for _, p := range products {
		if strings.EqualFold(p.Category, parts[2]) {
			result = append(result, p)
		}
	}
	if len(result) == 0 {
		http.Error(writer, http.StatusText(404), 404)
	} else {
		json.NewEncoder(writer).Encode(result)
	}
}

func getProducts(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("[%s]\t%s:%s\n", request.Method, time.Now(), request.URL)
	json.NewEncoder(writer).Encode(products)
}

func main() {
	products[0] = Product{Id: 9001, Title: "intel core i5", Category: "CPU", UnitPrice: 150.90}
	products[1] = Product{Id: 9021, Title: "intel core i7", Category: "CPU", UnitPrice: 200.35}
	products[2] = Product{Id: 7800, Title: "google mouse", Category: "OEM", UnitPrice: 44.35}
	products[3] = Product{Id: 1450, Title: "Logitech Wireless Keyboard", Category: "OEM", UnitPrice: 18.98}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/products/", getProduct)
	http.HandleFunc("/products", getProducts)
	log.Fatal(http.ListenAndServe(":8084", nil))
}

type Product struct {
	Id        int     "json:\"ID\""
	Title     string  "json:\"Title\""
	Category  string  "json:\"Category\""
	UnitPrice float32 "json:\"UnitPrice\""
}

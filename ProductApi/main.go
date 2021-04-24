package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	http.HandleFunc("/getProducts", getProducts)

	http.ListenAndServe(":8080", nil)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	var rafet []Product
	var mproducts, _ = db().Database("ProductDb").Collection("Products").Find(context.TODO(), bson.D{})
	mproducts.All(context.Background(), &rafet)

	fmt.Println(rafet)

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

//mongodb works

func db() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Connect to //MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

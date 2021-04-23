package main

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome Our Api"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Our Api Details "))
}

//querystring kullanımı
func getproductsHandler(w http.ResponseWriter, r *http.Request) {
	var marka = r.FormValue("marka")
	var pageIndex = r.FormValue("pageindex")
	var pageSize = r.FormValue("pagesize")

	data := "Marka: " + marka + " Page:" + pageIndex + " PageSize:" + pageSize

	w.Write([]byte(data))
}

func main() {

	// http.ListenAndServe(":8080", nil)

	mesaj1 := &messageHandler{message: "Mesaj1"}
	mesaj2 := &messageHandler{message: "Mesaj2"}

	mux := http.NewServeMux()

	mux.Handle("/mesaj1", mesaj1)
	mux.Handle("/mesaj2", mesaj2)

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/getproducts", getproductsHandler)

	http.ListenAndServe(":8080", mux)
}

//custom handler

type messageHandler struct {
	message string
}

func (x *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(x.message))
}

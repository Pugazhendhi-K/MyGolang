package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ProductNo    string `json:"ProductNo"`
	ProductName  string `json:"Name"`
	ProductType  string `json:"ProductType"`
	Date         string `json:"Date"`
	ProductOwner `json:"ProductOwner"`
}

var Products []Product

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Product Services!")
	fmt.Println("Endpoint Hit: Home")
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllProducts")
	json.NewEncoder(w).Encode(Products)
}

func returnSingleProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ProductNo"]

	for _, product := range Products {
		if Product.ProductNo == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func createNewProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product Products
	json.Unmarshal(reqBody, &product)
	Products = append(Products, product)

	json.NewEncoder(w).Encode(product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productNo := vars["ProductNo"]

	for index, product := range Products {
		if product.ProductNo == productNo {
			Products = append(Products[:index], Products[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", Home)
	myRouter.HandleFunc("/product", returnAllProducts)
	myRouter.HandleFunc("/product", createNewProduct).Methods("POST")
	myRouter.HandleFunc("/product/{productNo}", deleteProduct).Methods("DELETE")
	myRouter.HandleFunc("/product/{productNo}", returnSingleProduct)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Products = []Product{
		Product{ProductNo: "1", ProductName: "Hard Disk", ProductType: "Hardwares", Date: "10-06-2017", ProductOwner: "Segate"},
		Product{ProductNo: "2", ProductName: "Honey", ProductType: "Groceries", Date: "24-09-2020", ProductOwner: "Dabur"},
		Product{ProductNo: "3", ProductName: "T-shirt", ProductType: "Garments", Date: "14-01-2019", ProductOwner: "Polo"},
		Product{ProductNo: "4", ProductName: "Sofa-set", ProductType: "Furnitures", Date: "02-12-2018", ProductOwner: "RoyalOak"},
	}
	handleRequests()
}

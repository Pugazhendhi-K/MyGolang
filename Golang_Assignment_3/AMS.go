package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Article struct {
	ArticleNo string `json:"ArticleNo"`
	Title     string `json:"Title"`
	Author    string `json:"Author"`
	PublishedDate  string `json:"dop"`
}

var Articles []Article

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Article Management System !")
	fmt.Println("Endpoint Hit: Home")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ArticleNo"]

	for _, article := range Articles {
		if article.ArticleNo == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleNo := vars["ArticleNo"]

	for index, article := range Articles {
		if article.ArticleNo == articleNo {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", Home)
	myRouter.HandleFunc("/article", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{articleNo}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{articleNo}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Articles = []Article{
		Article{ArticleNo: "4012", Title: "A View on Cyber Security", Author: "Aano Joanathon", PublishedDate: "25 May 2016"},
		Article{ArticleNo: "2408", Title: "Cloud - The Vritual World", Author: "Rajiv Sharma", PublishedDate: "04 Nov 2019"},
		Article{ArticleNo: "1509", Title: "Future with AI", Author: "Sadiq Ansari", PublishedDate: "29 Feb 2020"}
	}
	handleRequests()
}

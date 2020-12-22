package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func postHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value), // Register No
	})
}
func homePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome",
	})
}
func QueryStrings(c *gin.Context) {
	course := c.Query("course")
	marks := c.Query("marks")
	c.JSON(200, gin.H{
		"course": course,
		"marks":  marks,
	})
}
func PathParameters(c *gin.Context) {
	course := c.Param("course")
	marks := c.Param("marks")
	c.JSON(200, gin.H{
		"course": course,
		"marks":  marks,
	})
}

func main() {
	fmt.Println("Welcome for Result")
	r := gin.Default()
	r.GET("/", homePage)
	r.POST("/", postHomePage)
	r.GET("/query", QueryStrings)
	r.GET("/path/:name/:age", PathParameters)

	fmt.Printf("Starting server at port 8084\n")
	if err := http.ListenAndServe(":8084", nil); err != nil {
		log.Fatal(err)
}

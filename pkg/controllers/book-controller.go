package controllers

import (
	"bookstore/pkg/models"
	"bookstore/pkg/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var NewBook models.Book
func GetBookByID(c *gin.Context){
	query := c.Request.URL.Query()
	id := query["id"][0]
	fmt.Printf("%T %v\n", id, id)
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		fmt.Printf("Error while parsing id: %v %v\n", id, err)
	}
	bookDetails, _ := models.GetBookByID(Id)
	fmt.Println(bookDetails)
	c.IndentedJSON(http.StatusOK, &bookDetails)
	return
}

func GetAllBooks(c *gin.Context){
	fmt.Println("Inside all books")
	newBooks := models.GetAllBooks()
	fmt.Println(newBooks)
	c.IndentedJSON(http.StatusOK, &newBooks)
	return
}


func CreateBook(c *gin.Context){
	CreateBook := &models.Book{}
	utils.ParseBody(c, CreateBook)
	book := CreateBook.CreateBook()
	c.IndentedJSON(http.StatusOK, &book)
	return
}

func DeleteBook(c *gin.Context){
	id, ok := c.GetQuery("id")
	if !ok{
		log.Fatalf("Failed to get id: %v", id)
		return
	}
	Id, err := strconv.ParseInt(id, 0, 0)
	if err != nil{
		log.Fatalf("Failed to parse id: %v", id)
		return
	}
	book := models.DeleteBook(Id)

	c.IndentedJSON(http.StatusOK, &book)
	return
}


func UpdateBook(c *gin.Context){
	var updateBook = &models.Book{}
	utils.ParseBody(c, updateBook)
	id, ok := c.GetQuery("id")
	if !ok{
		log.Fatalf("Failed to get id: %v", id)
		return
	}
	Id, err := strconv.ParseInt(id, 0, 0)
	if err != nil{
		log.Fatalf("Failed to parse id: %v", id)
		return
	}
	bookDetails, db := models.GetBookByID(Id)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	c.IndentedJSON(http.StatusOK, &bookDetails)
	return
}
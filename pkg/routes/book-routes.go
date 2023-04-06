package routes

import (
	"bookstore/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine){
	r.POST("/book", controllers.CreateBook)
	r.GET("/book/:id", controllers.GetBookByID)
	r.GET("/book", controllers.GetAllBooks)
	r.PUT("/book/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.DeleteBook)
}
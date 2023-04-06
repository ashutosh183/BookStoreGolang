package main

import (
	"bookstore/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	routes.Route(r)
	r.Run(":8080")
}
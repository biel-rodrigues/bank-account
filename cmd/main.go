package main

import (
	"net/http"

	"github.com/biel-rodrigues/bank-account/internal/customer"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	r.GET("/customer/create", func(c *gin.Context) {
		c.String(http.StatusOK, customer.Create)
	})

	r.Run()
}

package main

import (
	"encoding/json"
	"net/http"

	"github.com/bank-account/internal/customer"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	r.GET("/customer/create", func(c *gin.Context) {
		returnCustomer(c.Writer, c.Request)
	})

	r.Run()
}

func returnCustomer(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(customer.Create())
}

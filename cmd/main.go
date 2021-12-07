package main

import (
	"net/http"

	"github.com/bank-account/internal/config"
	"github.com/bank-account/internal/customer"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	r := gin.Default()
	loadDB()
	loadRoutes(r)
	r.Run()
}

func loadDB() {
	DB = config.InitDb()
	DB.AutoMigrate(&customer.Address{})
	DB.AutoMigrate(&customer.Customer{})
}

func loadRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	r.POST("/customer/create", func(c *gin.Context) {
		customer.Create(c, DB)
	})

	r.GET("/customer/read/:id", func(c *gin.Context) {
		customer.Read(c, DB)
	})

	r.PUT("/customer/update/:id", func(c *gin.Context) {
		customer.Update(c, DB)
	})

	r.DELETE("/customer/delete/:id", func(c *gin.Context) {
		customer.Delete(c, DB)
	})
}

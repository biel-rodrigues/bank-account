package main

import (
	"net/http"

	"github.com/bank-account/internal/account"
	"github.com/bank-account/internal/config"
	"github.com/bank-account/internal/customer"
	"github.com/bank-account/internal/transfer"
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
	DB.AutoMigrate(&account.Account{})
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

	r.POST("/account/create", func(c *gin.Context) {
		account.Create(c, DB)
	})

	r.GET("/account/read/:id", func(c *gin.Context) {
		account.Read(c, DB)
	})

	r.GET("/account/read-by-number/:number", func(c *gin.Context) {
		account.ReadByNumber(c, DB)
	})

	r.GET("/account/read-all/:customer_id", func(c *gin.Context) {
		account.ReadAllByCustomerID(c, DB)
	})

	r.POST("/transfer", func(c *gin.Context) {
		transfer.Create(c, DB)
	})
}

package customer

import (
	"net/http"

	"github.com/bank-account/internal/dtos/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(c *gin.Context, DB *gorm.DB) {
	var dto request.Customer
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := create(DB, dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"vish": err})
	}

	c.JSON(http.StatusOK, gin.H{"response": customer})
}

func Read(c *gin.Context, DB *gorm.DB) {
	id := c.Param("id")
	customer := read(DB, id)
	c.JSON(http.StatusOK, gin.H{"response": customer})
}

func Update(c *gin.Context, DB *gorm.DB) {
	var dto request.Customer
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	update(DB, id, dto)

	c.String(http.StatusOK, "Customer updated with success")
}

func Delete(c *gin.Context, DB *gorm.DB) {
	id := c.Param("id")
	delete(DB, id)

	c.String(http.StatusOK, "Customer deleted with success")
}

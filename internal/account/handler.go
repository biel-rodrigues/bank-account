package account

import (
	"net/http"
	"strconv"

	"github.com/bank-account/internal/dtos/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(c *gin.Context, DB *gorm.DB) {
	var DTO request.Account
	if err := c.ShouldBindJSON(&DTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := create(DB, DTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, account)
}

func Read(c *gin.Context, DB *gorm.DB) {
	ID := c.Param("id")
	convertedID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an int"})
		return
	}

	account, err := read(DB, convertedID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func ReadByNumber(c *gin.Context, DB *gorm.DB) {
	n := c.Param("number")
	convN, err := strconv.Atoi(n)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Number must be an int"})
		return
	}

	account, err := UseCaseReadByNumber(DB, convN)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func ReadAllByCustomerID(c *gin.Context, DB *gorm.DB) {
	cID := c.Param("customer_id")
	ccID, err := strconv.Atoi(cID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an int"})
		return
	}

	accounts, err := readAllByCustomerID(DB, ccID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, accounts)
}

package account

import (
	"net/http"

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
		c.JSON(http.StatusBadRequest, gin.H{"vish": err})
	}

	c.JSON(http.StatusOK, gin.H{"response": account})
}

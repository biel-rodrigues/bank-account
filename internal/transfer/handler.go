package transfer

import (
	"net/http"

	"github.com/bank-account/internal/dtos/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(c *gin.Context, DB *gorm.DB) {
	var DTO request.Transfer
	if err := c.ShouldBindJSON(&DTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t, err := create(DB, DTO.SN, DTO.RN, DTO.A)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)
}

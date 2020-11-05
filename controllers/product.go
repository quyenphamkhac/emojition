package controllers

import (
	"net/http"

	"github.com/quyenphamkhac/emojition/db"

	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/emojition/forms"
	"github.com/quyenphamkhac/emojition/models"
)

// ProductController struct
type ProductController struct{}

var productModel = new(models.Product)

// Create method
func (p ProductController) Create(c *gin.Context) {
	var input forms.CreateProductInput
	db := db.GetDB()
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	book := models.Product{Price: input.Price, Code: input.Code}
	book.CreateProduct(db)
}

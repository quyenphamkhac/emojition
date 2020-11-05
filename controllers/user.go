package controllers

import (
	"net/http"

	"github.com/quyenphamkhac/emojition/models"

	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/emojition/forms"
)

// UserController ...
type UserController struct{}

var userModel = new(models.UserModel)

// Create ...
func (u UserController) SignUp(c *gin.Context) {
	var input forms.SignUpInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := userModel.SignUp(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mbkstartup/helper"
	"mbkstartup/user"
	"net/http"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userInput, errInput := h.userService.RegisterUser(input)
	if errInput != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println("idnya", userInput.ID)

	formatter := user.FormatUser(userInput, "token in feature")

	response := helper.ApiResponse("Account has been register", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

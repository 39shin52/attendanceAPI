package handler

import (
	"net/http"

	"github.com/39shin52/attendanceApp/app/interfaces/request"
	"github.com/39shin52/attendanceApp/app/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	createUserRequest := new(request.CreateUserRequest)
	if err := c.Bind(&createUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	user, err := uh.userUsecase.CreateUser(c.Request.Context(), createUserRequest.Name, createUserRequest.Affiliation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": user.ID,
	})
}

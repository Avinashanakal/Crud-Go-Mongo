package controllers

import (
	"net/http"

	"github.com/Avinashanakal/models"
	"github.com/Avinashanakal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := uc.UserService.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Created User",
	})
}

func (uc *UserController) GetUser(ctx *gin.Context) {

	name := ctx.Param("name")

	user, err := uc.UserService.GetUser(&name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(ctx *gin.Context) {

	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := uc.UserService.UpdateUser(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {

	name := ctx.Param("name")

	if err := uc.UserService.DeleteUser(&name); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Deleted User",
	})
}

func (uc *UserController) RegisterRoutes(rg *gin.RouterGroup) {
	userRoutes := rg.Group("/user")
	userRoutes.POST("/create", uc.CreateUser)
	userRoutes.GET("/get/all", uc.GetAll)
	userRoutes.GET("/get/:name", uc.GetUser)
	userRoutes.PATCH("/update", uc.UpdateUser)
	userRoutes.DELETE("/delete/:name", uc.DeleteUser)
}

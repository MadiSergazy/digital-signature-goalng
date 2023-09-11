package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	// "mado/internal"
	"mado/internal/auth"
	"mado/internal/auth/model"

	// "mado/internal/controller/http/httperr"
	"mado/internal/core/user"
)

type ECP struct {
	Ecp string `json:"ecp"    binding:"required,ecp"`
}

// UserService is a user service interface.
type UserService interface {
	// Create(ctx context.Context, dto user.CreateDTO) (user.User, error)
	Login(model.LoginRequirements) (*user.User, error)
	LogOut(ctx context.Context, user *user.User) error
}

type userDeps struct {
	router *gin.RouterGroup

	userService UserService
}

type userHandler struct {
	userService UserService
}

func newUserHandler(deps userDeps) {
	handler := userHandler{
		userService: deps.userService,
	}

	usersGroup := deps.router.Group("/users")
	{
		usersGroup.GET("/", handler.getUser)
		usersGroup.POST("/", handler.createUser)    // api/users/
		usersGroup.POST("/login", handler.sendLink) // api/users/login
		usersGroup.POST("/confirm", handler.confirmCredentials)

	}

}

func (h userHandler) createUser(c *gin.Context) {
	fmt.Println("createUser")
	c.IndentedJSON(http.StatusOK, "User created")

}

func (h userHandler) getUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, http.StatusText(http.StatusOK))
	fmt.Println("GetUser")
}

// TODO implement this properly
type loginUserRequest struct {
	User struct {
		Email    string `json:"email"    binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	} `json:"user" binding:"required"`
}

type loginUserResponse struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	IIN      *string `json:"iin"`
	BIN      *string `json:"bin"`
}

// TODO implement this properly
// func (h userHandler) loginUser(c *gin.Context) {
// 	var request loginUserRequest
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		httperr.BadRequest(c, "invalid-request", err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{})
// }

func (h userHandler) sendLink(c *gin.Context) {
	egovMobileLink, qrSigner, nonce := auth.PreparationStep()
	if egovMobileLink == nil || qrSigner == nil || nonce == nil {
		fmt.Println("egovMobileLink: ", egovMobileLink)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": "egovMobileLink or qrSigner or nonce is nil"})
		return
	}
	requirements := model.LoginRequirements{QrSigner: qrSigner, Nonce: nonce}
	// go h.userService.Login(context.Background(), qrSigner, nonce)
	c.JSON(http.StatusOK, gin.H{"link": egovMobileLink, "requirements": requirements})
	return
}

func (h userHandler) confirmCredentials(c *gin.Context) {
	var request model.LoginRequirements
	if err := c.ShouldBindJSON(request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	h.userService.Login(request)
	c.JSON(http.StatusOK, gin.H{})
	return
}

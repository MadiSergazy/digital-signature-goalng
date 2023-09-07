package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"mado/internal/auth"
	"mado/internal/controller/http/httperr"
	"mado/internal/core/user"
)

type ECP struct {
	Ecp string `json:"ecp"    binding:"required,ecp"`
}

// UserService is a user service interface.
type UserService interface {
	// Create(ctx context.Context, dto user.CreateDTO) (user.User, error)
	Login(ctx context.Context, user *user.User) (*user.User, error)
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
		usersGroup.POST("/", handler.createUser)     // api/users/
		usersGroup.POST("/login", handler.loginUser) // api/users/login
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
func (h userHandler) loginUser(c *gin.Context) {
	var request loginUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		httperr.BadRequest(c, "invalid-request", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h userHandler) authUser(c *gin.Context) {
	var request ECP
	if err := c.ShouldBind(&request); err != nil {
		httperr.BadRequest(c, "invalid-request", err)
		return
	}
	egovMobileLink, qrSigner := auth.PreparationStep()
	c.JSON(http.StatusOK, gin.H{"link": egovMobileLink})
	auth.GetNonceSignature(qrSigner)
}

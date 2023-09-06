package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"mado/internal/controller/http/httperr"
	"mado/internal/core/user"
)

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

	// userEntity, err := h.userService.Create(logger.FromRequestToContext(c), user.CreateDTO{
	// 	Email:    request.User.Email,
	// 	Username: request.User.Username,
	// 	Password: request.User.Password,
	// })
	// if err != nil {
	// 	httperr.RespondWithSlugError(c, err)
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		// "user": loginUserResponse{
		// 	Email:    userEntity.Email,
		// 	Username: userEntity.Username,
		// 	IIN:      userEntity,
		// 	BIN:      userEntity,
		// },
	})
}

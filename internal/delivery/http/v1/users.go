package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.GET("/", getAllUsers)
		users.GET("/:id", getOneUser)
		users.POST("/", createUser)
		users.PUT("/:id", updateUser)
		users.DELETE("/:id", deleteUser)
	}
}

func updateUser(context *gin.Context) {
	context.String(http.StatusOK, "User updated")
}

func deleteUser(context *gin.Context) {
	context.String(http.StatusOK, "User deleted")
}

func createUser(context *gin.Context) {
	context.String(http.StatusOK, "User created")
}

func getOneUser(context *gin.Context) {
	id := context.Param("id")
	if id == "1" {
		context.String(http.StatusBadRequest, "Invalid id of user")
	} else {
		context.String(http.StatusOK, fmt.Sprintf("Get user with id=%s", id))
	}
}

func getAllUsers(context *gin.Context) {
	context.String(http.StatusOK, "Get all users")
}

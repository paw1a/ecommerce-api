package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

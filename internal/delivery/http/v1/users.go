package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"net/http"
)

func (h *Handler) getAllUsersAdmin(context *gin.Context) {
	users, err := h.services.Users.FindAll(context.Request.Context())
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	usersArray := make([]domain.User, len(users))
	if users != nil {
		usersArray = users
	}

	successResponse(context, usersArray)
}

func (h *Handler) getUserByIdAdmin(context *gin.Context) {
	id, err := parseIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.Users.FindByID(context.Request.Context(), id)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, user)
}

func (h *Handler) createUserAdmin(context *gin.Context) {
	var userDTO dto.CreateUserDTO
	err := context.BindJSON(&userDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "Invalid input body")
		return
	}
	user, err := h.services.Users.Create(context.Request.Context(), userDTO)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, user)
}

func (h *Handler) updateUserAdmin(context *gin.Context) {
	var userDTO dto.UpdateUserDTO

	err := context.BindJSON(&userDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}

	userID, err := parseIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.Users.Update(context.Request.Context(), userDTO, userID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, user)
}

func (h *Handler) deleteUserAdmin(context *gin.Context) {
	userID, err := parseIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Users.Delete(context, userID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.Status(http.StatusOK)
}

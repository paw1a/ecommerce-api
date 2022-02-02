package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// GetUsers godoc
// @Summary   Get all users
// @Tags      admin-users
// @Accept    json
// @Produce   json
// @Success   200  {array}   success
// @Failure   401   {object}  failure
// @Failure   404   {object}  failure
// @Failure   500   {object}  failure
// @Security  AdminAuth
// @Router    /admins/users [get]
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

// GetUserById godoc
// @Summary   Get user by id
// @Tags      admin-users
// @Accept    json
// @Produce   json
// @Param     id   path      string  true  "user id"
// @Success   200   {object}  success
// @Failure   400   {object}  failure
// @Failure   401   {object}  failure
// @Failure   404   {object}  failure
// @Failure   500   {object}  failure
// @Security  AdminAuth
// @Router    /admins/users/{id} [get]
func (h *Handler) getUserByIdAdmin(context *gin.Context) {
	id, err := getIdFromPath(context, "id")
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

// CreateUser godoc
// @Summary   Create user
// @Tags      admin-users
// @Accept    json
// @Produce   json
// @Param     user  body      dto.CreateUserDTO  true  "user"
// @Success   200   {object}  success
// @Failure   400   {object}  failure
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  AdminAuth
// @Router    /admins/users [post]
func (h *Handler) createUserAdmin(context *gin.Context) {
	var userDTO dto.CreateUserDTO
	err := context.BindJSON(&userDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}
	user, err := h.services.Users.Create(context.Request.Context(), userDTO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			errorResponse(context, http.StatusInternalServerError,
				fmt.Sprintf("user with email %s already exists", userDTO.Email))
		} else {
			errorResponse(context, http.StatusInternalServerError, err.Error())
		}
		return
	}

	createdResponse(context, user)
}

// UpdateUser godoc
// @Summary   Update user
// @Tags      admin-users
// @Accept    json
// @Produce   json
// @Param     id    path      string             true  "user id"
// @Param     user  body      dto.UpdateUserDTO  true  "user update fields"
// @Success   200  {object}  success
// @Failure   400  {object}  failure
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  AdminAuth
// @Router    /admins/users/{id} [put]
func (h *Handler) updateUserAdmin(context *gin.Context) {
	var userDTO dto.UpdateUserDTO

	err := context.BindJSON(&userDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}

	userID, err := getIdFromPath(context, "id")
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

// DeleteUser godoc
// @Summary   Delete user
// @Tags      admin-users
// @Accept    json
// @Produce   json
// @Param     id   path      string  true  "user id"
// @Success   200  {object}  success
// @Failure   400  {object}  failure
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  AdminAuth
// @Router    /admins/users/{id} [delete]
func (h *Handler) deleteUserAdmin(context *gin.Context) {
	userID, err := getIdFromPath(context, "id")
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

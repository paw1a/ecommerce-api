package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func (h *Handler) initCartRoutes(api *gin.RouterGroup) {

}

// GetCarts godoc
// @Summary   Get all carts
// @Tags      admin-carts
// @Accept    json
// @Produce   json
// @Success   200  {array}   success
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  AdminAuth
// @Router    /admins/carts [get]
func (h *Handler) getAllCartsAdmin(context *gin.Context) {
	carts, err := h.services.Carts.FindAll(context.Request.Context())
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	cartsArray := make([]domain.Cart, len(carts))
	if carts != nil {
		cartsArray = carts
	}

	successResponse(context, cartsArray)
}

// GetCartByIdAdmin godoc
// @Summary   Get cart by id
// @Tags      admin-carts
// @Accept    json
// @Produce   json
// @Param     id   path      string  true  "cart id"
// @Success   200  {object}  success
// @Failure   400  {object}  failure
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  AdminAuth
// @Router    /admins/carts/{id} [get]
func (h *Handler) getCartByIdAdmin(context *gin.Context) {
	id, err := getIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	cart, err := h.services.Carts.FindByID(context.Request.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorResponse(context, http.StatusInternalServerError,
				fmt.Sprintf("no carts with id: %s", id.Hex()))
		} else {
			errorResponse(context, http.StatusInternalServerError, err.Error())
		}
		return
	}

	successResponse(context, cart)
}

// DeleteCart godoc
// @Summary   Delete cart
// @Tags      admin-carts
// @Accept    json
// @Produce   json
// @Param     id   path      string  true  "cart id"
// @Success   200  {object}  success
// @Failure   400  {object}  failure
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  AdminAuth
// @Router    /admins/carts/{id} [delete]
func (h *Handler) deleteCartAdmin(context *gin.Context) {
	cartID, err := getIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Carts.Delete(context, cartID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.Status(http.StatusOK)
}

package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

func (h *Handler) initCartRoutes(api *gin.RouterGroup) {
	cart := api.Group("/cart", h.extractCartId)
	{
		cart.GET("/", h.getCartItems)
		cart.POST("/", h.createCartItem)
		cart.PUT("/:id", h.updateCartItem)
		cart.DELETE("/:id", h.deleteCartItem)
	}
}

func (h *Handler) extractCartId(context *gin.Context) {
	cartIDHex, err := context.Cookie("cartID")
	if err == nil {
		cartID, err := primitive.ObjectIDFromHex(cartIDHex)
		if err == nil {
			_, err = h.services.Carts.FindByID(context.Request.Context(), cartID)
			if err == nil {
				context.Set("cartID", cartID)
				return
			} else {
				log.Warnf("cart with id: %s not found in db", cartIDHex)
			}
		} else {
			log.Warnf("failed to convert cookie %s to objectID", cartIDHex)
		}
	}

	log.Warn("cartID cookie not found")
	userID, err := h.extractIdFromAuthHeader(context, "userID")
	if err == nil {
		user, err := h.services.Users.FindByID(context.Request.Context(), userID)
		if err != nil {
			errorResponse(context, http.StatusInternalServerError,
				fmt.Sprintf("user with id: %s not found", userID))
			return
		}

		_, err = h.services.Carts.FindByID(context.Request.Context(), user.CartID)
		if err != nil {
			log.Warnf("user with id: %s don't have cart with id: %s", userID, user.CartID)

			newCart, err := h.services.Carts.Create(context.Request.Context(), dto.CreateCartDTO{
				ExpireAt:  time.Now().Add(30 * time.Hour * 24),
				CartItems: nil,
			})
			if err != nil {
				errorResponse(context, http.StatusInternalServerError, err.Error())
				return
			}
			user.CartID = newCart.ID

			_, err = h.services.Users.Update(context, dto.UpdateUserDTO{CartID: &user.CartID}, userID)
			if err != nil {
				errorResponse(context, http.StatusInternalServerError, err.Error())
				return
			}
		}
		context.Set("cartID", user.CartID)
		return
	}

	log.Warnf("user not authenticated")
	newCart, err := h.services.Carts.Create(context.Request.Context(), dto.CreateCartDTO{
		ExpireAt:  time.Now().Add(30 * time.Hour * 24),
		CartItems: nil,
	})
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	context.Set("cartID", newCart.ID)
}

func (h *Handler) getCartItems(context *gin.Context) {
	cartIDHex, ok := context.Get("cartID")
	if !ok {
		errorResponse(context, http.StatusInternalServerError, "failed to get cart id")
		return
	}
	cartID := cartIDHex.(primitive.ObjectID)

	cartItems, err := h.services.Carts.FindCartItems(context, cartID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, cartItems)
}

func (h *Handler) createCartItem(context *gin.Context) {
	cartIDHex, ok := context.Get("cartID")
	if !ok {
		errorResponse(context, http.StatusInternalServerError, "failed to get cart id")
		return
	}
	cartID := cartIDHex.(primitive.ObjectID)

	var cartItemInput domain.CartItem
	err := context.BindJSON(&cartItemInput)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}

	cartItem, err := h.services.Carts.AddCartItem(context, cartItemInput, cartID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	createdResponse(context, cartItem)
}

func (h *Handler) updateCartItem(context *gin.Context) {
	cartIDHex, ok := context.Get("cartID")
	if !ok {
		errorResponse(context, http.StatusInternalServerError, "failed to get cart id")
		return
	}
	cartID := cartIDHex.(primitive.ObjectID)

	productID, err := getIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	var cartItemInput dto.UpdateCartItemDTO
	err = context.BindJSON(&cartItemInput)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}

	cartItem, err := h.services.Carts.UpdateCartItem(context, domain.CartItem{
		ProductID: productID,
		Quantity:  cartItemInput.Quantity,
	}, cartID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, cartItem)
}

func (h *Handler) deleteCartItem(context *gin.Context) {
	cartIDHex, ok := context.Get("cartID")
	if !ok {
		errorResponse(context, http.StatusInternalServerError, "failed to get cart id")
		return
	}
	cartID := cartIDHex.(primitive.ObjectID)

	productID, err := getIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Carts.DeleteCartItem(context, productID, cartID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.Status(http.StatusOK)
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

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"net/http"
)

func (h *Handler) initOrdersRoutes(api *gin.RouterGroup) {

}

// GerUserOrders godoc
// @Summary   User order List
// @Tags      user
// @Accept    json
// @Produce   json
// @Success   200  {array}   success
// @Failure   401    {object}  failure
// @Failure   404    {object}  failure
// @Failure   500    {object}  failure
// @Security  UserAuth
// @Router    /users/orders [get]
func (h *Handler) getUserOrders(context *gin.Context) {
	userID, err := getIdFromRequestContext(context, "userID")
	if err != nil {
		errorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}

	orders, err := h.services.Orders.FindByUserID(context.Request.Context(), userID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, orders)
}

// CreateOrder godoc
// @Summary   Create order
// @Tags      user
// @Accept    json
// @Produce   json
// @Param     order  body      dto.CreateOrderDTO  true  "contact info"
// @Success   201    {object}  success
// @Failure   400  {object}  failure
// @Failure   401    {object}  failure
// @Failure   404    {object}  failure
// @Failure   500    {object}  failure
// @Security  UserAuth
// @Router    /users/orders [post]
func (h *Handler) createOrder(context *gin.Context) {
	userID, err := getIdFromRequestContext(context, "userID")
	if err != nil {
		errorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.services.Users.FindByID(context.Request.Context(), userID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	cart, err := h.services.Carts.FindByID(context.Request.Context(), user.CartID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	if len(cart.CartItems) == 0 {
		errorResponse(context, http.StatusBadRequest, "user cart is empty")
		return
	}

	orderItems := make([]domain.OrderItem, len(cart.CartItems))
	for i, cartItem := range cart.CartItems {
		orderItems[i] = domain.OrderItem{
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
		}
	}

	var createOrderDTO dto.CreateOrderDTO
	err = context.BindJSON(&createOrderDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}

	order, err := h.services.Orders.Create(context.Request.Context(), dto.CreateOrderDTO{
		OrderItems:  orderItems,
		ContactInfo: createOrderDTO.ContactInfo,
		UserID:      userID,
	})

	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Carts.ClearCart(context.Request.Context(), cart.ID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, "cart can't be cleared")
		return
	}

	successResponse(context, order)
}

// PaymentLink godoc
// @Summary   Get order payment link
// @Tags      user
// @Accept    json
// @Produce   json
// @Param     id   path      string  true  "order id"
// @Success   200  {object}  success
// @Failure   400    {object}  failure
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  UserAuth
// @Router    /users/orders/{id}/payment [get]
func (h *Handler) getOrderPaymentLink(context *gin.Context) {
	orderID, err := getIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	order, err := h.services.Orders.FindByID(context.Request.Context(), orderID)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	if order.Status != "reserved" {
		errorResponse(context, http.StatusBadRequest,
			fmt.Sprintf("order with id: %s already paid", order.OrderID))
		return
	}

	link, err := h.services.Payment.GetPaymentUrl(order)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, link)
}

// GetOrdersAdmin godoc
// @Summary   Get all orders
// @Tags      admin-orders
// @Accept    json
// @Produce   json
// @Success   200  {array}   success
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  AdminAuth
// @Router    /admins/orders [get]
func (h *Handler) getAllOrdersAdmin(context *gin.Context) {
	orders, err := h.services.Orders.FindAll(context.Request.Context())
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	orderArray := make([]domain.Order, len(orders))
	if orders != nil {
		orderArray = orders
	}

	successResponse(context, orderArray)
}

// UpdateOrder godoc
// @Summary   Update order
// @Tags      admin-orders
// @Accept    json
// @Produce   json
// @Param     id     path      string              true  "order id"
// @Param     order  body      dto.UpdateOrderDTO  true  "order update fields"
// @Success   200    {object}  success
// @Failure   400    {object}  failure
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  AdminAuth
// @Router    /admins/orders/{id} [put]
func (h *Handler) updateOrderAdmin(context *gin.Context) {
	var orderDTO dto.UpdateOrderDTO

	err := context.BindJSON(&orderDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}

	orderID, err := getIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	order, err := h.services.Orders.Update(context.Request.Context(), orderDTO, orderID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, order)
}

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"net/http"
)

func (h *Handler) initOrdersRoutes(api *gin.RouterGroup) {

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
// @Param     id       path      string                true  "order id"
// @Param     order  body      dto.UpdateOrderDTO  true  "order update fields"
// @Success   200  {object}  success
// @Failure   400  {object}  failure
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

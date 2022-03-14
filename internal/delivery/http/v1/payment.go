package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	log "github.com/sirupsen/logrus"
	"github.com/stripe/stripe-go/v72"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (h *Handler) initPaymentRoutes(api *gin.RouterGroup) {
	api.POST("/payment/webhook", h.webhookSessionCompleted)
}

func (h *Handler) webhookSessionCompleted(context *gin.Context) {
	var event stripe.Event
	err := context.BindJSON(&event)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "can't parse event from request")
		return
	}

	orderIDHex := event.Data.Object["client_reference_id"].(string)
	orderID, err := primitive.ObjectIDFromHex(orderIDHex)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, "can't parse orderID")
		return
	}

	_, err = h.services.Orders.Update(context.Request.Context(), dto.UpdateOrderDTO{Status: "paid"}, orderID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	log.Warn("PAID")
}

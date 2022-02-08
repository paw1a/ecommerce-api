package v1

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) initPaymentRoutes(api *gin.RouterGroup) {
	api.POST("/payment/webhook", h.webhook)
}

func (h *Handler) webhook(context *gin.Context) {
	log.Error("HELLO from webhook")
}

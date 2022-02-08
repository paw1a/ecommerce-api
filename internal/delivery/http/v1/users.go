package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/auth/sign-in", h.userSignIn)
		users.POST("/auth/sign-up", h.userSignUp)
		users.POST("/auth/refresh", h.userRefresh)

		authenticated := users.Group("/", h.verifyUser)
		{
			authenticated.GET("/account", h.getUserAccount)
			authenticated.GET("/reviews", h.getAllReviewsUser)
			orders := authenticated.Group("/orders")
			{
				orders.GET("/", h.getUserOrders)
				orders.POST("/", h.createOrder)
				orders.GET("/:id/payment", h.getOrderPaymentLink)
			}
		}
	}
}

// GerUserReviews godoc
// @Summary   User reviews List
// @Tags      user
// @Accept    json
// @Produce   json
// @Success   200  {array}   success
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  UserAuth
// @Router    /users/reviews [get]
func (h *Handler) getAllReviewsUser(context *gin.Context) {
	userID, err := getIdFromRequestContext(context, "userID")
	if err != nil {
		errorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}

	reviews, err := h.services.Reviews.FindByUserID(context.Request.Context(), userID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, reviews)
}

// UserAccount godoc
// @Summary   User account
// @Tags      user
// @Accept    json
// @Produce   json
// @Success   200  {object}  auth.AuthDetails
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  UserAuth
// @Router    /users/account [get]
func (h *Handler) getUserAccount(context *gin.Context) {
	userID, err := getIdFromRequestContext(context, "userID")
	if err != nil {
		errorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}

	userInfo, err := h.services.Users.FindUserInfo(context.Request.Context(), userID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, userInfo)
}

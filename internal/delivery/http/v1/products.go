package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func (h *Handler) initProductsRoutes(api *gin.RouterGroup) {
	products := api.Group("/products")
	{
		products.GET("/", h.getAllProducts)
		products.GET("/:id", h.getProductById)
		products.GET("/:id/reviews", h.getProductReviews)

		authenticated := products.Group("/", h.verifyUser)
		{
			authenticated.POST("/:id/reviews", h.createProductReview)
		}
	}
}

func (h *Handler) getAllProducts(context *gin.Context) {

}

func (h *Handler) getProductById(context *gin.Context) {

}

func (h *Handler) getProductReviews(context *gin.Context) {

}

func (h *Handler) createProductReview(context *gin.Context) {

}

// GetProducts godoc
// @Summary   Get all products
// @Tags      admin-products
// @Accept    json
// @Produce   json
// @Success   200  {array}   success
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  AdminAuth
// @Router    /admins/products [get]
func (h *Handler) getAllProductsAdmin(context *gin.Context) {
	products, err := h.services.Products.FindAll(context.Request.Context())
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	productsArray := make([]domain.Product, len(products))
	if products != nil {
		productsArray = products
	}

	successResponse(context, productsArray)
}

// GetProductById godoc
// @Summary   Get product by id
// @Tags      admin-products
// @Accept    json
// @Produce   json
// @Param     id   path      string  true  "product id"
// @Success   200  {object}  success
// @Failure   400  {object}  failure
// @Failure   401      {object}  failure
// @Failure   404      {object}  failure
// @Failure   500      {object}  failure
// @Security  AdminAuth
// @Router    /admins/products/{id} [get]
func (h *Handler) getProductByIdAdmin(context *gin.Context) {
	id, err := getIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	product, err := h.services.Products.FindByID(context.Request.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorResponse(context, http.StatusInternalServerError,
				fmt.Sprintf("no products with id: %s", id.Hex()))
		} else {
			errorResponse(context, http.StatusInternalServerError, err.Error())
		}
		return
	}

	successResponse(context, product)
}

// CreateProduct godoc
// @Summary   Create product
// @Tags      admin-products
// @Accept    json
// @Produce   json
// @Param     product  body      dto.CreateProductDTO  true  "product"
// @Success   201      {object}  success
// @Failure   400      {object}  failure
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  AdminAuth
// @Router    /admins/products [post]
func (h *Handler) createProductAdmin(context *gin.Context) {
	var productDTO dto.CreateProductDTO
	err := context.BindJSON(&productDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}
	product, err := h.services.Products.Create(context.Request.Context(), productDTO)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, product)
}

// UpdateProduct godoc
// @Summary   Update product
// @Tags      admin-products
// @Accept    json
// @Produce   json
// @Param     id       path      string                true  "product id"
// @Param     product  body      dto.UpdateProductDTO  true  "product update fields"
// @Success   200  {object}  success
// @Failure   400  {object}  failure
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  AdminAuth
// @Router    /admins/products/{id} [put]
func (h *Handler) updateProductAdmin(context *gin.Context) {
	var productDTO dto.UpdateProductDTO

	err := context.BindJSON(&productDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}

	productID, err := getIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.services.Products.Update(context.Request.Context(), productDTO, productID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, product)
}

// DeleteProduct godoc
// @Summary   Delete product
// @Tags      admin-products
// @Accept    json
// @Produce   json
// @Param     id   path      string  true  "product id"
// @Success   200  {object}  success
// @Failure   400  {object}  failure
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  AdminAuth
// @Router    /admins/products/{id} [delete]
func (h *Handler) deleteProductAdmin(context *gin.Context) {
	productID, err := getIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Products.Delete(context, productID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.Status(http.StatusOK)
}

// GetProductReviews godoc
// @Summary   Get product reviews list
// @Tags      admin-products
// @Accept    json
// @Produce   json
// @Param     id   path      string  true  "product id"
// @Success   200      {object}  success
// @Failure   400      {object}  failure
// @Failure   401      {object}  failure
// @Failure   404      {object}  failure
// @Failure   500      {object}  failure
// @Security  AdminAuth
// @Router    /admins/products/{id}/reviews [get]
func (h *Handler) getProductReviewsAdmin(context *gin.Context) {
	productID, err := getIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	reviews, err := h.services.Reviews.FindByProductID(context, productID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, reviews)
}

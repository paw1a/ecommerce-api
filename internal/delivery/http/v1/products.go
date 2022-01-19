package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"github.com/paw1a/ecommerce-api/internal/domain/dto"
	"net/http"
)

func (h *Handler) getAllProducts(context *gin.Context) {
	products, err := h.services.Products.FindAll(context.Request.Context())
	if err != nil {
		newResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	productsArray := make([]domain.Product, len(products))
	if products != nil {
		productsArray = products
	}

	context.JSON(http.StatusOK, dataResponse{Data: productsArray})
}

func (h *Handler) getProductById(context *gin.Context) {
	id, err := parseIdFromPath(context, "id")
	if err != nil {
		newResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	product, err := h.services.Products.FindByID(context.Request.Context(), id)
	if err != nil {
		newResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, dataResponse{Data: product})
}

func (h *Handler) createProduct(context *gin.Context) {
	var productDTO dto.CreateProductDTO
	err := context.BindJSON(&productDTO)
	if err != nil {
		newResponse(context, http.StatusBadRequest, "Invalid input body")
		return
	}
	product, err := h.services.Products.Create(context.Request.Context(), productDTO)
	if err != nil {
		newResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, dataResponse{Data: product})
}

func (h *Handler) updateProduct(context *gin.Context) {
	var productDTO dto.UpdateProductDTO

	err := context.BindJSON(&productDTO)
	if err != nil {
		newResponse(context, http.StatusBadRequest, "Invalid input body")
		return
	}

	productID, err := parseIdFromPath(context, "id")
	if err != nil {
		newResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	productDTO.ID = productID

	product, err := h.services.Products.Update(context.Request.Context(), productDTO)
	if err != nil {
		newResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, dataResponse{Data: product})
}

func (h *Handler) deleteProduct(context *gin.Context) {
	productID, err := parseIdFromPath(context, "id")
	if err != nil {
		newResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Products.Delete(context, productID)
	if err != nil {
		newResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.Status(http.StatusOK)
}

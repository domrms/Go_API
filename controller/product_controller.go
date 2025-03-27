package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(c *gin.Context) {

	products, err := p.productUseCase.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(c *gin.Context) {
	var product model.Product
	c.BindJSON(&product)

	err := p.productUseCase.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (p *productController) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response := model.Response{
			Message: "ID is required",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{}
		response.Message = "ID must be a number"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	product, err := p.productUseCase.GetProductByID(productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

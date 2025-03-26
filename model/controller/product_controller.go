package controller

import (
	"go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type producController struct {
}

func NewProductController() producController {
	return producController{}
}

func (p *producController) GetProducts(c *gin.Context) {

	products := []model.Product{
		{
			ID:    1,
			Name:  "Laptop",
			Price: 1000.00,
		},
	}

	c.JSON(http.StatusOK, products)

}

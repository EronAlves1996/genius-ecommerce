package product

import (
	"github.com/gin-gonic/gin"
)

var getAllProducts = GetAll

func RegisterController(r *gin.Engine) {
	r.GET("/products", func(c *gin.Context) {
		ps, err := getAllProducts()
		if err != nil {
			c.Error(err)
		}
		c.JSON(200, ps)
	})
}

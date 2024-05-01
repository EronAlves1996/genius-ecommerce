package product

import "github.com/gin-gonic/gin"

func RegisterController(r *gin.Engine) {

	r.GET("/products", func(c *gin.Context) {
		ps := []Product{
			{
				Id:       1,
				Name:     "Test Product",
				Price:    20.90,
				ImageUrl: "./test.jpg",
			},
			{
				Id:       2,
				Name:     "Second Product",
				Price:    10.20,
				ImageUrl: "./test2.jpg",
			},
		}

		c.JSON(200, ps)
	})
}

package product

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterController(r *gin.Engine, db *sql.DB) {

	r.GET("/products", func(c *gin.Context) {
		ps, err := GetAll(db)
		if err != nil {
			c.Error(err)
		}
		c.JSON(200, ps)
	})
}

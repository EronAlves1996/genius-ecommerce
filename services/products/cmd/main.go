package main

import (
	"log"

	"github.com/EronAlves1996/services/products/internal/product"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	doMigration()
	g := registerControllers()
	g.Run(":8080")
}

func doMigration() {
	m, err := migrate.New("file://db/migrations", "postgres://root:root@localhost:5432/gennius-products?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

func registerControllers() *gin.Engine {
	g := gin.Default()
	product.RegisterController(g)
	return g
}

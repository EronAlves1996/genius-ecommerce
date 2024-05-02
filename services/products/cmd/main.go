package main

import (
	"database/sql"
	"log"

	"github.com/EronAlves1996/services/products/internal/product"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	doMigration()
	db := createDbConnection()
	g := registerControllers(db)
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

func createDbConnection() *sql.DB {
	connStr := "postgres://root:root@localhost:5432/gennius-products?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func registerControllers(db *sql.DB) *gin.Engine {
	g := gin.Default()
	product.RegisterController(g, db)
	return g
}

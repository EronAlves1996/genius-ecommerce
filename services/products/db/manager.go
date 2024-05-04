package db

import (
	"database/sql"
	"log"
	"sync"
)

var connStr = "postgres://root:root@localhost:5432/gennius-products?sslmode=disable"
var connection *sql.DB
var m sync.Mutex

func GetConnection() *sql.DB {
	if connection == nil {
		acquireConnection()
	}
	return connection
}

func acquireConnection() {
	m.Lock()
	defer m.Unlock()
	var err error
	// Double check here, because some other goroutine
	// can enter lock, and we don't leak because of
	// another goroutine awaiting the lock get free
	if connection == nil {
		connection, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
	}
}

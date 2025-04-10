package service

import (
	"context"
	"database/sql"
	"time"
)

// OpenDB will start a new connection using the selected database driver and data source
func OpenDB(driver string, dsn string) (*sql.DB, error) {
	// open the database
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	//try to connect to pool within 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	// return the connection pool
	return db, nil
}

// Services will eventally provide the database connetion pool to various services
type Services struct{}

// NewServices instantiates a new Service struct
func NewServices(db *sql.DB) *Services {
	return &Services{}
}

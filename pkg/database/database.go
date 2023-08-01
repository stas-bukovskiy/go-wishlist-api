package database

import (
	"context"
)

type Database interface {
	// Close closes the connection to database.
	Close() error
	// Ping - checks if database is available.
	Ping(ctx context.Context) error
}

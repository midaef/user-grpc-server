package pkg

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // ...
)

// Connection ...
type Connection struct {
	databaseURI string
	DB          *sqlx.DB
}

// NewConnection ...
func NewConnection(uri string) *Connection {
	return &Connection{
		databaseURI: uri,
	}
}

// Open ...
func (conn *Connection) Open() error {
	ctx := context.Background()

	db, err := sqlx.ConnectContext(ctx, "postgres", conn.databaseURI)
	if err != nil {
		return err
	}

	conn.DB = db

	return nil
}

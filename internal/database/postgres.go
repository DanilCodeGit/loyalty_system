package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

type DB struct {
	Conn *pgxpool.Pool
	//mu   sync.RWMutex
}

func NewDataBase(ctx context.Context, dsn string) (*DB, error) {
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err)
		return nil, err
	}
	log.Println("Успешное подключение")

	return &DB{Conn: conn}, nil
}

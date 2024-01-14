package repository

import (
	"context"
	"github.com/DanilCodeGit/loyalty_system/go-musthave-diploma-tpl/internal/entity"
	"github.com/jackc/pgx/v5"
)

type User struct {
	db *pgx.Conn
}

func NewDB(db *pgx.Conn) User {
	return User{db: db}
}

func (u User) Registration(ctx context.Context, user entity.User) error {
	return nil
}

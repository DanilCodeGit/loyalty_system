package repository

import (
	"context"
	"github.com/DanilCodeGit/loyalty_system/go-musthave-diploma-tpl/internal/entity"
	"github.com/jackc/pgx/v5"
)

type Users struct {
	db *pgx.Conn
}

func NewDB(db *pgx.Conn) Users {
	return Users{db: db}
}

func (u Users) Registration(ctx context.Context, user entity.User) error {
	return nil
}

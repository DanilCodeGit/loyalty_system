package repository

import (
	"context"
	"github.com/DanilCodeGit/loyalty_system/internal/database"
	"github.com/DanilCodeGit/loyalty_system/internal/domain"
	"github.com/DanilCodeGit/loyalty_system/internal/entity"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type Users struct {
	db database.DB
}

func NewUsers(db *database.DB) Users {
	return Users{db: *db}
}

func (u Users) Registration(ctx context.Context, user domain.Users) error {
	query := `insert into users (login , password) values ($1,$2)`
	_, err := u.db.Conn.Exec(ctx, query, user.Login, user.Password)
	if err != nil {
		return errors.WithMessage(err, "exec query")
	}
	return nil
}

func (u Users) IsExists(ctx context.Context, user domain.Users) error {
	query := `select login from users where login=$1`
	var res entity.User
	err := u.db.Conn.QueryRow(ctx, query, user.Login).Scan(&res)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil // Пользователь не найден, возвращаем nil
		}
		return err
	}

	return errors.New("user with this login already exists")
}

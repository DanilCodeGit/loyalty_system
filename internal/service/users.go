package service

import (
	"context"
	"github.com/DanilCodeGit/loyalty_system/go-musthave-diploma-tpl/internal/entity"
)

type UsersRepo interface {
	Registration(ctx context.Context, user entity.User) error
}

type Users struct {
	r UsersRepo
}

func NewUsers(repo UsersRepo) Users {
	return Users{
		r: repo,
	}
}

func (u Users) Registration(ctx context.Context, user entity.User) error {
	return nil
}

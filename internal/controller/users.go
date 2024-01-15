package controller

import (
	"context"
	"github.com/DanilCodeGit/loyalty_system/go-musthave-diploma-tpl/internal/entity"
)

type UserService interface {
	Registration(ctx context.Context, user entity.User) error
}

type Users struct {
	s UserService
}

func NewUsers(s UserService) Users {
	return Users{s: s}
}

func Registration(ctx context.Context, user entity.User) error {

}

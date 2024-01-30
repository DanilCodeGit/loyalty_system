package service

import (
	"context"
	"github.com/DanilCodeGit/loyalty_system/internal/domain"
	"github.com/pkg/errors"
)

type UsersRepo interface {
	Registration(ctx context.Context, user domain.Users) error
	IsExists(ctx context.Context, user domain.Users) error
}

type Users struct {
	r UsersRepo
}

func NewUsers(repo UsersRepo) Users {
	return Users{
		r: repo,
	}
}

func (u Users) Registration(ctx context.Context, user domain.Users) error {
	err := u.r.IsExists(ctx, user)
	if err != nil {
		return errors.WithMessage(err, "user already exists")
	}
	err = u.r.Registration(ctx, user)
	if err != nil {
		return errors.WithMessage(err, "user service")
	}
	//_, err = auth.GenerateJWT(user)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (u Users) Authentication(ctx context.Context, user domain.Users) error {
	err := u.r.IsExists(ctx, user)
	if err == nil {
		return errors.WithMessage(err, "User does not exists")
	}
	return nil
}

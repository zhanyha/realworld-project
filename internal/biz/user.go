package biz

import (
	"context"

	v1 "realworld/api/conduit/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// User is a User model.
type User struct {
	Email    string
	Password string
}

type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	ListByHello(context.Context, string) ([]*User, error)
	ListAll(context.Context) ([]*User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc UserUsecase) Login(ctx context.Context, user *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("biz.. User Login: %v", user)
	return uc.repo.Save(ctx, user)
}

func (uc UserUsecase) Register(ctx context.Context, user *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("biz.. User Register: %v", user)
	return uc.repo.Save(ctx, user)
}

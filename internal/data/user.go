package data

import (
	"context"

	"realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UserRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}

func (r *UserRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}

func (r *UserRepo) FindByID(context.Context, int64) (*biz.User, error) {
	return nil, nil
}

func (r *UserRepo) ListByHello(context.Context, string) ([]*biz.User, error) {
	return nil, nil
}

func (r *UserRepo) ListAll(context.Context) ([]*biz.User, error) {
	return nil, nil
}

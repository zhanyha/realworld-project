package service

import (
	"context"

	v1 "realworld/api/conduit/v1"
	"realworld/internal/biz"
)

type UserService struct {
	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (u *UserService) Login(ctx context.Context, reqParam *v1.LoginRequest) (*v1.UserReply, error) {
	loginUser, err := u.uc.Login(ctx, &biz.User{Email: reqParam.Email, Password: reqParam.Password})
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{Email: loginUser.Email, Token: "logintoken:" + loginUser.Password}, nil
}

func (u *UserService) Register(ctx context.Context, reqParam *v1.RegisterRequest) (*v1.UserReply, error) {
	registerUser, err := u.uc.Register(ctx, &biz.User{Email: reqParam.Email, Password: reqParam.Password})
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{Email: registerUser.Email, Token: "registertoken:" + registerUser.Password}, nil
}

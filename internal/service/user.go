package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/metadata"
	"kratos-git/helper"
	"kratos-git/models"

	pb "kratos-git/api/git"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	if meta, ok := metadata.FromServerContext(ctx); ok {
		fmt.Println("username:", meta.Get("username"))
		fmt.Println("identity:", meta.Get("identity"))
	}
	ub := models.UserBasic{}
	err := models.DB.Where("username = ? AND password = ?", req.Username, helper.GetMd5(req.Password)).Find(ub).Error
	if err != nil {
		return nil, err
	}
	//generate token
	token, err := helper.GenerateToken(ub.Identity, ub.UserName)
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		Jwt: token,
	}, nil
}

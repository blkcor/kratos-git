package service

import (
	"context"
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

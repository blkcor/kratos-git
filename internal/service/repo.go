package service

import (
	"context"
	"errors"
	"kratos-git/helper"
	"kratos-git/models"

	pb "kratos-git/api/git"
)

type RepoService struct {
	pb.UnimplementedRepoServer
}

func NewRepoService() *RepoService {
	return &RepoService{}
}

func (s *RepoService) ListRepo(ctx context.Context, req *pb.ListRepoRequest) (*pb.ListRepoReply, error) {
	rbList := make([]models.RepoBasic, 0)
	var cnt int64
	err := models.DB.Model(&models.RepoBasic{}).Count(&cnt).Offset(int((req.Page - 1) * req.Size)).Limit(int(req.Size)).Find(&rbList).Error
	if err != nil {
		return nil, err
	}
	list := make([]*pb.ListRepoItem, 0)
	for _, v := range rbList {
		list = append(list, &pb.ListRepoItem{
			Identity: v.Identity,
			Name:     v.Name,
			Desc:     v.Desc,
			Path:     v.Path,
			Star:     int32(v.Star),
		})
	}
	return &pb.ListRepoReply{
		List: list,
		Cnt:  cnt,
	}, nil
}

func (s *RepoService) CreateRepo(ctx context.Context, req *pb.CreateRepoRequest) (*pb.CreateRepoReply, error) {
	//查询仓库是否存在
	var cnt int64 = 0
	err := models.DB.Model(&models.RepoBasic{}).Where("path = ?", req.Path).Count(&cnt).Error
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("仓库已存在")
	}

	//创建仓库
	repo := models.RepoBasic{
		Identity: helper.GetUUID(),
		Name:     req.Name,
		Desc:     req.Desc,
		Path:     req.Path,
		Type:     int(req.Type),
	}
	err = models.DB.Model(&models.RepoBasic{}).Create(&repo).Error
	if err != nil {
		return nil, err
	}
	return &pb.CreateRepoReply{
		Identity: repo.Identity,
		Name:     repo.Name,
		Star:     int32(repo.Star),
		Desc:     repo.Desc,
		Path:     repo.Path,
	}, nil
}

func (s *RepoService) UpdateRepo(ctx context.Context, req *pb.UpdateRepoRequest) (*pb.UpdateRepoReply, error) {
	err := models.DB.Model(&models.RepoBasic{}).Where("identity = ?", req.Identity).Updates(map[string]interface{}{
		"name": req.Name,
		"desc": req.Desc,
		"type": req.Type,
	}).Error
	if err != nil {
		return nil, err
	}
	return &pb.UpdateRepoReply{}, nil
}

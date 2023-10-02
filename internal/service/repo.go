package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	pb "kratos-git/api/git"
	"kratos-git/define"
	"kratos-git/helper"
	"kratos-git/models"
	"os"
	"os/exec"
	"path/filepath"
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
	//创建仓库和创建文件夹放在一个事务里面
	models.DB.Transaction(func(tx *gorm.DB) error {
		err = models.DB.Model(&models.RepoBasic{}).Create(&repo).Error
		if err != nil {
			return err
		}
		//创建仓库文件夹
		gitPath := define.RepoPath + string(filepath.Separator) + req.Path
		err = os.MkdirAll(gitPath, 0755)
		if err != nil {
			return err
		}
		//git init --bare
		err = exec.Command("/bin/bash", "-c", "cd "+gitPath+";"+"git init --bare").Run()
		if err != nil {
			return err
		}
		return nil
	})
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

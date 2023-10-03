package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/metadata"
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
	md, exit := metadata.FromServerContext(ctx)
	if !exit {
		return nil, errors.New("no auth!")
	}
	identity := md.Get("identity")
	//查询用户的基本信息
	var ub models.UserBasic
	err := models.DB.Model(&models.UserBasic{}).Where("identity = ?", identity).First(&ub).Error
	if err != nil {
		return nil, err
	}

	//查询仓库是否存在
	var cnt int64 = 0
	err = models.DB.Model(&models.RepoBasic{}).Where("path = ?", req.Path).Count(&cnt).Error
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("仓库已存在")
	}

	//创建仓库
	repo := models.RepoBasic{
		Uid:      int(ub.ID),
		Identity: helper.GetUUID(),
		Name:     req.Name,
		Desc:     req.Desc,
		Path:     req.Path,
		Type:     int(req.Type),
	}
	//创建仓库和创建文件夹放在一个事务里面
	err = models.DB.Transaction(func(tx *gorm.DB) error {
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

		//信息入RepoUser表
		err = models.DB.Model(&models.RepoUser{}).Create(&models.RepoUser{
			Uid:  int(ub.ID),
			Rid:  int(repo.ID),
			Type: repo.Type,
		}).Error
		if err != nil {
			return err
		}
		return nil
	})

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

func (s *RepoService) DeleteRepo(ctx context.Context, req *pb.DeleteRepoRequest) (*pb.DeleteRepoReply, error) {
	//1、获取repo信息
	var repo = &models.RepoBasic{}
	err := models.DB.Model(&models.RepoBasic{}).Where("identity = ?", req.Identity).First(repo).Error
	if err != nil {
		return nil, err
	}
	//2、删除repo
	err = models.DB.Transaction(func(tx *gorm.DB) error {
		//1.1、删除仓库
		gitPath := define.RepoPath + string(filepath.Separator) + repo.Path
		err = os.RemoveAll(gitPath)
		if err != nil {
			return err
		}
		//1.2、删除db
		err = tx.Where("identity = ?", req.Identity).Delete(&models.RepoBasic{}).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &pb.DeleteRepoReply{}, nil
}

func (s *RepoService) RepoAuth(ctx context.Context, req *pb.RepoAuthRequest) (*pb.RepoAuthReply, error) {
	//获取当前用户信息
	md, exit := metadata.FromServerContext(ctx)
	if !exit {
		return nil, errors.New("no auth!")
	}
	var ub = &models.UserBasic{}
	err := models.DB.Model(&models.UserBasic{}).Where("identity = ?", md.Get("identity")).First(ub).Error
	if err != nil {
		return nil, err
	}
	//获取被授权用户信息
	var userAuth = &models.UserBasic{}
	err = models.DB.Model(&models.UserBasic{}).Where("identity = ?", req.UserIdentity).First(userAuth).Error
	if err != nil {
		return nil, err
	}
	//获取仓库信息
	var repo = &models.RepoBasic{}
	err = models.DB.Model(&models.RepoBasic{}).Where("identity = ?", req.RepoIdentity).First(repo).Error
	if err != nil {
		return nil, err
	}
	//查询当前用户权限
	var cnt int64
	err = models.DB.Model(&models.RepoUser{}).Where("uid = ? AND rid = ? AND type = 1", ub.ID, repo.ID).Count(&cnt).Error
	if err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, errors.New("非法操作！")
	}

	//判断被授权用户是否已经有权限
	err = models.DB.Model(&models.RepoUser{}).Where("uid = ? AND rid = ?", userAuth.ID, repo.ID).Count(&cnt).Error
	if err != nil {
		return nil, err
	}
	//有权限 直接返回
	if cnt > 0 {
		return &pb.RepoAuthReply{}, nil
	}
	//无权限 入db
	err = models.DB.Model(&models.RepoUser{}).Create(&models.RepoUser{
		Uid:  int(userAuth.ID),
		Rid:  int(repo.ID),
		Type: 2,
	}).Error
	if err != nil {
		return nil, err
	}
	return &pb.RepoAuthReply{}, nil
}

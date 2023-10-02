package models

import "gorm.io/gorm"

type RepoBasic struct {
	gorm.Model
	Uid      int    `gorm:"column:uid;type:int(11);not null;comment:用户id" json:"uid"`
	Identity string `gorm:"column:identity;type:varchar(36);not null;unique;comment:id" json:"identity"`
	Name     string `gorm:"column:name;type:varchar(255)" json:"name"`
	Desc     string `gorm:"column:desc;type:varchar(255)" json:"desc"`
	Star     int    `gorm:"column:star;type:int(11);default:0" json:"star"`
	Path     string `gorm:"column:path;type:varchar(255);not null;unique;comment:仓库路径" json:"path"`
}

func (table *RepoBasic) TableName() string {
	return "repo_basic"
}

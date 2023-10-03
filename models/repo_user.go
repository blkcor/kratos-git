package models

import "gorm.io/gorm"

type RepoUser struct {
	gorm.Model
	Uid  int `gorm:"column:uid;type:int(11);not null;comment:用户id" json:"uid"`
	Rid  int `gorm:"column:rid;type:int(11);not null;comment:仓库id" json:"rid"`
	Type int `gorm:"column:type;type:tinyint(1);not null;comment:仓库类型" json:"type"`
}

func (table *RepoUser) TableName() string {
	return "repo_user"
}

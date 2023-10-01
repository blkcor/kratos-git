package models

import "gorm.io/gorm"

type RepoStar struct {
	gorm.Model
	Uid int `gorm:"column:uid;type:int(11);not null;comment:用户id" json:"uid"`
	Rid int `gorm:"column:rid;type:int(11);not null;comment:仓库id" json:"rid"`
}

func (table *RepoStar) TableName() string {
	return "repo_star"
}

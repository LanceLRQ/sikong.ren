package models

import "gorm.io/gorm"

// 谜题
type Riddles struct {
	gorm.Model
	Author      Account `gorm:"foreignKey:AuthorId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AuthorId    uint
	Keywords    string      // 关键词，逗号分隔
	Type        int         // 题目类型
	ImageHash   string      // 图片Hash
}
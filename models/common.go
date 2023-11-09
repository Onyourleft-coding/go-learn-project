package models

import (
	"gorm.io/gorm"
	"time"
)

// 自增ID 主键

type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

// 创建、更新时间

type Timestamps struct {
	CreateAt time.Time `json:"create_at" gorm:"column:create_at;type:datetime(0);autoUpdateTime"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at;type:datetime(0);autoUpdateTime"`
}

// 软删除

type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

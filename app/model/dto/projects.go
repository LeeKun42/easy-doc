// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dto

import (
	"time"
)

const TableNameProject = "projects"

// Project mapped from table <projects>
type Project struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string    `gorm:"column:name;not null;comment:项目名称" json:"name"`
	OwnerUserID int       `gorm:"column:owner_user_id;not null;comment:项目所有者/创建人" json:"owner_user_id"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName Project's table name
func (*Project) TableName() string {
	return TableNameProject
}
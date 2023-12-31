// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dto

import (
	"time"
)

const TableNameProjectDirectory = "project_directory"

// ProjectDirectory mapped from table <project_directory>
type ProjectDirectory struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ParentID  int       `gorm:"column:parent_id;not null;comment:父级目录id" json:"parent_id"`
	ProjectID int       `gorm:"column:project_id;not null;comment:所属项目id" json:"project_id"`
	Name      string    `gorm:"column:name;comment:目录名称" json:"name"`
	Desc      string    `gorm:"column:desc;comment:目录说明" json:"desc"`
	Seq       int       `gorm:"column:seq;not null;comment:排序号从小到大排序" json:"seq"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`

	Apis 	  []ProjectAPI `gorm:"foreignKey:directory_id" json:"apis"`
}

// TableName ProjectDirectory's table name
func (*ProjectDirectory) TableName() string {
	return TableNameProjectDirectory
}

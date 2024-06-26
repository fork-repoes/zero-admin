// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysUser = "sys_user"

// SysUser 用户管理
type SysUser struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"`                          // 编号
	Name       string     `gorm:"column:name;not null;comment:用户名" json:"name"`                                          // 用户名
	NickName   *string    `gorm:"column:nick_name;comment:昵称" json:"nick_name"`                                          // 昵称
	Avatar     *string    `gorm:"column:avatar;comment:头像" json:"avatar"`                                                // 头像
	Password   string     `gorm:"column:password;not null;comment:密码" json:"password"`                                   // 密码
	Salt       string     `gorm:"column:salt;not null;comment:加密盐" json:"salt"`                                          // 加密盐
	Email      *string    `gorm:"column:email;comment:邮箱" json:"email"`                                                  // 邮箱
	Mobile     *string    `gorm:"column:mobile;comment:手机号" json:"mobile"`                                               // 手机号
	Status     int32      `gorm:"column:status;not null;comment:状态  0：禁用   1：正常" json:"status"`                          // 状态  0：禁用   1：正常
	DeptID     int64      `gorm:"column:dept_id;not null;comment:部门id" json:"dept_id"`                                   // 部门id
	JobID      int64      `gorm:"column:job_id;not null;comment:岗位id" json:"job_id"`                                     // 岗位id
	CreateBy   string     `gorm:"column:create_by;not null;comment:创建者" json:"create_by"`                                // 创建者
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateBy   *string    `gorm:"column:update_by;comment:更新者" json:"update_by"`                                         // 更新者
	UpdateTime *time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`          // 更新时间
	DelFlag    int32      `gorm:"column:del_flag;not null;default:1;comment:是否删除  0：已删除  1：正常" json:"del_flag"`          // 是否删除  0：已删除  1：正常
}

// TableName SysUser's table name
func (*SysUser) TableName() string {
	return TableNameSysUser
}

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysDict = "sys_dict"

// SysDict 字典表
type SysDict struct {
	ID          int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"`                 // 编号
	Value       string     `gorm:"column:value;not null;comment:数据值" json:"value"`                               // 数据值
	Label       string     `gorm:"column:label;not null;comment:标签名" json:"label"`                               // 标签名
	Type        string     `gorm:"column:type;not null;comment:类型" json:"type"`                                  // 类型
	Description string     `gorm:"column:description;not null;comment:描述" json:"description"`                    // 描述
	Sort        int32      `gorm:"column:sort;not null;comment:排序（升序）" json:"sort"`                              // 排序（升序）
	CreateBy    string     `gorm:"column:create_by;not null;comment:创建人" json:"create_by"`                       // 创建人
	CreateTime  *time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateBy    *string    `gorm:"column:update_by;comment:更新人" json:"update_by"`                                // 更新人
	UpdateTime  *time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
	Remarks     *string    `gorm:"column:remarks;comment:备注信息" json:"remarks"`                                   // 备注信息
	DelFlag     int32      `gorm:"column:del_flag;not null;comment:是否删除  -1：已删除  0：正常" json:"del_flag"`          // 是否删除  -1：已删除  0：正常
}

// TableName SysDict's table name
func (*SysDict) TableName() string {
	return TableNameSysDict
}

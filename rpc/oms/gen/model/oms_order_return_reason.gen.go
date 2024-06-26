// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOmsOrderReturnReason = "oms_order_return_reason"

// OmsOrderReturnReason 退货原因表
type OmsOrderReturnReason struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name       string    `gorm:"column:name;not null;comment:退货类型" json:"name"`                // 退货类型
	Sort       int32     `gorm:"column:sort;not null;comment:排序" json:"sort"`                  // 排序
	Status     int32     `gorm:"column:status;not null;comment:状态：0->不启用；1->启用" json:"status"` // 状态：0->不启用；1->启用
	CreateTime time.Time `gorm:"column:create_time;not null;comment:添加时间" json:"create_time"`  // 添加时间
}

// TableName OmsOrderReturnReason's table name
func (*OmsOrderReturnReason) TableName() string {
	return TableNameOmsOrderReturnReason
}

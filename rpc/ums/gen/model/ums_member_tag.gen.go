// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUmsMemberTag = "ums_member_tag"

// UmsMemberTag 用户标签表
type UmsMemberTag struct {
	ID                int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name              string  `gorm:"column:name;not null;comment:标签名称" json:"name"`                                      // 标签名称
	FinishOrderCount  int32   `gorm:"column:finish_order_count;not null;comment:自动打标签完成订单数量" json:"finish_order_count"`   // 自动打标签完成订单数量
	FinishOrderAmount float64 `gorm:"column:finish_order_amount;not null;comment:自动打标签完成订单金额" json:"finish_order_amount"` // 自动打标签完成订单金额
}

// TableName UmsMemberTag's table name
func (*UmsMemberTag) TableName() string {
	return TableNameUmsMemberTag
}
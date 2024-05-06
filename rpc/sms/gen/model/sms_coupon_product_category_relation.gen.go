// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSmsCouponProductCategoryRelation = "sms_coupon_product_category_relation"

// SmsCouponProductCategoryRelation 优惠券和产品分类关系表
type SmsCouponProductCategoryRelation struct {
	ID                  int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CouponID            int64  `gorm:"column:coupon_id;not null;comment:优惠券id" json:"coupon_id"`                          // 优惠券id
	ProductCategoryID   int64  `gorm:"column:product_category_id;not null;comment:产品分类id" json:"product_category_id"`     // 产品分类id
	ProductCategoryName string `gorm:"column:product_category_name;not null;comment:产品分类名称" json:"product_category_name"` // 产品分类名称
	ParentCategoryName  string `gorm:"column:parent_category_name;not null;comment:父分类名称" json:"parent_category_name"`    // 父分类名称
}

// TableName SmsCouponProductCategoryRelation's table name
func (*SmsCouponProductCategoryRelation) TableName() string {
	return TableNameSmsCouponProductCategoryRelation
}

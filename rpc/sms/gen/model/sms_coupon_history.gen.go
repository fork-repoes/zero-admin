// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSmsCouponHistory = "sms_coupon_history"

// SmsCouponHistory 优惠券使用、领取历史表
type SmsCouponHistory struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CouponID       int64     `gorm:"column:coupon_id;not null;comment:优惠券id" json:"coupon_id"`                       // 优惠券id
	MemberID       int64     `gorm:"column:member_id;not null;comment:会员id" json:"member_id"`                        // 会员id
	CouponCode     string    `gorm:"column:coupon_code;not null;comment:优惠码" json:"coupon_code"`                     // 优惠码
	MemberNickname string    `gorm:"column:member_nickname;not null;comment:领取人昵称" json:"member_nickname"`           // 领取人昵称
	GetType        int32     `gorm:"column:get_type;not null;comment:获取类型：0->后台赠送；1->主动获取" json:"get_type"`          // 获取类型：0->后台赠送；1->主动获取
	CreateTime     time.Time `gorm:"column:create_time;not null;comment:领取时间" json:"create_time"`                    // 领取时间
	UseStatus      int32     `gorm:"column:use_status;not null;comment:使用状态：0->未使用；1->已使用；2->已过期" json:"use_status"` // 使用状态：0->未使用；1->已使用；2->已过期
	UseTime        time.Time `gorm:"column:use_time;not null;comment:使用时间" json:"use_time"`                          // 使用时间
	OrderID        int64     `gorm:"column:order_id;not null;comment:订单编号" json:"order_id"`                          // 订单编号
	OrderSn        string    `gorm:"column:order_sn;not null;comment:订单号码" json:"order_sn"`                          // 订单号码
}

// TableName SmsCouponHistory's table name
func (*SmsCouponHistory) TableName() string {
	return TableNameSmsCouponHistory
}

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOmsOrder = "oms_order"

// OmsOrder 订单表
type OmsOrder struct {
	ID                    int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:订单id" json:"id"`                               // 订单id
	MemberID              int64     `gorm:"column:member_id;not null;comment:会员id" json:"member_id"`                                      // 会员id
	CouponID              int64     `gorm:"column:coupon_id;not null;comment:优惠券id" json:"coupon_id"`                                     // 优惠券id
	OrderSn               string    `gorm:"column:order_sn;not null;comment:订单编号" json:"order_sn"`                                        // 订单编号
	CreateTime            time.Time `gorm:"column:create_time;not null;comment:提交时间" json:"create_time"`                                  // 提交时间
	MemberUsername        string    `gorm:"column:member_username;not null;comment:用户帐号" json:"member_username"`                          // 用户帐号
	TotalAmount           float64   `gorm:"column:total_amount;not null;comment:订单总金额" json:"total_amount"`                               // 订单总金额
	PayAmount             float64   `gorm:"column:pay_amount;not null;comment:应付金额（实际支付金额）" json:"pay_amount"`                            // 应付金额（实际支付金额）
	FreightAmount         float64   `gorm:"column:freight_amount;not null;comment:运费金额" json:"freight_amount"`                            // 运费金额
	PromotionAmount       float64   `gorm:"column:promotion_amount;not null;comment:促销优化金额（促销价、满减、阶梯价）" json:"promotion_amount"`          // 促销优化金额（促销价、满减、阶梯价）
	IntegrationAmount     float64   `gorm:"column:integration_amount;not null;comment:积分抵扣金额" json:"integration_amount"`                  // 积分抵扣金额
	CouponAmount          float64   `gorm:"column:coupon_amount;not null;comment:优惠券抵扣金额" json:"coupon_amount"`                           // 优惠券抵扣金额
	DiscountAmount        float64   `gorm:"column:discount_amount;not null;comment:管理员后台调整订单使用的折扣金额" json:"discount_amount"`              // 管理员后台调整订单使用的折扣金额
	PayType               int32     `gorm:"column:pay_type;not null;comment:支付方式：0->未支付；1->支付宝；2->微信" json:"pay_type"`                    // 支付方式：0->未支付；1->支付宝；2->微信
	SourceType            int32     `gorm:"column:source_type;not null;comment:订单来源：0->PC订单；1->app订单" json:"source_type"`                 // 订单来源：0->PC订单；1->app订单
	Status                int32     `gorm:"column:status;not null;comment:订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单" json:"status"` // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	OrderType             int32     `gorm:"column:order_type;not null;comment:订单类型：0->正常订单；1->秒杀订单" json:"order_type"`                    // 订单类型：0->正常订单；1->秒杀订单
	DeliveryCompany       string    `gorm:"column:delivery_company;not null;comment:物流公司(配送方式)" json:"delivery_company"`                  // 物流公司(配送方式)
	DeliverySn            string    `gorm:"column:delivery_sn;not null;comment:物流单号" json:"delivery_sn"`                                  // 物流单号
	AutoConfirmDay        int32     `gorm:"column:auto_confirm_day;not null;comment:自动确认时间（天）" json:"auto_confirm_day"`                   // 自动确认时间（天）
	Integration           int32     `gorm:"column:integration;not null;comment:可以获得的积分" json:"integration"`                               // 可以获得的积分
	Growth                int32     `gorm:"column:growth;not null;comment:可以活动的成长值" json:"growth"`                                        // 可以活动的成长值
	PromotionInfo         string    `gorm:"column:promotion_info;not null;comment:活动信息" json:"promotion_info"`                            // 活动信息
	BillType              int32     `gorm:"column:bill_type;not null;comment:发票类型：0->不开发票；1->电子发票；2->纸质发票" json:"bill_type"`              // 发票类型：0->不开发票；1->电子发票；2->纸质发票
	BillHeader            string    `gorm:"column:bill_header;not null;comment:发票抬头" json:"bill_header"`                                  // 发票抬头
	BillContent           string    `gorm:"column:bill_content;not null;comment:发票内容" json:"bill_content"`                                // 发票内容
	BillReceiverPhone     string    `gorm:"column:bill_receiver_phone;not null;comment:收票人电话" json:"bill_receiver_phone"`                 // 收票人电话
	BillReceiverEmail     string    `gorm:"column:bill_receiver_email;not null;comment:收票人邮箱" json:"bill_receiver_email"`                 // 收票人邮箱
	ReceiverName          string    `gorm:"column:receiver_name;not null;comment:收货人姓名" json:"receiver_name"`                             // 收货人姓名
	ReceiverPhone         string    `gorm:"column:receiver_phone;not null;comment:收货人电话" json:"receiver_phone"`                           // 收货人电话
	ReceiverPostCode      string    `gorm:"column:receiver_post_code;not null;comment:收货人邮编" json:"receiver_post_code"`                   // 收货人邮编
	ReceiverProvince      string    `gorm:"column:receiver_province;not null;comment:省份/直辖市" json:"receiver_province"`                    // 省份/直辖市
	ReceiverCity          string    `gorm:"column:receiver_city;not null;comment:城市" json:"receiver_city"`                                // 城市
	ReceiverRegion        string    `gorm:"column:receiver_region;not null;comment:区" json:"receiver_region"`                             // 区
	ReceiverDetailAddress string    `gorm:"column:receiver_detail_address;not null;comment:详细地址" json:"receiver_detail_address"`          // 详细地址
	Note                  string    `gorm:"column:note;not null;comment:订单备注" json:"note"`                                                // 订单备注
	ConfirmStatus         int32     `gorm:"column:confirm_status;not null;comment:确认收货状态：0->未确认；1->已确认" json:"confirm_status"`            // 确认收货状态：0->未确认；1->已确认
	DeleteStatus          int32     `gorm:"column:delete_status;not null;comment:删除状态：0->未删除；1->已删除" json:"delete_status"`                // 删除状态：0->未删除；1->已删除
	UseIntegration        int32     `gorm:"column:use_integration;not null;comment:下单时使用的积分" json:"use_integration"`                      // 下单时使用的积分
	PaymentTime           time.Time `gorm:"column:payment_time;not null;comment:支付时间" json:"payment_time"`                                // 支付时间
	DeliveryTime          time.Time `gorm:"column:delivery_time;not null;comment:发货时间" json:"delivery_time"`                              // 发货时间
	ReceiveTime           time.Time `gorm:"column:receive_time;not null;comment:确认收货时间" json:"receive_time"`                              // 确认收货时间
	CommentTime           time.Time `gorm:"column:comment_time;not null;comment:评价时间" json:"comment_time"`                                // 评价时间
	ModifyTime            time.Time `gorm:"column:modify_time;not null;comment:修改时间" json:"modify_time"`                                  // 修改时间
}

// TableName OmsOrder's table name
func (*OmsOrder) TableName() string {
	return TableNameOmsOrder
}

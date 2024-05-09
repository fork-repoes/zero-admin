// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCmsPrefrenceArea = "cms_prefrence_area"

// CmsPrefrenceArea 优选专区
type CmsPrefrenceArea struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name       string `gorm:"column:name;not null" json:"name"`
	SubTitle   string `gorm:"column:sub_title;not null" json:"sub_title"`
	Pic        string `gorm:"column:pic;not null;comment:展示图片" json:"pic"` // 展示图片
	Sort       int32  `gorm:"column:sort;not null" json:"sort"`
	ShowStatus int32  `gorm:"column:show_status;not null" json:"show_status"`
}

// TableName CmsPrefrenceArea's table name
func (*CmsPrefrenceArea) TableName() string {
	return TableNameCmsPrefrenceArea
}
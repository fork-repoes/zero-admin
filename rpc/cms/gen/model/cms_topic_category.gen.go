// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCmsTopicCategory = "cms_topic_category"

// CmsTopicCategory 话题分类表
type CmsTopicCategory struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name         string `gorm:"column:name;not null" json:"name"`
	Icon         string `gorm:"column:icon;not null;comment:分类图标" json:"icon"`                   // 分类图标
	SubjectCount int32  `gorm:"column:subject_count;not null;comment:专题数量" json:"subject_count"` // 专题数量
	ShowStatus   int32  `gorm:"column:show_status;not null" json:"show_status"`
	Sort         int32  `gorm:"column:sort;not null" json:"sort"`
}

// TableName CmsTopicCategory's table name
func (*CmsTopicCategory) TableName() string {
	return TableNameCmsTopicCategory
}

package form

import (
	"time"
)

// DynamicFormIssued 表单下发
type DynamicFormIssued struct {
	ID           uint64    `gorm:"column:id;primary_key" json:"id"`
	BusinessID   int       `gorm:"column:business_id" json:"business_id"`
	ModuleID     int       `gorm:"column:module_id" json:"module_id"`
	ShowcaseID   uint64    `gorm:"column:showcase_id" json:"showcase_id"`
	FormID       uint64    `gorm:"column:form_id" json:"form_id"`			// 对应表单数据表ID
	ParentID     uint64    `gorm:"column:parent_id" json:"parent_id"`		// 对应本表父ID
	IsRequired   int8      `gorm:"column:is_required" json:"is_required"`
	SortIndex    int8      `gorm:"column:sort_index" json:"sort_index"`
	Tag          string    `gorm:"column:tag" json:"tag"`					// 对于一些有特殊处理的输入框的标签
	TableNameStr string    `gorm:"column:table_name" json:"table_name"`
	FieldName    string    `gorm:"column:field_name" json:"field_name"`
	Status       int8      `gorm:"column:status" json:"status"`
	CreateAt     time.Time `gorm:"column:create_at" json:"create_at"`
}

// TableName sets the insert table name for this struct type
func (t *DynamicFormIssued) TableName() string {
	return "t_dynamic_form_issued"
}

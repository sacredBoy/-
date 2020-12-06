package form

import (
	"time"
)

// DynamicFormConfig 动态表单配置
type DynamicFormConfig struct {
	ID           uint64    `gorm:"column:id;primary_key" json:"id"`
	Name         string    `gorm:"column:name" json:"name"`
	InputType    int       `gorm:"column:input_type" json:"input_type"`
	CheckType    int       `gorm:"column:check_type" json:"check_type"`
	SizeRange    string    `gorm:"column:size_range" json:"size_range"`
	SubOption    string    `gorm:"column:sub_option" json:"sub_option"` // 单选框选项列表
	Hint         string    `gorm:"column:hint" json:"hint"`       		// 输入提示
	ErrMsg       string    `gorm:"column:err_msg" json:"err_msg"` 		// 输入错误文案提示
	CreateAt     time.Time `gorm:"column:create_at" json:"create_at"`
}

// TableName sets the insert table name for this struct type
func (t *DynamicFormConfig) TableName() string {
	return "t_dynamic_form_config"
}

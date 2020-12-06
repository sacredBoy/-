package form

import (
	"github.com/jinzhu/gorm"
	formety "dynamic_form/models/entity/form"
)

// GetConfigByIDs 通过ID批量获取
func GetConfigByIDs(IDs []uint64) []*formety.DynamicFormConfig {
	// TODO: create DB
	db := &gorm.DB{}
	ety := make([]*formety.DynamicFormConfig, 0)
	err := db.Where("`id` in (?)", IDs).Find(&ety).Error
	if nil != err && gorm.ErrRecordNotFound != err {
		return nil
	}
	return ety
}

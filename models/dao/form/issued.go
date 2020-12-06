package form

import (
	"github.com/jinzhu/gorm"
	formety "dynamic_form/models/entity/form"
)

// GetIssuedForms 获取下发表单列表
func GetIssuedForms(businessID, moduleID, showcaseID, status int) []*formety.DynamicFormIssued {
	// TODO: create DB
	db := &gorm.DB{}
	ety := make([]*formety.DynamicFormIssued, 0)
	err := db.Where(
		"`business_id` = ? AND `module_id` = ? AND `showcase_id` = ? AND `status` = ?",
			businessID, moduleID, showcaseID, status).
		Order("sort_index").
		Find(&ety).Error
	if nil != err && gorm.ErrRecordNotFound != err {
		return nil
	}
	return ety
}

// GetIssuedFormsCount 获取下发表单配置数量
func GetIssuedFormsCount(businessID, status int, moduleShowcaseIDMap map[int]int) int {
	if len(moduleShowcaseIDMap) < 1 {
		return 0
	}
	// TODO: create DB
	db := &gorm.DB{}
	for moduleID, showcaseID := range moduleShowcaseIDMap {
		db = db.Or("`business_id` = ? AND `module_id` in (?) AND `showcase_id` = ? AND `status` = ?",
			businessID, moduleID, showcaseID, status)
	}
	var total int
	err := db.Model(&formety.DynamicFormConfig{}).Count(&total).Error
	if err != nil {
		return 0
	}
	return total
}

// GetIssuedByIDs 通过ID批量获取
func GetIssuedByIDs(IDs []uint64) []*formety.DynamicFormIssued {
	// TODO: create DB
	db := &gorm.DB{}
	ety := make([]*formety.DynamicFormIssued, 0)
	err := db.Where("`id` in (?)", IDs).Find(&ety).Error
	if nil != err && gorm.ErrRecordNotFound != err {
		return nil
	}
	return ety
}

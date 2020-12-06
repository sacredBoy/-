package formsvc

import (
	"encoding/json"
	"errors"
	"dynamic_form/constant/formdef"
	formdao "dynamic_form/models/dao/form"
	formety "dynamic_form/models/entity/form"
)

// GetFormConfig 获取表单配置
func GetFormConfig(businessID, moduleID, showcaseID, status int) ([]*FormConfigItem, error) {
	depthFormsMap := getDepthFormsMap(businessID, moduleID, showcaseID, status)
	if len(depthFormsMap) < 1 || len(depthFormsMap[0]) < 1 {
		return nil, errors.New("get forms error")
	}
	// 串联N叉树，此处map遍历忌用range，必须保证有序性
	depth := formdef.DepthTop
	for {
		nextDepthForms, ok := depthFormsMap[depth + 1]
		if !ok {
			break
		}
		parentIDFormsMap, ok := depthFormsMap[depth]
		if !ok {
			break
		}
		curDepthForms := make([]*FormConfigItem, 0)
		for _, forms := range parentIDFormsMap {
			curDepthForms = append(curDepthForms, forms...)
		}
		for _, form := range curDepthForms {
			childForms, ok := nextDepthForms[form.ID]
			if ok && len(childForms) > 0 {
				form.SubInput = childForms
				continue
			}
			form.SubInput = nil
		}
		depth++
	}
	if len(depthFormsMap[formdef.DepthTop][formdef.ParentIDTop]) < 1 {
		return nil, errors.New("build forms error")
	}
	// 返回N叉树的头指针列表
	return depthFormsMap[formdef.DepthTop][formdef.ParentIDTop], nil
}

// getDepthFormsMap 获取表单深度与表单列表map的map
func getDepthFormsMap(businessID, moduleID, showcaseID, status int) map[tDepth]map[tParentID][]*FormConfigItem {
	depthFormsMap := make(map[tDepth]map[tParentID][]*FormConfigItem)
	issuedForms, formConfigs := GetDynamicFormBaseData(businessID, moduleID, showcaseID, status)
	formIDConfigMap := make(map[uint64]*formety.DynamicFormConfig, len(formConfigs))
	for _, formConfig := range formConfigs {
		formIDConfigMap[formConfig.ID] = formConfig
	}
	lastDepthIDFlag := map[uint64]bool{formdef.ParentIDTop:true}
	for depth := formdef.DepthTop; len(lastDepthIDFlag) > 0; depth++ {
		currentDepthIDFlag := make(map[uint64]bool)
		pIDFormsMap := make(map[tParentID][]*FormConfigItem)
		for i, issuedForm := range issuedForms {
			if issuedForm == nil {
				continue
			}
			_, ok := lastDepthIDFlag[issuedForm.ParentID]
			if ok {
				pIDFormsMap[issuedForm.ParentID] = append(
					pIDFormsMap[issuedForm.ParentID],
					NewFormConfigItem(issuedForm, formIDConfigMap[issuedForm.FormID]))
				currentDepthIDFlag[issuedForm.ID] = true
				// 剪枝，减少后续无意义遍历
				issuedForms[i] = nil
			}
		}
		depthFormsMap[depth] = pIDFormsMap
		lastDepthIDFlag = currentDepthIDFlag
	}
	return depthFormsMap
}

// GetDynamicFormBaseData 获取动态表单基础数据
func GetDynamicFormBaseData(businessID, moduleID, showcaseID, status int) (
	[]*formety.DynamicFormIssued, []*formety.DynamicFormConfig) {
	issuedForms := formdao.GetIssuedForms(businessID, moduleID, showcaseID, status)
	issuedFormsLen := len(issuedForms)
	if issuedFormsLen < 1 {
		return nil, nil
	}
	formIDs := make([]uint64, issuedFormsLen)
	for i, issuedForm := range issuedForms {
		formIDs[i] = issuedForm.FormID
	}
	formConfigs := formdao.GetConfigByIDs(formIDs)
	if len(formConfigs) != issuedFormsLen {
		return nil, nil
	}
	return issuedForms, formConfigs
}

// GetFormConfigItemByIDs 通过下发ID(issued中的ID)获取表单列表信息(不串联树)
func GetFormConfigItemByIDs(IDs []uint64) []*FormConfigItem {
	issuedForms := formdao.GetIssuedByIDs(IDs)
	issuedFormsLen := len(issuedForms)
	if issuedFormsLen < 1 {
		return nil
	}
	formIDs := make([]uint64, issuedFormsLen)
	for i, issuedForm := range issuedForms {
		formIDs[i] = issuedForm.FormID
	}
	formConfigs := formdao.GetConfigByIDs(formIDs)
	if len(formConfigs) != issuedFormsLen {
		return nil
	}
	formIDConfigMap := make(map[uint64]*formety.DynamicFormConfig, len(formConfigs))
	for _, formConfig := range formConfigs {
		formIDConfigMap[formConfig.ID] = formConfig
	}
	formConfigItems := make([]*FormConfigItem, issuedFormsLen)
	for i, issuedForm := range issuedForms {
		formConfigItems[i] = NewFormConfigItem(issuedForm, formIDConfigMap[issuedForm.FormID])
	}
	return formConfigItems
}

// NewFormConfigItem 构造表单服务所需结构体数据
func NewFormConfigItem(issued *formety.DynamicFormIssued, config *formety.DynamicFormConfig) *FormConfigItem {
	return &FormConfigItem{
		ID:         issued.ID,
		Name:       config.Name,
		InputType:  config.InputType,
		IsRequired: issued.IsRequired,
		TableName:  issued.TableNameStr,
		FieldName:  issued.FieldName,
		Tag:        issued.Tag,
		CheckType:  config.CheckType,
		Hint:       config.Hint,
		ErrMsg:     config.ErrMsg,
		SizeRange:  FormatSizeRange(config.SizeRange),
		SubOption:  FormatSubOption(config.SubOption),
		SubCountry: FormatSubCountry(config.InputType),
	}
}

// FormatSizeRange 格式化数据存储为具体结构
func FormatSizeRange(sizeRangeStr string) *SizeRange {
	sizeRange := &SizeRange{}
	err := json.Unmarshal([]byte(sizeRangeStr), sizeRange)
	if err != nil {
		return nil
	}
	return sizeRange
}

// FormatSubOption 格式化数据存储为具体结构
func FormatSubOption(subOptionStr string) []*SubOption {
	subOption := make([]*SubOption, 10)
	err := json.Unmarshal([]byte(subOptionStr), &subOption)
	if err != nil {
		return nil
	}
	return subOption
}

// FormatSubCountry 格式化数据存储为具体结构
func FormatSubCountry(inputType int) *SubCountry {
	if inputType != InputTypeCountry {
		return nil
	}
	return &SubCountry{
		HotCountries: HotCountries,
		Countries:    Countries,
	}
}

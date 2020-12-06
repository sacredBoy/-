package formsvc

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
)


// Bfs 为N叉树提供广度优先遍历
func (f *FormConfigItem) Bfs(op func(item *FormConfigItem)) {
	if f == nil {
		return
	}
	queue := FormConfigItems{f}
	for queue.size() > 0 {
		size := queue.size()
		for i := 0; i < size; i++ {
			node := queue.remove()
			op(node)
			queue.addAll(node.SubInput)
		}
	}
}

// add 为支持广度优先搜索的队列操作
func (f *FormConfigItems) add(node *FormConfigItem) bool {
	*f = append(*f, node)
	return true
}

// remove 为支持广度优先搜索的队列操作
func (f *FormConfigItems) remove() *FormConfigItem {
	if len(*f) == 0 {
		return nil
	}
	first := []*FormConfigItem(*f)[0]
	*f = []*FormConfigItem(*f)[1:]
	return first
}

// element 为支持广度优先搜索的队列操作
func (f *FormConfigItems) element() *FormConfigItem {
	first := []*FormConfigItem(*f)[0]
	return first
}

// size 为支持广度优先搜索的队列操作
func (f *FormConfigItems) size() int {
	return len(*f)
}

// addAll 为支持广度优先搜索的队列操作
func (f *FormConfigItems) addAll(nodes []*FormConfigItem) bool {
	*f = append(*f, nodes...)
	return true
}

// CheckSubmitForms 验证表单提交
func CheckSubmitForms(formConfigItems []*FormConfigItem, submitFormIDDataMap map[uint64]string) error {
	for _, form := range formConfigItems {
		data, ok := submitFormIDDataMap[form.ID]
		dataLen := len(data)
		if !ok || (form.IsRequired == RequiredTrue && form.InputType != InputTypeChiL && dataLen < 1) {
			return errors.New("missing params")
		}
		if dataLen < 1 {
			continue
		}
		switch form.InputType {
		case InputTypeText, InputTypeNum, InputTypeLong, InputTypeImage, InputTypeCountry:
			if err := checkSubmitType(data, form.CheckType); err != nil {
				return err
			}
		case InputTypeSdEmail:
			if err := checkVerifyData(data, submitFormIDDataMap, InputTypeSdEmail, formConfigItems); err != nil {
				return err
			}
		// 下发给客户端的checkType与服务端实际的checkType不一样，单独处理
		case InputTypeDate, InputTypeSin:
			if err := checkSubmitType(data, CheckTypeNum); err != nil {
				return err
			}
		case InputTypeMul:
			IDs := make([]int, 0)
			err := json.Unmarshal([]byte(data), &IDs)
			if err != nil {
				return errors.New("params error")
			}
		}
	}
	return nil
}

// checkSubmitType 校验提交数据类型是否正确
func checkSubmitType(data string, checkType int) error {
	switch checkType {
	case CheckTypeNull:
		return nil
	case CheckTypeNum:
		_, err := strconv.Atoi(data)
		if err != nil {
			return errors.New("need pure numbers")
		}
	case CheckTypeEmail:
		isMatch, _ := regexp.MatchString(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`, data)
		if !isMatch {
			return errors.New("please enter a valid email address. ")
		}
	case CheckTypeLink:
		isMatch, _ := regexp.MatchString(`http(s)?.*`, data)
		if !isMatch {
			return errors.New("the link you filled in is incorrect, please check and refill. ")
		}
	default:
		return nil
	}
	return nil
}

// checkVerifyData 检查校验类型数据，该类数据不会保存至数据库，如邮件验证码
func checkVerifyData(
	data string,
	submitFormIDDataMap map[uint64]string,
	inputType int,
	formConfigItems []*FormConfigItem) error {
	switch inputType {
	case InputTypeSdEmail:
		if err := checkSubmitType(data, CheckTypeEmail); err != nil {
			return err
		}
		fieldName := inputTypeCheckFieldNameMap[InputTypeSdEmail]
		for _, form := range formConfigItems {
			if form.FieldName != fieldName {
				continue
			}
			// TODO: check email code
		}
	}
	return nil
}

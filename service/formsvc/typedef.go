package formsvc

// tDepth 表单深度
type tDepth = int
// tParentID 表单父级ID
type tParentID = uint64

// check_type字段：校验用户输入类型(新类型从最后一行追加)
const (
	CheckTypeNull  = iota // 无
	CheckTypeNum          // 纯数字
	CheckTypeEmail        // 邮箱
	CheckTypeLink         // 链接
	CheckTypePhone		  // 电话号码
)

// input_type字段：用户输入类型(新类型从最后一行追加，客户端传值统一字符串)
const (
	InputTypeText    = iota // 文本类型(客户端传值后处理为string)
	InputTypeNum            // 纯数字(客户端传值后处理为int64)
	InputTypeLong           // 长文本，客户端需要跳转至长文本编辑界面(客户端传值后处理为string)
	InputTypeImage          // 图片，客户端需要做上传图片交互(客户端传图片url值后处理为string)
	InputTypeDate           // 日期，客户端需要做日期选择(客户端传值时间戳后处理为int64)
	InputTypeSin            // 单选框，客户端需要做下拉框选择(客户端传选项id值后处理为int8)
	InputTypeMul            // 复选框，客户端需要做下拉框选择(客户端传选项id列表后处理为int8[])
	InputTypeChiL           // 子集输入框，客户端需要跳转再做表单处理(客户端传值：按照以上类型处理)
	InputTypeSdEmail        // 发送邮件，用于验证码，客户端需要请求邮件发送接口(客户端传值后处理为string)
	InputTypeCountry        // 输入国家，会下发不同于单选框的选项列表(客户端传值后处理为string)
)

// is_required字段：是否必填
const (
	RequiredTrue  = 1 // 必填
	RequiredFalse = 0 // 非必填
)

// SizeRange 大小范围
type SizeRange struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}

// SubOption 子选项
type SubOption struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// SubCountry 子国家选项
type SubCountry struct {
	HotCountries []string `json:"hot_countries"`
	Countries    []string `json:"countries"`
}

// FormConfigItem 结果元素
type FormConfigItem struct {
	ID         uint64 			 `json:"id"`
	Name       string 			 `json:"name"`
	InputType  int   			 `json:"input_type"`
	IsRequired int8   			 `json:"is_required"`
	// 返回时为nil，业务所需时自行填充
	Default    string 			 `json:"default"`
	// 该输入框最终最终添加到数据表中对应的字段名
	TableName  string			 `json:"table_name,omitempty"`
	FieldName  string            `json:"field_name,omitempty"`
	Tag 	   string			 `json:"tag,omitempty"`
	CheckType  int               `json:"check_type"`
	SizeRange  *SizeRange        `json:"size_range"`
	Hint       string            `json:"hint"`
	ErrMsg     string            `json:"err_msg"`
	SubOption  []*SubOption      `json:"sub_option"`
	SubCountry *SubCountry       `json:"sub_country"`
	SubInput   []*FormConfigItem `json:"sub_input"`
}

// FormConfigItems 结果
type FormConfigItems []*FormConfigItem

// SubmitFormItem 表单提交元素
type SubmitFormItem struct {
	ID   uint64	`json:"id"`
	Data string	`json:"data"`
}

// SubmitFormItems 表单提交数据
type SubmitFormItems []*SubmitFormItem

var inputTypeCheckFieldNameMap = map[int]string{
	InputTypeSdEmail: "_check_code",
}

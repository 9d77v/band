package base

type FuzzyField struct {
	Name  string      `json:"name"`  //查询字段名称
	Value interface{} `json:"value"` //查询字段值
}

type TermField struct {
	Name  string        `json:"name"`  //查询字段名称
	Value []interface{} `json:"value"` //查询字段值
}

type LTField struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type LTEField struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type GTField struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type GTEField struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type JSONFuzzyField struct {
	Field string      `json:"field"` //json字段名
	Name  string      `json:"name"`  //查询字段名称
	Value interface{} `json:"value"` //查询字段值
}

type JSONTermField struct {
	Field string        `json:"field"` //json字段名
	Name  string        `json:"name"`  //查询字段名称
	Value []interface{} `json:"value"` //查询字段值
}

type JSONLTField struct {
	Field string  `json:"field"` //json字段名
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type JSONLTEField struct {
	Field string  `json:"field"` //json字段名
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type JSONGTField struct {
	Field string  `json:"field"` //json字段名
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type JSONGTEField struct {
	Field string  `json:"field"` //json字段名
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type JSONField struct {
	TermFields  []JSONTermField
	FuzzyFields []JSONFuzzyField
	LTEFields   []JSONLTEField
	LTFields    []JSONLTField
	GTEFields   []JSONGTEField
	GTFields    []JSONGTField
	ArrayFields []TermField
}

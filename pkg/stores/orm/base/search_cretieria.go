package base

import (
	"context"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type SearchField struct {
	TermFields  []TermField
	FuzzyFields []FuzzyField
	LTEFields   []LTEField
	LTFields    []LTField
	GTEFields   []GTEField
	GTFields    []GTField
	JSONField   JSONField
}

// SearchCriteria 搜索条件
type SearchCriteria struct {
	PageQuery
	Columns string
	Table   string
	SearchField
	Joins       []string
	Order       string
	Group       string
	OrFuzzys    []OrFuzzyField
	NoPageQuery bool
}

type OrFuzzyField struct {
	Names []string    `json:"names"` //查询字段名称
	Value interface{} `json:"value"` //查询字段值
}

func NewSearchCriteriaOfPageQuery(query PageQuery) *SearchCriteria {
	return &SearchCriteria{PageQuery: query}
}

func NewJoinSearchCriteria(columns, table string, joins []string) *SearchCriteria {
	return &SearchCriteria{
		Columns: columns,
		Table:   table,
		Joins:   joins,
	}
}

func NewJoinSearchCriteriaOfPageQuery(query PageQuery, columns, table string, joins []string) *SearchCriteria {
	return &SearchCriteria{
		PageQuery: query,
		Columns:   columns,
		Table:     table,
		Joins:     joins,
	}
}

// 构建interface数组
func BuildArray[T any](data []T) []interface{} {
	values := []interface{}{}
	for _, v := range data {
		values = append(values, v)
	}
	return values
}

// 关键词查询
func (sc *SearchCriteria) WithKeyword(names []string, keyword string) *SearchCriteria {
	if keyword != "" {
		sc.OrFuzzys = append(sc.OrFuzzys, OrFuzzyField{
			Names: names,
			Value: keyword,
		})
	}
	return sc
}

func (sc *SearchCriteria) BuildDB(db *gorm.DB) *gorm.DB {
	if sc.Columns != "" {
		db = db.Select(sc.Columns)
	}
	if sc.Table != "" {
		db = db.Table(sc.Table)
	}
	if sc.Order != "" {
		db = db.Order(sc.Order)
	}
	if sc.Group != "" {
		db = db.Group(sc.Group)
	}
	for _, v := range sc.Joins {
		db = db.Joins(v)
	}
	for _, v := range sc.TermFields {
		if len(v.Value) == 1 {
			if v.Value[0] == nil {
				db = db.Where(fmt.Sprintf("%s is null", v.Name))
			} else {
				db = db.Where(fmt.Sprintf("%s = ?", v.Name), v.Value[0])
			}
		} else if len(v.Value) > 1 {
			db = db.Where(fmt.Sprintf("%s in ?", v.Name), v.Value)
		}
	}
	for _, v := range sc.LTEFields {
		db = db.Where(fmt.Sprintf("%s <= ?", v.Name), v.Value)
	}
	for _, v := range sc.LTFields {
		db = db.Where(fmt.Sprintf("%s < ?", v.Name), v.Value)
	}
	for _, v := range sc.GTEFields {
		db = db.Where(fmt.Sprintf("%s >= ?", v.Name), v.Value)
	}
	for _, v := range sc.GTFields {
		db = db.Where(fmt.Sprintf("%s > ?", v.Name), v.Value)
	}
	for _, v := range sc.FuzzyFields {
		db = db.Where(fmt.Sprintf("%s LIKE ?", v.Name), "%"+fmt.Sprintf("%v", v.Value)+"%")
	}
	for _, v := range sc.JSONField.TermFields {
		if len(v.Value) == 1 {
			db = db.Where(fmt.Sprintf("%s->'$.%s' = ?", v.Field, v.Name), v.Value[0])
		} else {
			db = db.Where(fmt.Sprintf("%s->'$.%s' in ?", v.Field, v.Name), v.Value)
		}
	}
	for _, fuzzy := range sc.OrFuzzys {
		orArr := []string{}
		values := []interface{}{}
		for _, v := range fuzzy.Names {
			orArr = append(orArr, " "+v+" like ? ")
			values = append(values, "%"+fuzzy.Value.(string)+"%")
		}
		db = db.Where(strings.Join(orArr, "or"), values...)
	}
	return db
}

// RawSearchCriteria 原生查询
type RawSearchCriteria struct {
	SelectSql string
	CountSql  string
	Where     []string
	OrderSql  string
	Values    []interface{}
	PageQuery
}

// 封装分页列表查询操作
func Page[R any, T any](ctx context.Context, db *gorm.DB, t *T, query SearchCriteria) (data []*R, total int64, err error) {
	if t != nil {
		db = db.Model(t)
	}
	db = query.BuildDB(db)
	data = []*R{}
	if query.NoPageQuery {
		err = db.Find(&data).Error
		return data, 0, err
	}
	err = db.Count(&total).Error
	if err != nil || total == 0 {
		return data, 0, err
	}
	if query.Limit != -1 {
		db = db.Offset(query.Offset()).Limit(query.Limit)
	}
	err = db.Find(&data).Error
	return data, int64(total), err
}

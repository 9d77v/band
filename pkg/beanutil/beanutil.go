package beanutil

import (
	"errors"
	"fmt"
	"reflect"
)

func ConvertObject(dst, src interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)

	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}

	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return errors.New("src type should be a struct or a struct pointer")
	}

	dstType, dstValue = dstType.Elem(), dstValue.Elem()

	propertyNums := dstType.NumField()

	for i := 0; i < propertyNums; i++ {
		property := dstType.Field(i)
		propertyValue := srcValue.FieldByName(property.Name)
		if !propertyValue.IsValid() || property.Type != propertyValue.Type() {
			continue
		}
		if dstValue.Field(i).CanSet() {
			dstValue.Field(i).Set(propertyValue)
		}
	}
	return nil
}

func ConvertList[T, S any](soure []*S) ([]*T, error) {
	resps := []*T{}
	for _, v := range soure {
		resp := new(T)
		err := ConvertObject(resp, v)
		if err != nil {
			return resps, err
		}
		resps = append(resps, resp)
	}
	return resps, nil
}

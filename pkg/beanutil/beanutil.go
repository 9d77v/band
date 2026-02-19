package beanutil

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"maps"
	"reflect"
)

func ConvertObject(dst, src any) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)

	if dstType.Kind() != reflect.Pointer || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}

	if srcType.Kind() == reflect.Pointer {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return errors.New("src type should be a struct or a struct pointer")
	}

	dstType, dstValue = dstType.Elem(), dstValue.Elem()

	propertyNums := dstType.NumField()

	for i := range propertyNums {
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

func ConvertToMap(content any) map[string]any {
	var name map[string]any
	if marshalContent, err := json.Marshal(content); err != nil {
		log.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber()
		if err := d.Decode(&name); err != nil {
			log.Println(err)
		} else {
			maps.Copy(name, name)
		}
	}
	return name
}

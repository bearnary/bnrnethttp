package bnrnethttp

import (
	"fmt"
	"net/url"
	"reflect"
	"time"
)

func (c *defaultClient) CreateQueryParamString(model interface{}) string {
	t := reflect.TypeOf(model).Elem()
	v := reflect.ValueOf(model).Elem()

	params := url.Values{}
	timeType := reflect.TypeOf(time.Time{})

	for i := 0; i < v.NumField(); i++ {

		valueField := v.Field(i)

		if valueField.Kind() == reflect.Struct {
			it := valueField.Type()

			for j := 0; j < it.NumField(); j++ {
				inboundTypeField := it.Field(j)
				inboundValueField := valueField.Field(j)
				if inboundValueField.IsNil() {
					continue
				}

				tag := inboundTypeField.Tag.Get("form")
				switch inboundValueField.Kind() {
				case reflect.Slice:
					panic("not support slice in query param")
				default:
					value := inboundValueField.Elem()
					param := fmt.Sprintf("%v", value)
					params.Add(tag, param)
				}
			}
			continue
		} else if valueField.IsNil() {
			continue
		}

		tag := t.Field(i).Tag.Get("form")
		switch v.Field(i).Kind() {
		case reflect.Slice:
			panic("not support slice in query param")
		default:
			value := v.Field(i).Elem()
			param := fmt.Sprintf("%v", value)
			if value.Type() == timeType {
				timeValue := value.Interface().(time.Time)
				param = fmt.Sprintf("%v", timeValue.Format(time.RFC3339))
			}
			params.Add(tag, param)
		}

	}

	return params.Encode()

}

package my_cast

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
)

func ToMapByFieldE(key string, data interface{}) (map[string]interface{}, error) {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Slice {
		result := make(map[string]interface{}, val.Len())

		for i := 0; i < val.Len(); i++ {
			e := val.Index(i).Elem()
			Id := reflect.Indirect(e).FieldByName(key) // reflect.Indirect(e).FieldByName(key)

			if !Id.IsValid() {
				return nil, fmt.Errorf("interface `%s` does not have the field `%s`", val.Type(), key)
			}
			var (
				mapKey string
				err    error
			)

			if mapKey, err = cast.ToStringE(Id.Interface()); err != nil {
				return nil, err
			}

			var mapVal interface{}
			if e.Kind() == reflect.Ptr {
				mapVal = e.Elem().Addr().Interface()
			} else {
				mapVal = e.Interface()
			}
			result[mapKey] = mapVal
		}
		return result, nil
	}
	return nil, fmt.Errorf("incoming data is not slice")
}

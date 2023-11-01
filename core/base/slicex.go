package base

import "reflect"

// InArrayMixed 判断值是否在另一个切片中，支持多类型, 需注意值类型与数组类型要一致
func InArrayMixed(val any, array any) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

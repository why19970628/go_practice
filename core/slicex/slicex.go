package slicex

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

// ChunkSlice 切片分块
func ChunkSlice[T any](slice []T, chunkSize int) [][]T {
	var result [][]T

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}

	return result
}

func RemoveElement[T int | int8 | int32 | int16 | int64 | string](slice []T, element T) []T {
	if len(slice) == 0 {
		return []T{}
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == element {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}
	return slice

}

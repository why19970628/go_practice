package slicex

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	type args[T interface {
		int | int8 | int32 | int16 | int64
	}] struct {
		slice   []T
		element T
	}
	type testCase[T interface {
		int | int8 | int32 | int16 | int64
	}] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int64]{
		{
			name: "64",
			args: args[int64]{[]int64{0, 1, 2, 3, 0}, 0},
			want: []int64{1, 2, 3},
		},
		{
			name: "64_2",
			args: args[int64]{[]int64{}, 0},
			want: []int64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElement(tt.args.slice, tt.args.element); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

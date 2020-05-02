package search

import (
	"reflect"
	"testing"
)

func Test_getLeastNumbers(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{
			name: "one",
			args: args{
				nums: []int{3, 2, 1},
				k:    2,
			},
			want: []int{1, 2},
		},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLeastNumbers(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSortTopK() = %v, want %v", got, tt.want)
			}
		})
	}
}

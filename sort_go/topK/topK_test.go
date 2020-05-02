package topk

import (
	"reflect"
	"testing"
)

func TestTopK(t *testing.T) {
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
			name: "topk",
			args: args{
				nums: []int{3, 4, 5, 6, 7, 9, 7, 5, 4, 3, 2, 1, 2, 3, 4, 5, 6, 0, 8, 9, 9, 54, 2, 435, 11, 234, 324, 12, 43, 78, 5657, 66, 87, 867, 65, 324, 87, 23, 213, 12, 89, 23, 12, 4, 45, 434, 6, -43, 656, 4, 323, 43, 323},
				k:    5,
			},
			want: []int{5657, 867, 656, 324, 435},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopK(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopK() = %v, want %v", got, tt.want)
			}
		})
	}
}

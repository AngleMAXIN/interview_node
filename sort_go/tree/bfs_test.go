package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	type args struct {
		tree map[string][]string
		head string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "A is head",
			args: args{
				head: "A",
				tree: map[string][]string{
					"A": []string{"B", "C"},
					"B": []string{"A", "C", "D"},
					"C": []string{"A", "B", "D", "E"},
					"D": []string{"B", "C", "E", "F"},
					"E": []string{"C", "D"},
					"F": []string{"D"},
				},
			},
			want: []string{"A", "B", "C", "D", "E", "F"},
		},
		{
			name: "E is head",
			args: args{
				head: "E",
				tree: map[string][]string{
					"A": []string{"B", "C"},
					"B": []string{"A", "C", "D"},
					"C": []string{"A", "B", "D", "E"},
					"D": []string{"B", "C", "E", "F"},
					"E": []string{"C", "D"},
					"F": []string{"D"},
				},
			},
			want: []string{"E", "C", "D", "A", "B", "F"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BFS(tt.args.tree, tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSliceAppend(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := []int{4, 5, 6}
		s = s[1:]
		if len(s) != 2 {
			fmt.Println(len(s), s)
		}
	}
}

func BenchmarkSliceCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := []int{4, 5, 6}
		copy(s, s[1:])
		if len(s) != 2 {
			fmt.Println(len(s), s)
		}
	}
}

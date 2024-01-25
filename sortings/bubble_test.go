package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	data := randomData(100, 1000)
	sorted := append([]int{}, data()...)
	sort.Ints(sorted)

	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sort reversed slice",
			args: args{data: []int{55, 34, 21, 13, 8, 5, 3, 2, 1, 1}},
			want: []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55},
		},
		{
			name: "sort reversed slice",
			args: args{
				data: data(),
			},
			want: sorted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	//data := []int{55, 34, 21, 13, 8, 5, 3, 2, 1, 1}
	data := generateData(b.N, 1000, 5000)
	//data := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(data[i])
	}
}

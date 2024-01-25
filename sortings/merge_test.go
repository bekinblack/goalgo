package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	data := randomData(100, 1000)
	sorted := append([]int{}, data()...)
	sort.Ints(sorted)

	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "sort reversed slice",
			arr:  []int{55, 34, 21, 13, 8, 5, 3, 2, 1, 1},
			want: []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55},
		},
		{
			name: "sort random slice",
			arr:  data(),
			want: sorted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	data := generateData(DataSet, DataLength, 5000)
	b.ResetTimer()
	for i := 0; i < DataSet; i++ {
		mergeSort(data[i])
	}
}

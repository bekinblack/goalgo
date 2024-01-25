package main

import (
	"reflect"
	"testing"
)

func Test_selectionSort(t *testing.T) {
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
			name: "sort random slice",
			args: args{data: []int{5, 2, 3, 1, 6, 7, 8, 4, 9, 0}},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectiomSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	data := generateData(DataSet, DataLength, 5000)
	b.ResetTimer()
	for i := 0; i < DataSet; i++ {
		SelectionSort(data[i])
	}
}

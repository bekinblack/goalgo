package main

import (
	"reflect"
	"sort"
	"testing"
)

//func Test_QuickSortWithTwoSlices(t *testing.T) {
//	data := randomData(100, 1000)
//	sorted := append([]int{}, data()...)
//	sort.Ints(sorted)
//
//	tests := []struct {
//		name string
//		arr  []int
//		want []int
//	}{
//		{
//			name: "sort reversed slice",
//			arr:  []int{55, 34, 21, 13, 8, 5, 3, 2, 1, 1},
//			want: []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55},
//		},
//		{
//			name: "sort random slice",
//			arr:  data(),
//			want: sorted,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got := QuickSortWithTwoSlices(tt.arr)
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("QuickSortWithTwoSlices() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func Benchmark_quickSort1(b *testing.B) {
//	data := generateData(b.N, 1000, 5000)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		QuickSortWithTwoSlices(data[i])
//	}
//}

func Test_QuickSortWithPartition(t *testing.T) {
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
			name: "sort random slice",
			args: args{data: data()},
			want: sorted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := QuickSortWithPartition(tt.args.data, 0, len(tt.args.data)-1)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectiomSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Benchmark_QuickSortWithPartition(b *testing.B) {
	data := generateData(DataSet, DataLength, 5000)
	b.ResetTimer()
	for i := 0; i < DataSet; i++ {
		QuickSortWithPartition(data[i], 0, 9)
	}
}

//
//func Test_QuickSortWithThreeSlice(t *testing.T) {
//	data := randomData(100, 1000)
//	sorted := append([]int{}, data()...)
//	sort.Ints(sorted)
//
//	type args struct {
//		data []int
//	}
//	tests := []struct {
//		name string
//		args args
//		want []int
//	}{
//		{
//			name: "sort reversed slice",
//			args: args{data: []int{55, 34, 21, 13, 8, 5, 3, 2, 1, 1}},
//			want: []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55},
//		},
//		{
//			name: "sort random slice",
//			args: args{data: data()},
//			want: sorted,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := QuickSortWithThreeSlices(tt.args.data); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("SelectiomSort() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func BenchmarkQuickSortWithThreeSlice(b *testing.B) {
//	data := generateData(b.N, 10, 5000)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		QuickSortWithThreeSlices(data[i])
//	}
//}

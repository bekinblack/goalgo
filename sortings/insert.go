package main

func InsertSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {

		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}

	return arr
}

func InsertSort2(arr []int) []int {
	n := len(arr)
	for i := 1; i <= n; i++ {

		j := i - 1
		for j > 0 && arr[j] < arr[j-1] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j--
		}
	}
	return arr
}

package main

func QuickSortWithThreeSlices(arr []int) []int {
	n := len(arr)

	if n <= 1 {
		return arr
	}

	pivot := arr[n/2]
	low := make([]int, 0, n)
	middle := make([]int, 0, n)
	high := make([]int, 0, n)

	for _, item := range arr {
		switch {
		case item < pivot:
			low = append(low, item)
		case item == pivot:
			middle = append(middle, item)
		case item > pivot:
			high = append(high, item)
		}
	}

	low = QuickSortWithThreeSlices(low)
	high = QuickSortWithThreeSlices(high)
	low = append(low, middle...)
	return append(low, high...)
}

// QuickSortWithPartition with partition func
func QuickSortWithPartition(arr []int, low, high int) []int {
	if low < high {
		pivot := partition(arr, low, high)
		QuickSortWithPartition(arr, low, pivot-1)
		QuickSortWithPartition(arr, pivot+1, high)
	}

	return arr
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func QuickSortWithTwoSlices(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)-1]
	var left, right []int

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = QuickSortWithTwoSlices(left)
	right = QuickSortWithTwoSlices(right)

	return append(append(left, pivot), right...)
}

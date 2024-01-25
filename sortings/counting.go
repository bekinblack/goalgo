package main

func countingSort(arr []int) {
	max := findMax(arr)
	count := make([]int, max+1)
	sortedArr := make([]int, len(arr))

	for _, num := range arr {
		count[num]++
	}

	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		num := arr[i]
		sortedArr[count[num]-1] = num
		count[num]--
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = sortedArr[i]
	}
}

func findMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

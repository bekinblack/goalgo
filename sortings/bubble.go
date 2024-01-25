package main

func BubbleSort(data []int) []int {
	n := len(data)

	swapped := true
	i := 1
	for swapped {
		swapped = false
		for j := 0; j < n-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				swapped = true
			}
		}
		i++
	}

	return data
}

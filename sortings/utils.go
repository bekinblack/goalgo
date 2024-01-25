package main

import (
	"math/rand"
	"time"
)

const DataLength = 5000
const DataSet = 1000

func randomData(n, max int) func() []int {
	var data []int
	return func() []int {
		if data != nil {
			return data
		}

		rand.Seed(time.Now().UnixNano())
		for i := 0; i < n; i++ {
			data = append(data, rand.Intn(max))
		}

		return data
	}
}

func generateData(n, count, maxValue int) [][]int {
	var data [][]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		var cycleData []int
		for j := 0; j < count; j++ {
			cycleData = append(cycleData, rand.Intn(maxValue))
		}
		data = append(data, cycleData)
	}

	return data
}

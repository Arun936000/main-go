package main

import (
	"fmt"
	"math"
)

func secondlargest(arr []int) (int, error) {
	if len(arr) < 2 {
		return 0, fmt.Errorf("Array can be of min two")
	}
	large, secondlar := math.MinInt64, math.MinInt64
	for _, num := range arr {
		if num > large {
			secondlar = large
			large = num
		} else if num > secondlar && num != large {
			secondlar = num
		}
	}
	if secondlar == math.MinInt64 {
		return 0, fmt.Errorf("No second large element is found")
	}
	return secondlar, nil
}
func main() {
	arr := []int{10, 40, 50, 20, 60}
	secondlar, err := secondlargest(arr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The second largest element is:", secondlar)
	}
}

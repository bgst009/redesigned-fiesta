package VariadicFunctions

import "fmt"

func max(values ...int) int {

	shouldReturn, returnValue := isEmpty(values)
	if shouldReturn {
		return returnValue
	}

	max := values[0]
	for _, val := range values {
		if max < val {
			max = val
		}
	}
	return max
}

func isEmpty(values []int) (bool, int) {
	if len(values) == 0 {
		fmt.Println("require at least one arguement")
		return true, 0
	}
	return false, 0
}

func min(values ...int) int {

	shouldReturn, returnValue := isEmpty(values)
	if shouldReturn {
		return returnValue
	}

	min := values[0]
	for _, val := range values {
		if min > val {
			min = val
		}
	}
	return min
}

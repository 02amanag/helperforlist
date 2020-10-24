package learn

import "math"

func Nearest(value float64, list []float64) float64 {
	smallestDistance := math.Abs(list[0] - value)
	number := list[0]
	for i, _ := range list {
		newDistance := math.Abs(list[i] - value)
		if newDistance < smallestDistance {
			smallestDistance = newDistance
			number = list[i]
		}
	}
	return number
}

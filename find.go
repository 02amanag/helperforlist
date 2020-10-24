package learn

import (
	"errors"
)

func Find(value float64, list []float64) error {
	for _, i := range list {
		if i == value {
			return nil
		}
	}
	return errors.New("not found")
}

func IfFindRemove(value float64, list []float64) []float64 {
	var newlist []float64
	for _, i := range list {
		if i != value {
			newlist = append(newlist, i)
		}
	}
	return newlist
}

func IfNotFindAdd(value float64, list []float64) []float64 {
	var newlist []float64
	for _, i := range list {
		if i == value {
			return list
		}
	}
	newlist = append(list, value)
	return newlist
}
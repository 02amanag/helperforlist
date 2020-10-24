package learn

import "errors"

func DuplicacyCheck(list []float64) error {
	for _, i := range list {
		counter := 0
		for _, j := range list {
			if i == j {
				counter++
			}
			if counter > 1 {
				return errors.New("Data Duplicacy")
			}
		}
	}
	return nil
}

func DuplicacyRemove(list []float64) []float64 {
	var newlist []float64
	for _, i := range list {
		counter := 0
		for _, j := range newlist {
			if i == j {
				counter++
			}
		}
		if counter == 0 {
			newlist = append(newlist, i)
		}
	}
	return newlist
}

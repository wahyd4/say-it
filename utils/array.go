package utils

func Contains(item int, array []int) bool {
	for _, one := range array {
		if one == item {
			return true
		}
	}
	return false
}

package github.com/Gabriel0110/go-utilities

/*
	Go utility functions for common tasks that aren't readily available like Python's
*/

// Check if a slice contains an int value
func ContainsInt(slice []int, num int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}

// Check if a slice contains a float64 value
func ContainsFloat64(slice []float64, num float64) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}

// Check if a slice contains a string value
func ContainsString(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

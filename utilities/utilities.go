package utilities

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

func binarySearch(arr []int, l int, r int, x int) int {
	if r >= l {
		mid := l + (r-1)/2

		if arr[mid] == x {
			return mid
		} else if arr[mid] > x {
			return binarySearch(arr, l, mid-1, x)
		} else {
			return binarySearch(arr, mid+1, r, x)
		}
	} else {
		return -1
	}
}

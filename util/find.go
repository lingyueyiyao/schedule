package util

import "schedule/service"

// 1 3 5 6 7 9
// 0 1 2 3 4 5
func BinaryFind(c []*service.Customer, n int) int {
	i, j := 0, len(c) - 1
	for i < j {
		mid := i + (j-i)/2
		if c[mid].GetSerialNum() == n {
			return mid
		} else if c[mid].GetSerialNum() < n {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	if c[i].GetSerialNum() == n {
		return i
	}
	return -1
}

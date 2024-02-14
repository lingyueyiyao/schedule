package main

import "fmt"

var (
	numberOfCommon = 0
	numberOfBarrel = 0
)

const (
	threshNumber = 100
)

func GetNumber() int {
	ret := numberOfCommon%threshNumber + 1
	numberOfCommon = (numberOfCommon + 1) % threshNumber

	return ret
}

func main() {
	t := []int{1}
	i := 0
	t = append(t[0:i], t[i+1:]...)
	fmt.Println(t)
}

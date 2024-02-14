package util

import (
	"fmt"
	"schedule/service"
	"testing"
)

func TestBinaryFind(t *testing.T) {
	c := make([]*service.Customer, 0)
	c = append(c, service.NewCustomer(1, 1, service.TypeShower))
	c = append(c, service.NewCustomer(1, 3, service.TypeShower))
	c = append(c, service.NewCustomer(1, 5, service.TypeShower))
	c = append(c, service.NewCustomer(1, 6, service.TypeShower))
	c = append(c, service.NewCustomer(1, 7, service.TypeShower))
	c = append(c, service.NewCustomer(1, 9, service.TypeShower))
	fmt.Println("ret: ", BinaryFind(c, 6))
}

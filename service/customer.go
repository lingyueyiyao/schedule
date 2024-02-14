package service

import (
	"fmt"
)

type BathType int

const (
	TypeShower = iota
	TypeBarrel
)

type Customer struct {
	number int
	serialNum int
	bathType BathType
	message string
}

func NewCustomer(number int, serialNum int, bathType BathType) *Customer {
	t := "淋浴"
	if bathType == TypeBarrel {
		t = "木桶"
	}
	return &Customer{
		number: number,
		serialNum: serialNum,
		bathType: bathType,
		message: fmt.Sprintf("%s%d号%d人", t, serialNum, number),
	}
}

func (c *Customer) GetNumber() int {
	return c.number
}

func (c *Customer) GetSerialNum() int {
	return c.serialNum
}

func (c *Customer) GetBathType() BathType {
	return c.bathType
}

func (c *Customer) GetMessage() string {
	return c.message
}

func (c *Customer) SetNumber(n int) {
	c.number = n
}

func (c *Customer) SetSerial(n int) {
	c.serialNum = n
}

func (c *Customer) SetBathType(t BathType) {
	c.bathType = t
}

func (c *Customer) SetMessage(m string) {
	c.message = m
}

func (c *Customer) Cover(customer *Customer) {
	c.number = customer.number
	c.message = customer.message
}

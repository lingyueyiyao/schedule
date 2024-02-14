package service

import (
	"errors"
	"fmt"
	"sync"
)

type WaitQueue struct {
	showerQueue        []*Customer
	barrelQueue        []*Customer
	showerMu, barrelMu sync.Mutex
}

func NewWaitQueue() *WaitQueue {
	return &WaitQueue{
		showerQueue: []*Customer{},
		barrelQueue: []*Customer{},
	}
}

func (w *WaitQueue) SetShower(number int, serialNum int) {
	customer := NewCustomer(number, serialNum, TypeShower)
	w.SetShowerCustomer(customer)
}

func (w *WaitQueue) SetShowerCustomer(customer *Customer) {
	w.showerMu.Lock()
	w.showerQueue = append(w.showerQueue, customer)
	w.showerMu.Unlock()
}

func (w *WaitQueue) SetBarrel(number int, serialNum int) {
	customer := NewCustomer(number, serialNum, TypeBarrel)
	w.SetBarrelCustomer(customer)
}

func (w *WaitQueue) SetBarrelCustomer(customer *Customer) {
	w.barrelMu.Lock()
	w.barrelQueue = append(w.barrelQueue, customer)
	w.barrelMu.Unlock()
}

func (w *WaitQueue) GetShower() []*Customer {
	return w.showerQueue
}

func (w *WaitQueue) GetBarrel() []*Customer {
	return w.barrelQueue
}

func (w *WaitQueue) DelShower(n int) {
	w.showerMu.Lock()
	w.showerQueue = append(w.showerQueue[0:n], w.showerQueue[n+1:]...)
	w.showerMu.Unlock()
}

func (w *WaitQueue) DelBarrel(n int) {
	w.barrelMu.Lock()
	w.barrelQueue = append(w.barrelQueue[0:n], w.barrelQueue[n+1:]...)
	w.barrelMu.Unlock()
}

func (w *WaitQueue) DelCustomerBySerial(t BathType, n int) (*Customer, error) {
	q := w.showerQueue
	if t == TypeBarrel {
		q = w.barrelQueue
	}

	idx := BinaryFind(q, n)
	if idx == -1 {
		return nil, errors.New(fmt.Sprintf("序号%d不存在", n))
	}
	customer := q[idx]

	if t == TypeShower {
		w.DelShower(idx)
	} else {
		w.DelBarrel(idx)
	}
	return customer, nil
}

func (w *WaitQueue) GetShowerCustomerBySerial(serialNum int) *Customer {
	for _, customer := range w.showerQueue {
		if customer.serialNum == serialNum {
			return customer
		}
	}
	return nil
}

func BinaryFind(c []*Customer, n int) int {
	i, j := 0, len(c)-1
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

// BinaryFindInsert 查找序号n的插入位置
func BinaryFindInsert(c []*Customer, n int) int {
	i, j := 0, len(c)-1
	for i <= j {
		mid := i + (j-i)/2
		if c[mid].GetSerialNum() < n {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}

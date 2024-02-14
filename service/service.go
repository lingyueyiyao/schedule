package service

import (
	"errors"
	"fmt"
	"math"
	"schedule/room"
)

var (
	serialNumberOfCommon = 0
	serialNumberOfBarrel = 0
)

const (
	threshSerialNumber = 100
)

type Core struct {
	waitQueue  *WaitQueue
	roomRecord *room.Record
}

func NewCore() *Core {
	return &Core{
		waitQueue:  NewWaitQueue(),
		roomRecord: room.NewRecord(),
	}
}

func (c *Core) GetSerialNumber(t BathType, n int) int {
	ret := serialNumberOfCommon % threshSerialNumber + 1
	if t == TypeBarrel {
		ret = serialNumberOfBarrel % threshSerialNumber + 1
		serialNumberOfBarrel = (serialNumberOfBarrel + 1) % threshSerialNumber
		c.waitQueue.SetBarrel(n, ret)
	} else {
		serialNumberOfCommon = (serialNumberOfCommon + 1) % threshSerialNumber
		c.waitQueue.SetShower(n, ret)
	}

	return ret
}

func (c *Core) DelSerialNumber(t BathType, n int) error {
	_, err := c.waitQueue.DelCustomerBySerial(t, n)
	if err != nil {
		return errors.New(fmt.Sprintf("删除的号码%d不存在", n))
	}

	return nil
}

// ShareRoom 注意不要在min为第一个序号时拼房间，防止min已经被叫号
func (c *Core) ShareRoom(serialNums ...int) error {
	if len(serialNums) <= 1 {
		return errors.New("请输入足够数量的交换者")
	}

	min := minSlice(serialNums)
	first := c.waitQueue.GetShowerCustomerBySerial(min)
	message := first.message
	for _, num := range serialNums {
		if num == min {
			continue
		}
		customer, err := c.waitQueue.DelCustomerBySerial(TypeShower, num)
		if err != nil {
			return err
		}
		first.SetNumber(first.number + 1)
		message += customer.message
	}
	first.SetMessage(message + "拼房间")
	return nil
}

func (c *Core) CallSerialNumber(roomNum int) error {
	r := room.Rooms[roomNum]
	switch r.GetRoomType() {
	case room.TypeRoomTwo:
		customer := c.findCustomerForTwoType()
		if customer != nil {
			c.roomRecord.DispatchIdleRoom(roomNum)
			fmt.Println(fmt.Sprintf("请%s到%d号房间", customer.message, roomNum))
		} else {
			c.roomRecord.SetIdle(roomNum)
			if len(c.waitQueue.showerQueue) != 0 {
				fmt.Println("请分配更大的房间")
			} else {
				fmt.Println("没有淋浴顾客")
			}
		}
	case room.TypeRoomThree:
		customer := c.findCustomerForThreeType()
		if customer != nil {
			c.roomRecord.DispatchIdleRoom(roomNum)
			fmt.Println(fmt.Sprintf("请%s到%d号房间", customer.message, roomNum))
		} else {
			c.roomRecord.SetIdle(roomNum)
			fmt.Println("没有淋浴顾客")
		}
	case room.TypeRoomBarrel:
		customer := c.findCustomerForBarrel()
		if customer != nil {
			c.roomRecord.DispatchIdleRoom(roomNum)
			fmt.Println(fmt.Sprintf("请%s到%d号房间", customer.message, roomNum))
		} else {
			c.roomRecord.SetIdle(roomNum)
			if len(c.waitQueue.showerQueue) != 0 {
				fmt.Println("请分配更大的房间")
			} else {
				fmt.Println("没有木桶顾客")
			}
		}
	}
	return nil
}

// 查找第一个人数大于等于n的顾客
func (c *Core) findCustomerByBigThanNum(n int) (int, *Customer) {
	for i, customer := range c.waitQueue.showerQueue {
		if customer.number >= n {
			c.waitQueue.DelShower(i)
			return i, customer
		}
	}

	return -1, nil
}

func (c *Core) findCustomerByLessThanNum(n int) (int, *Customer) {
	for i, customer := range c.waitQueue.showerQueue {
		if customer.number <= n {
			c.waitQueue.DelShower(i)
			return i, customer
		}
	}

	return -1, nil
}

func (c *Core) findCustomerForTwoType() *Customer {
	idx, customer := c.findCustomerBetween(1, 2)
	if idx == -1 {
		return nil
	}
	return customer
}

func (c *Core) findCustomerBetween(lo int, hi int) (int, *Customer) {
	for i, customer := range c.waitQueue.showerQueue {
		if customer.number <= hi && customer.number >= lo {
			c.waitQueue.DelShower(i)
			return i, customer
		}
	}

	return -1, nil
}

func (c *Core) findCustomerForThreeType() *Customer {
	idx, customer := c.findCustomerByBigThanNum(2)
	if idx == -1 {
		idx, customer = c.findCustomerByLessThanNum(1)
		if idx == -1 {
			return nil
		}
		return customer
	}
	return customer
}

func (c *Core) findCustomerForBarrel() *Customer {
	if len(c.waitQueue.barrelQueue) != 0 {
		ret := c.waitQueue.GetBarrel()[0]
		c.waitQueue.DelBarrel(0)
		return ret
	}
	idx, customer := c.findCustomerByLessThanNum(2)
	if idx != -1 {
		return customer
	}
	return nil
}

func (c *Core) ChangeSerialNumAward(t BathType, from int, to int) error {
	customer, err := c.waitQueue.DelCustomerBySerial(t, from)
	if err != nil {
		return err
	}
	q := c.waitQueue.showerQueue
	if t == TypeBarrel {
		q = c.waitQueue.barrelQueue
	}
	idx := BinaryFindInsert(q, to)
	if q[idx].GetSerialNum() == to {
		q[idx].Cover(customer)
		return errors.New(fmt.Sprintf("%d号已存在，已修改顾客信息", to))
	}

	customer.serialNum = to
	newSlice := make([]*Customer, 0)
	copy(newSlice, q[:idx])
	newSlice = append(append(newSlice, customer), q[idx:]...)

	switch t {
	case TypeShower:
		c.waitQueue.showerQueue = newSlice
	case TypeBarrel:
		c.waitQueue.barrelQueue = newSlice
	}
	return nil
}

func (c *Core) ExchangeBathType(from BathType, to BathType, n int) (int, error) {
	customer, err := c.waitQueue.DelCustomerBySerial(from, n)
	if err != nil {
		return -1, err
	}

	number := c.GetSerialNumber(to, customer.number)

	return number, nil
}

func minSlice(nums []int) int {
	min := math.MaxInt
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

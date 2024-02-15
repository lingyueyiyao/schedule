package room

import (
	"context"
	"errors"
	"time"
)

var (
	ErrRoomOccupied = errors.New("room has been occupied")
)

type TypeRoom uint32

const (
	TypeRoomTwo = iota
	TypeRoomThree
	TypeRoomBarrel
)

var Rooms map[int]*room

type room struct {
	roomType   TypeRoom
	number     int
	capacity   int
	useTime    int
	timeCtx    context.Context
	timeCancel context.CancelFunc
	empty      bool
}

func NewRoom(n int, c int, t TypeRoom) *room {
	return &room{
		roomType: t,
		number:   n,
		capacity: c,
		empty:    true,
	}
}

func (r *room) IsEmpty() bool {
	if r.empty {
		return true
	}
	return false
}

func (r *room) OccupyRoom() error {
	if !r.IsEmpty() {
		return ErrRoomOccupied
	}
	r.empty = false
	return nil
}

func (r *room) GetRoomType() TypeRoom {
	return r.roomType
}

func (r *room) GetNumber() int {
	return r.number
}

func (r *room) GetCapacity() int {
	return r.capacity
}

func (r *room) GetUseTime() int {
	return r.useTime
}

func (r *room) AddUseTime(n int) {
	r.useTime += n
}

func (r *room) StartTimer() {
	ctx, cancel := context.WithCancel(context.Background())
	r.timeCtx = ctx
	r.timeCancel = cancel

	go func() {
		for {
			select {
			case <-time.After(60 * time.Second):
				r.useTime++
			case <-r.timeCtx.Done():
				return
			}
		}
	}()
}

func (r *room) CancelTimer() {
	r.timeCancel()
}

func init() {
	Rooms = map[int]*room{
		201: NewRoom(201, 2, TypeRoomTwo),
		202: NewRoom(202, 2, TypeRoomTwo),
		203: NewRoom(203, 2, TypeRoomTwo),
		204: NewRoom(204, 2, TypeRoomTwo),
		205: NewRoom(205, 2, TypeRoomBarrel),
		206: NewRoom(206, 2, TypeRoomBarrel),
		207: NewRoom(207, 2, TypeRoomBarrel),
		208: NewRoom(208, 2, TypeRoomBarrel),
		209: NewRoom(209, 2, TypeRoomTwo),
		210: NewRoom(210, 2, TypeRoomTwo),
		211: NewRoom(211, 2, TypeRoomTwo),
		212: NewRoom(212, 2, TypeRoomTwo),
		213: NewRoom(213, 2, TypeRoomTwo),
		214: NewRoom(214, 2, TypeRoomTwo),
		215: NewRoom(215, 2, TypeRoomTwo),
		216: NewRoom(216, 2, TypeRoomTwo),
		217: NewRoom(217, 3, TypeRoomThree),
		218: NewRoom(218, 3, TypeRoomThree),
		219: NewRoom(219, 3, TypeRoomThree),
		220: NewRoom(220, 3, TypeRoomThree),
		221: NewRoom(221, 3, TypeRoomThree),
		222: NewRoom(222, 3, TypeRoomThree),
	}
}

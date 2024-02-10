package schedule

type roomType uint32

const (
	roomTypeTwo = iota
	roomTypeThree
	roomTypeBarrel
)

var rooms []*room

type room struct {
	number   uint32
	capacity uint32
	roomType roomType
	empty    bool
}

func newRoom(n uint32, c uint32, t roomType) *room {
	return &room{
		number:   n,
		capacity: c,
		roomType: t,
		empty:    true,
	}
}

func init() {
	rooms = []*room{
		newRoom(201, 2, roomTypeTwo),
		newRoom(202, 2, roomTypeTwo),
		newRoom(203, 2, roomTypeTwo),
		newRoom(204, 2, roomTypeTwo),
		newRoom(205, 2, roomTypeBarrel),
		newRoom(206, 2, roomTypeBarrel),
		newRoom(207, 2, roomTypeBarrel),
		newRoom(208, 2, roomTypeBarrel),
		newRoom(209, 2, roomTypeTwo),
		newRoom(210, 2, roomTypeTwo),
		newRoom(211, 2, roomTypeTwo),
		newRoom(212, 2, roomTypeTwo),
		newRoom(213, 2, roomTypeTwo),
		newRoom(214, 2, roomTypeTwo),
		newRoom(215, 2, roomTypeTwo),
		newRoom(216, 2, roomTypeTwo),
		newRoom(217, 3, roomTypeThree),
		newRoom(218, 3, roomTypeThree),
		newRoom(219, 3, roomTypeThree),
		newRoom(220, 3, roomTypeThree),
		newRoom(221, 3, roomTypeThree),
		newRoom(222, 3, roomTypeThree),
	}
}

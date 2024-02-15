package room

import "sync"

type Record struct {
	occupiedRoom map[int]*room
	idleRoom     map[int]*room
	occupiedMu   sync.Mutex
	idleMu       sync.Mutex
}

func NewRecord() *Record {
	r := &Record{
		occupiedRoom: make(map[int]*room),
		idleRoom:     make(map[int]*room),
	}
	for i := range Rooms {
		r.idleRoom[Rooms[i].number] = Rooms[i]
	}
	return r
}

func (r *Record) SetOccupy(n int) {
	r.occupiedMu.Lock()
	defer r.occupiedMu.Unlock()
	r.occupiedRoom[n] = Rooms[n]
	Rooms[n].empty = false
}

func (r *Record) SetIdle(n int) {
	r.idleMu.Lock()
	defer r.idleMu.Unlock()
	Rooms[n].CancelTimer()
	r.idleRoom[n] = Rooms[n]
	Rooms[n].empty = true
}

func (r *Record) DelOccupy(n int) {
	r.occupiedMu.Lock()
	defer r.occupiedMu.Unlock()
	delete(r.occupiedRoom, n)
}

func (r *Record) DelIdle(n int) {
	r.idleMu.Lock()
	defer r.idleMu.Unlock()
	delete(r.idleRoom, n)
}

func (r *Record) GetOccupy() map[int]*room {
	return r.occupiedRoom
}

func (r *Record) GetIdle() map[int]*room {
	return r.idleRoom
}

func (r *Record) DispatchIdleRoom(n int) {
	r.DelIdle(n)
	r.SetOccupy(n)
	Rooms[n].StartTimer()
}

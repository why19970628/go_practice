package room

import "testing"

func TestMachine_RoomState(t *testing.T) {
	room := RoomContext{Num: 101}
	free := new(FreeState)
	booked := new(BookedState)
	checkedIn := new(CheckedInState)
	room.SetState(free)
	room.SetState(booked)
	room.SetState(checkedIn)
}

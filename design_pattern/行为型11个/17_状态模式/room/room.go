package room

import "fmt"

type RoomState interface {
	Handle(int)
}
type FreeState struct {
}

func (this *FreeState) Handle(num int) {
	fmt.Printf("%d 房间空闲，可以入住。", num)
}

type BookedState struct {
}

func (this *BookedState) Handle(num int) {
	fmt.Printf("%d 房间已经有人预定。", num)
}

type CheckedInState struct {
}

func (this *CheckedInState) Handle(num int) {
	fmt.Printf("%d 房间已经有人入住。", num)
}

type RoomContext struct {
	Num   int // 房间号码
	State RoomState
}

func (this *RoomContext) SetState(state RoomState) {
	this.State = state
	this.State.Handle(this.Num)
}
func main() {
}

package revert_lib

import "fmt"

type ICourse interface {
	ID() int
	Name() string
	SetUser(IUser)
	Study()
}

type GolangCourse struct {
	iID          int
	sName        string
	xCurrentUser IUser
}

func NewGolangCourse() ICourse {
	return &GolangCourse{
		iID:          11,
		sName:        "golang",
		xCurrentUser: nil,
	}
}

func (me *GolangCourse) ID() int {
	return me.iID
}

func (me *GolangCourse) Name() string {
	return me.sName
}

func (me *GolangCourse) SetUser(user IUser) {
	me.xCurrentUser = user
}

func (me *GolangCourse) Study() {
	fmt.Printf("%v is learning %v\n", me.xCurrentUser.Name(), me.Name())
}

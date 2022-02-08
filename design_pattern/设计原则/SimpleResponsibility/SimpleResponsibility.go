package SimpleResponsibility

import "fmt"

type Course interface {
	Name() string
	ID() int
	Controller() BaseController
}

type BaseController interface {
	Play()
	Stop()
}

type ReplayController interface {
	BaseController
	Forward(seconds int)
	Backward(seconds int)
}

type CourseInfo struct {
	iID   int
	sName string
}

func (me *CourseInfo) ID() int {
	return me.iID
}

func (me *CourseInfo) Name() string {
	return me.sName
}

// 直播课
type LiveCourse struct {
	CourseInfo
}

func NewLiveCourse(courseInfo CourseInfo) Course {
	return &LiveCourse{CourseInfo: courseInfo}
}

func (me *LiveCourse) Controller() BaseController {
	return me
}

func (me *LiveCourse) Play() {
	fmt.Printf("%v play\n", me.Name())
}

func (me *LiveCourse) Stop() {
	fmt.Printf("%v pause\n", me.Name())
}

// 重播课
type ReplayCourse struct {
	CourseInfo
}

func NewReplayCourse(courseInfo CourseInfo) Course {
	return &ReplayCourse{CourseInfo: courseInfo}
}

func (me *ReplayCourse) Controller() BaseController {
	return me
}

func (me *ReplayCourse) Play() {
	fmt.Printf("%v play\n", me.Name())
}

func (me *ReplayCourse) Stop() {
	fmt.Printf("%v pause\n", me.Name())
}

func (me *ReplayCourse) Forward(seconds int) {
	fmt.Printf("%v forward %v\n", me.Name(), seconds)
}

func (me *ReplayCourse) Backward(seconds int) {
	fmt.Printf("%v backward %v\n", me.Name(), seconds)
}

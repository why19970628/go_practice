package 单一原则

// 单一职责原则（Simple Responsibility Principle, SRP）指不要存在一个以上导致类变更的原因。
// 假设有一个Class负责两个职责，一旦发生需求变更，修改其中一个职责的逻辑代码，有可能会导致另一个职责的功能发生故障。

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

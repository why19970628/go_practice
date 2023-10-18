package single

import "testing"

func Test_SimpleResponsibility(t *testing.T) {
	//fnTestBadCourse := func(bc ) {
	//	bc.Play()
	//	bc.Pause()
	//	bc.Forward(30)
	//	bc.Backward(30)
	//}
	//fnTestBadCourse(NewLiveCourse(CourseInfo{1, "直播课"}))
	//fnTestBadCourse(NewReplayCourse(CourseInfo{2, "录播课"}))

	fnTestGoodCourse := func(gc Course) {
		pc := gc.Controller()
		pc.Play()
		pc.Stop()
		if rc, ok := pc.(ReplayController); ok {
			rc.Forward(30)
			rc.Backward(30)
		}
	}

	fnTestGoodCourse(NewLiveCourse(CourseInfo{11, "直播课"}))
	fnTestGoodCourse(NewReplayCourse(CourseInfo{12, "录播课"}))
}

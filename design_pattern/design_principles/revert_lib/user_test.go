package revert_lib

import "testing"

func TestDIP_test(t *testing.T) {
	//bu := dip.NewBadUser(1, "Tom")
	//bu.StudyGolangCourse()

	gu := NewGoodUser(2, "Mike")
	gu.Study(NewGolangCourse())
}

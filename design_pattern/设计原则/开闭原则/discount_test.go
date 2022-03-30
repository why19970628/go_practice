package 开闭原则

import (
	"testing"
)

//开闭原则（Open-Closed Principle, OCP）
// 指一个软件实体如类、模块和函数应该对扩展开放，对修改关闭。
// 所谓开闭，也正是对扩展和修改两个行为的一个原则。

//开闭原则, 就是尽量避免修改, 改以扩展的方式, 实现系统功能的增加
//增加"优惠折扣"接口 - IDiscount
//增加"折后golang课程" - DiscountedGolangCourse, 同时实现课程接口和折扣接口
//DiscountedGolangCourse继承自GolangCourse, 添加实现折扣接口, 并覆盖ICourse.price()方法

func Test_open_close(t *testing.T) {
	fnShowCourse := func(it ICourse) {
		t.Logf("id=%v, name=%v, price=%v\n", it.ID(), it.Name(), it.Price())
	}

	c1 := NewGolangCourse(1, "golang课程", 100)
	fnShowCourse(c1)

	c2 := NewDiscountedGolangCourse(2, "golang优惠课程", 100, 0.6)
	fnShowCourse(c2)
}

package 开闭原则

type IDiscount interface {
	Discount() float64
}

type DiscountedGolangCourse struct {
	GolangCourse
	fDiscount float64
}

func NewDiscountedGolangCourse(id int, name string, price float64, discount float64) ICourse {
	return &DiscountedGolangCourse{
		GolangCourse: GolangCourse{
			iID:    id,
			sName:  name,
			fPrice: price,
		},
		fDiscount: discount,
	}
}

// implements IDiscount.Discount
func (me *DiscountedGolangCourse) Discount() float64 {
	return me.fDiscount
}

// overwrite ICourse.Price
func (me *DiscountedGolangCourse) Price() float64 {
	return me.fDiscount * me.GolangCourse.Price()
}

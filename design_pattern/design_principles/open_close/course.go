package open_close

type ICourse interface {
	ID() int
	Name() string
	Price() float64
}

type GolangCourse struct {
	iID    int
	sName  string
	fPrice float64
}

func NewGolangCourse(id int, name string, price float64) ICourse {
	return &GolangCourse{
		iID:    id,
		sName:  name,
		fPrice: price,
	}
}

func (me *GolangCourse) ID() int {
	return me.iID
}

func (me *GolangCourse) Name() string {
	return me.sName
}

func (me *GolangCourse) Price() float64 {
	return me.fPrice
}

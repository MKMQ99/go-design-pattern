package prototype

type Shape struct {
	x     int
	y     int
	color string
}

func (sh Shape) Constructor(shape Shape) {
	sh.x = shape.x
	sh.y = shape.y
	sh.color = shape.color
}

func (sh Shape) Clone() Shape {
	return Shape{
		x:     sh.x,
		y:     sh.y,
		color: sh.color,
	}
}

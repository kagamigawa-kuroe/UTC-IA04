package mypackage

import "math"

type Point2D struct{
	x float32
	y float32
}

func (p *Point2D) NewPoint2D(x, y float64){
	p.x = float32(x)
	p.y = float32(y)
}

func (p *Point2D) GetterX() float32{
	return p.x
}

func (p *Point2D) GetterY() float32{
	return p.y
}

func (p *Point2D) SetterX(my_x float32){
	p.x = my_x
}

func (p *Point2D) SetterY(my_y float32){
	p.y = my_y
}

func (p *Point2D) clone(src Point2D){
	p.x = src.x
	p.y = src.y
}

func (p *Point2D) Module(src Point2D) float64{
	a := math.Pow(math.Abs(float64(p.x-src.x)),2) + math.Pow(math.Abs(float64(p.y-src.y)),2)
	return math.Sqrt(a)
}
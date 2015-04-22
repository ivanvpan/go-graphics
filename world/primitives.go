package world

import "math"
import "fmt"

type Vector struct {
	X, Y, Z float64
}

func (p Vector) Rotate(angle float64, axis byte) Vector {
	toRadians := func(angle float64) float64 {
		return math.Pi * (angle / 180.0)
	}
	radAngle := toRadians(angle)
	newX := math.Ceil(p.X*math.Cos(radAngle) - p.Y*math.Sin(radAngle))
	newY := math.Ceil(p.Y*math.Cos(radAngle) + p.X*math.Sin(radAngle))
	p.X = newX
	p.Y = newY
	fmt.Printf("%g %d %d\n", radAngle, p.X, p.Y)
	return p
}

func (p Vector) Translate(x, y float64) Vector {
	p.X = p.X + x
	p.Y = p.Y + y
	return p
}

func (p Vector) RotateAround(angle float64, x float64, y float64, axis byte) Vector {
	translated := p.Translate(-x, -y)
	rotated := translated.Rotate(angle, axis)
	return rotated.Translate(x, y)
}

type Line struct {
	Start, End Vector
}

type Polygon struct {
	Vertices []Vector
}

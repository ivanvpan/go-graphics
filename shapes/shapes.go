package shapes

import "math"
import "fmt"

import "github.com/veandco/go-sdl2/sdl"

func round(num float64) int32 {
	return int32(num + math.Copysign(0.5, num))
}

type Shape interface {
	draw(surface *sdl.Surface, color uint32)
}

type Point struct {
	X, Y float64
}

func (p Point) Rotate(angle int32) Point {
	toRadians := func(angle int32) float64 {
		return math.Pi * (float64(angle) / 180.0)
	}
	radAngle := toRadians(angle)
	newX := int32(math.Ceil(float64(p.X)*math.Cos(radAngle) - float64(p.Y)*math.Sin(radAngle)))
	newY := int32(math.Ceil(float64(p.Y)*math.Cos(radAngle) + float64(p.X)*math.Sin(radAngle)))
	p.X = newX
	p.Y = newY
	fmt.Printf("%g %d %d\n", radAngle, p.X, p.Y)
	return p
}

func (p Point) Translate(x, y int32) Point {
	p.X = p.X + x
	p.Y = p.Y + y
	return p
}

func (p Point) RotateAround(angle, x, y int32) Point {
	translated := p.Translate(-x, -y)
	rotated := translated.Rotate(angle)
	return rotated.Translate(x, y)
}

func (p Point) Draw(surface *sdl.Surface, color uint32) {
	rect := sdl.Rect{p.X, p.Y, 1, 1}
	surface.FillRect(&rect, color)
}

type Line struct {
	Start Point
	End   Point
}

func (l Line) Draw(surface *sdl.Surface, color uint32) {
	deltax := float64(l.End.X - l.Start.X)
	deltay := float64(l.End.Y - l.Start.Y)
	slope := math.Abs(deltay / deltax)
	var err float64 = 0
	var y int32 = l.Start.Y
	draw := func(x int32) {
		pixel := sdl.Rect{x, y, 1, 1}
		surface.FillRect(&pixel, color)
		err += slope
		for err >= 0.5 {
			err -= 1
			y += 1 * int32(math.Copysign(1, deltay))
			pixel := sdl.Rect{x, y, 1, 1}
			surface.FillRect(&pixel, color)
		}
	}
	if deltax > 0 {
		for x := l.Start.X; x <= l.End.X; x++ {
			draw(x)
		}
	} else if deltax < 0 {
		for x := l.Start.X; x >= l.End.X; x-- {
			draw(x)
		}
	} else {
		if deltay > 0 {
			for y := l.Start.Y; y <= l.End.Y; y++ {
				pixel := sdl.Rect{l.Start.X, y, 1, 1}
				surface.FillRect(&pixel, color)
			}
		} else {
			for y := l.Start.Y; y >= l.End.Y; y-- {
				pixel := sdl.Rect{l.Start.X, y, 1, 1}
				surface.FillRect(&pixel, color)
			}
		}
	}
}

type Polygon struct {
	Vertices []Point
}

func (p Polygon) Draw(surface *sdl.Surface, color uint32) {
	i1, i2 := 0, 1
	for i1 < len(p.Vertices) {
		var line Line
		if i1 == len(p.Vertices)-1 {
			line = Line{p.Vertices[i1], p.Vertices[0]}
		} else {
			line = Line{p.Vertices[i1], p.Vertices[i2]}
		}
		i1 += 1
		i2 += 1
		line.Draw(surface, color)
	}

}

type Circle struct {
	xCenter int32
	yCenter int32
	radius  int32
}

func (c Circle) Draw(surface *sdl.Surface, color uint32) {
}

package main

import "github.com/veandco/go-sdl2/sdl"

import "ivan/go-graphics/shapes"

func clearSurface(srcSurface, destSurface *sdl.Surface) {
	rect := sdl.Rect{0, 0, destSurface.W, destSurface.H}
	srcSurface.FillRect(&rect, 0x09000000)
	//srcSurface.FillRect(&rect, 0xff000000)
	srcSurface.Blit(&rect, destSurface, &rect)
}

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)

	const winHeight = 600
	const winWidth = 800

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	blittingSurface, err := sdl.CreateRGBSurface(0, surface.W, surface.H, 32, 0x00ff0000, 0x0000ff00, 0x000000ff, 0xff000000)
	if err != nil {
		panic(err)
	}

	quit := false
	polygon := shapes.Polygon{
		[]shapes.Point{
			shapes.Point{0, 0},
			shapes.Point{300, 300},
			shapes.Point{0, 300},
		},
	}
	line := shapes.Line{
		shapes.Point{0, 0},
		shapes.Point{(winWidth - 1) / 2, winHeight - 1},
	}
	point := shapes.Point{200, 200}
	for !quit {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				quit = true
			}
		}

		clearSurface(blittingSurface, surface)

		line.Draw(surface, 0xffff0000)
		polygon.Draw(surface, 0xffff0000)
		point.Draw(surface, 0xff00ffff)

		point = point.RotateAround(5, winWidth/2, winHeight/2)

		window.UpdateSurface()
		sdl.Delay(1000 / 24)
	}
	sdl.Quit()
}

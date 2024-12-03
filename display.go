package main

import rl "github.com/gen2brain/raylib-go/raylib"

func drawSun(sun Center, w int32, h int32) {
	w += 2 * int32(sun.Position[0])
	h += 2 * int32(sun.Position[1])
	rl.DrawCircle(w/2, h/2, 50, rl.Yellow)
}

func drawObject(obj SpaceObject) {
	rl.DrawCircle(int32(obj.Position[0])+screenWidth/2, -int32(obj.Position[1])+screenHeight/2, OBJ_RADIUS, obj.Color)
}

func drawSystem(system System) {
	drawSun(system.Sun, screenWidth, screenHeight)
	for i := 0; i < len(system.Objects); i++ {
		drawObject(system.Objects[i])
	}
}

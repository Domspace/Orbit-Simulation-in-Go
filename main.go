package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const screenWidth int32 = 1200
const screenHeight int32 = 1200

func getObjects() []SpaceObject {
	obj1 := SpaceObject{rl.Red, [2]float64{0.0, 250.0}, [2]float64{6.5, 0}, 7}
	obj2 := SpaceObject{rl.Blue, [2]float64{0.0, 150.0}, [2]float64{8, 0}, 4}
	obj3 := SpaceObject{rl.Green, [2]float64{0.0, 450.0}, [2]float64{3, 0}, 10}
	obj4 := SpaceObject{rl.Gray, [2]float64{0.0, 150.0}, [2]float64{6, 0}, 3}
	return []SpaceObject{obj1, obj2, obj3, obj4}
}

func main() {
	rl.InitWindow(screenWidth+200, screenHeight, "Orbit Simulation in Go")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	localsun := Center{[2]float64{0.0, 0.0}, 10000, 30}
	localsystem := System{localsun, getObjects()}
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		curkey := rl.GetKeyPressed()
		if curkey != 0 {
			fmt.Println(curkey)
		}
		rl.ClearBackground(rl.Black)
		drawSystem(localsystem)
		doStuff(&localsystem)
		drawUI()

		rl.EndDrawing()
	}
}

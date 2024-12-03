package main

import rl "github.com/gen2brain/raylib-go/raylib"

func drawUIBackGround() {
	rl.DrawRectangle(screenWidth, 0, 200, screenHeight, rl.Gray)
	rl.DrawRectangle(screenWidth+5, 5, 190, screenHeight-10, rl.DarkGray)
}

func drawUI() {
	drawUIBackGround()
}

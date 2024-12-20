package main

import rl "github.com/gen2brain/raylib-go/raylib"

type SpaceObject struct {
	Color    rl.Color
	Position [2]float64
	Velocity [2]float64
	Radius   float32
}
type Center struct {
	Position [2]float64
	Gravity  int
	Radius   float32
}
type System struct {
	Sun     Center
	Objects []SpaceObject
}

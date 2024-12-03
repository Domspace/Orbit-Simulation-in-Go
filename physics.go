package main

import (
	"math"
)

func get_dir(v1 [2]float64, v2 [2]float64) [2]float64 {
	dir := [2]float64{0.0, 0.0}
	dir[0] = v2[0] - v1[0]
	dir[1] = v2[1] - v1[1]
	return dir
}

func get_dist(v1 [2]float64, v2 [2]float64) float64 {
	var x float64
	x = math.Sqrt(math.Pow((v1[0]-v2[0]), 2) + math.Pow((v1[1]-v2[1]), 2))
	return x
}

func do_gravity(v1 [2]float64, v2 [2]float64, gravity int) [2]float64 {
	acc := float64(gravity) / (math.Pow(get_dist(v1, v2), 3))
	dir := get_dir(v1, v2)
	dir[0] *= acc
	dir[1] *= acc
	return dir
}
func do_velocity(pos [2]float64, currentVelocity [2]float64, newVelocity [2]float64) ([2]float64, [2]float64) {
	currentVelocity[0] += newVelocity[0]
	currentVelocity[1] += newVelocity[1]
	pos[0] += currentVelocity[0]
	pos[1] += currentVelocity[1]
	return pos, currentVelocity
}
func doStuff(system *System) {
	grav := system.Sun.Gravity
	gravPos := system.Sun.Position
	for i := 0; i < len(system.Objects); i++ {
		pos := system.Objects[i].Position
		//fmt.Println(pos)
		vel := system.Objects[i].Velocity
		system.Objects[i].Position, system.Objects[i].Velocity = do_velocity(pos, vel, do_gravity(pos, gravPos, grav))
	}
}

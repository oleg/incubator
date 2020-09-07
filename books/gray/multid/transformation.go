package multid

import "math"

//todo: where to put this methods?
func Translation(x, y, z float64) Matrix4 {
	return Matrix4{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	}
}
func Scaling(x, y, z float64) Matrix4 {
	return Matrix4{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	}
}
func RotationX(r float64) Matrix4 {
	return Matrix4{
		{1, 0, 0, 0},
		{0, math.Cos(r), -math.Sin(r), 0},
		{0, math.Sin(r), math.Cos(r), 0},
		{0, 0, 0, 1},
	}
}

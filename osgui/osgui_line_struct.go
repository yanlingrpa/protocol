package osgui

import "math"

type StraightLine struct {
	X1 int `json:"x1"`
	Y1 int `json:"y1"`
	X2 int `json:"x2"`
	Y2 int `json:"y2"`
}

func (l StraightLine) Clone() StraightLine {
	return StraightLine{
		X1: l.X1,
		Y1: l.Y1,
		X2: l.X2,
		Y2: l.Y2,
	}
}

func (l StraightLine) Length() float64 {
	dx := float64(l.X2 - l.X1)
	dy := float64(l.Y2 - l.Y1)
	return math.Hypot(dx, dy)
}

func (l StraightLine) AngleWithHorizontal() float64 {
	dx := float64(l.X2 - l.X1)
	dy := float64(l.Y2 - l.Y1)
	return math.Atan2(dy, dx) // 返回弧度，范围[-π, π]
}

func (l StraightLine) AngleWithHorizontalDeg() float64 {
	deg := l.AngleWithHorizontal() * 180 / math.Pi
	if deg < 0 {
		deg += 180
	}
	return deg
}

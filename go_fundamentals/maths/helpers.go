package clockface

import (
	"math"
	"time"
)

func SimpleTime(hour, minute, second int) time.Time {
	return time.Date(312, time.April, 28, hour, minute, second, 0, time.UTC)
}

func TestName(t time.Time) string {
	return t.Format("15:04:05")
}

func RoughlyEqualFloat64(num1, num2 float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(num1-num2) < equalityThreshold
}

func RoughlyEqualPoint(point1, point2 Point) bool {
	return RoughlyEqualFloat64(point1.X, point2.X) && RoughlyEqualFloat64(point1.Y, point2.Y)
}

func AngleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

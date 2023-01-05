package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCenterX     = 150
	clockCenterY     = 150
)

func SecondInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func SecondHandPoint(t time.Time) Point {
	return AngleToPoint(SecondInRadians(t))
}

func SecondHand(w io.Writer, t time.Time) {
	p := MakeHand(SecondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func MinutesInRadians(t time.Time) float64 {
	return (SecondInRadians(t) / 60) + math.Pi/(30/float64(t.Minute()))
}

func MinuteHandPoint(t time.Time) Point {
	return AngleToPoint(MinutesInRadians(t))
}

func MinuteHand(w io.Writer, t time.Time) {
	p := MakeHand(MinuteHandPoint(t), minuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:4.5px;"/>`, p.X, p.Y)
}

func HoursInRadians(t time.Time) float64 {
	return MinutesInRadians(t)/12 + math.Pi/(6/float64(t.Hour()%12))
}

func HourHandPoint(t time.Time) Point {
	return AngleToPoint(HoursInRadians(t))
}

func HourHand(w io.Writer, t time.Time) {
	p := MakeHand(HourHandPoint(t), hourHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:6px;"/>`, p.X, p.Y)
}

func MakeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}
	return p
}

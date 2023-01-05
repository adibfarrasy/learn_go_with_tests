package clockface_test

import (
	"math"
	"testing"
	"time"

	. "maths"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{SimpleTime(0, 0, 30), math.Pi},
		{SimpleTime(0, 0, 0), 0},
		{SimpleTime(0, 0, 45), math.Pi / 2 * 3},
		{SimpleTime(0, 0, 7), math.Pi / 30 * 7},
	}

	for _, test := range cases {
		t.Run(TestName(test.time), func(t *testing.T) {
			want := test.angle
			got := SecondInRadians(test.time)

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{SimpleTime(0, 0, 30), Point{0, -1}},
		{SimpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(TestName(c.time), func(t *testing.T) {
			got := SecondHandPoint(c.time)
			want := c.point

			if !RoughlyEqualPoint(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{SimpleTime(0, 30, 0), math.Pi},
		{SimpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, c := range cases {
		t.Run(TestName(c.time), func(t *testing.T) {
			got := MinutesInRadians(c.time)
			want := c.angle

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{SimpleTime(0, 30, 0), Point{0, -1}},
		{SimpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(TestName(c.time), func(t *testing.T) {
			got := MinuteHandPoint(c.time)
			want := c.point

			if !RoughlyEqualPoint(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{SimpleTime(6, 0, 0), math.Pi},
		{SimpleTime(0, 0, 0), 0},
		{SimpleTime(21, 0, 0), math.Pi * 1.5},
		{SimpleTime(0, 1, 30), math.Pi / (6 * 60 * 60 / 90)},
	}

	for _, c := range cases {
		t.Run(TestName(c.time), func(t *testing.T) {
			got := HoursInRadians(c.time)
			want := c.angle

			if !RoughlyEqualFloat64(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{SimpleTime(6, 0, 0), Point{0, -1}},
		{SimpleTime(21, 0, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(TestName(c.time), func(t *testing.T) {
			got := HourHandPoint(c.time)
			if !RoughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

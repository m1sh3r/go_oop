package main

import "testing"

func TestCircleArea(t *testing.T) {
	circle := Circle{Radius: 10}
	got := circle.Area()
	want := 314.0

	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

func TestSquareArea(t *testing.T) {
	square := Square{A: 10}
	got := square.Area()
	want := 100.0

	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

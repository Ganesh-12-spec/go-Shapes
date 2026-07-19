package main

import "testing"

func TestCircleArea(t *testing.T){
	c := Circle{radius:5}
	got := c.Area()
	want := 78.53981633974483

	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}

	r := Rectangle{width: 5, length: 5}
	got = r.Area()
	want = 25.0

	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}
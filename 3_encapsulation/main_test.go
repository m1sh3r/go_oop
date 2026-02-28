package main

import (
	person1 "incaps/person"
	"testing"
)

func TestPersonIntroduceFromMainPackage(t *testing.T) {
	p := person1.Person{}
	p.SetFirstName("Михаил")
	p.SetLastName("")

	got := p.Introduce()
	want := "Привет, я Михаил "

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

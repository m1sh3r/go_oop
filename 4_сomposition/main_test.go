package main

import (
	person1 "compose/person"
	"testing"
)

func TestPersonIntroduceFromMainPackage(t *testing.T) {
	p := person1.Person{}
	p.SetFirstName("Михаил")
	p.SetLastName("")
	p.SetAddress(person1.Address{City: "Тамбов"})

	got := p.Introduce()
	want := "Привет, я Михаил  из Тамбова"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

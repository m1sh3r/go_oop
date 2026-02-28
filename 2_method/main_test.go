package main

import (
	"testing"

	"github.com/akozadaev/go_oop/2_method/person"
)

func TestPersonIntroduceFromMainPackage(t *testing.T) {
	p := person.Person{FirstName: "Михаил", LastName: ""}
	got := p.Introduce()
	want := "Привет, я Михаил "

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

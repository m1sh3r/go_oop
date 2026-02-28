package main

import (
	"fmt"
	"testing"
)

func TestAnimalFields(t *testing.T) {
	animal := Animal{"Майя", "Кошка", 1}

	if animal.Name != "Майя" {
		t.Errorf("name = %q, want %q", animal.Name, "Майя")
	}

	if animal.Species != "Кошка" {
		t.Errorf("species = %q, want %q", animal.Species, "Кошка")
	}

	if animal.Age != 1 {
		t.Errorf("age = %d, want %d", animal.Age, 1)
	}
}

func TestAnimalFormat(t *testing.T) {
	animal := Animal{"Майя", "Кошка", 1}

	got := fmt.Sprintf("Кличка: %s,\n что за зверь: %s,\n возраст: %d год (лет)", animal.Name, animal.Species, animal.Age)
	want := "Кличка: Майя,\n что за зверь: Кошка,\n возраст: 1 год (лет)"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

package main

import "testing"

func TestAnimalsSpeak(t *testing.T) {
	tests := []struct {
		name   string
		animal Animal
		want   string
	}{
		{"dog", Dog{}, "Гав"},
		{"cat", Cat{}, "Мяу"},
		{"duck", Duck{}, "Кря-кря"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.animal.Speak()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

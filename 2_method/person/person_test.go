package person

import "testing"

func TestPersonIntroduce(t *testing.T) {
	p := Person{FirstName: "Михаил", LastName: ""}

	got := p.Introduce()
	want := "Привет, я Михаил "

	if got != want {
		t.Fatalf("Introduce() = %q, want %q", got, want)
	}
}

func TestPersonIntroduceWithEmptyNames(t *testing.T) {
	p := Person{}

	got := p.Introduce()
	want := "Привет, я  "

	if got != want {
		t.Fatalf("Introduce() with empty names = %q, want %q", got, want)
	}
}

package person

import "testing"

func TestSettersAndGetters(t *testing.T) {
	p := &Person{}
	p.SetFirstName("Михаил")
	p.SetLastName("")

	if p.GetFirstName() != "Михаил" {
		t.Fatalf("GetFirstName() = %q, want %q", p.GetFirstName(), "Михаил")
	}

	if p.GetLastName() != "" {
		t.Fatalf("GetLastName() = %q, want %q", p.GetLastName(), "")
	}
}

func TestIntroduce(t *testing.T) {
	p := &Person{}
	p.SetFirstName("Михаил")
	p.SetLastName("")

	got := p.Introduce()
	want := "Привет, я Михаил "

	if got != want {
		t.Fatalf("Introduce() = %q, want %q", got, want)
	}
}

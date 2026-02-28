package person

import "testing"

func TestSetAndGetAddress(t *testing.T) {
	p := &Person{}
	address := Address{City: "Тамбов"}
	p.SetAddress(address)

	got := p.GetAddress()
	if got.City != "Тамбов" {
		t.Fatalf("GetAddress().City = %q, want %q", got.City, "Тамбов")
	}
}

func TestIntroduceWithAddress(t *testing.T) {
	p := &Person{}
	p.SetFirstName("Михаил")
	p.SetLastName("")
	p.SetAddress(Address{City: "Тамбов"})

	got := p.Introduce()
	want := "Привет, я Михаил  из Тамбова"

	if got != want {
		t.Fatalf("Introduce() = %q, want %q", got, want)
	}
}

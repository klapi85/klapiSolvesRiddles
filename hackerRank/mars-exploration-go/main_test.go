package main

import "testing"

func TestCheckLetter(t *testing.T) {
	got := checkLetter("S", "S")
	want := 0

	if got != want {
		// t.Errorf zgłasza błąd i kontynuuje wykonywanie testu
		t.Errorf("checkLetter(\"S\", \"S\") = %d; expected %d", got, want)
	}
}

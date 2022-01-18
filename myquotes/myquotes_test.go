package myquotes

import "testing"

func TestDisplay(t *testing.T) {
	wanted := "I can eat glass and it doesn't hurt me."
	state := Display()
	if state != wanted {
		t.Errorf("Feil, fikk %v, ønsket %v.", state, wanted)
	}
}

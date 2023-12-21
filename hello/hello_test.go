package hello

import "testing"

func TestHello(t *testing.T) {
	box := NewSpeakerBox(5)
	box.Speak()
}

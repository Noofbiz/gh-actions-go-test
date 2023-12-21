package bleep

import "testing"

func TestBleepingCount(t *testing.T) {
	bleep := Bleep{Count: 1, Phrase: "hi"}
	bleep.Increase(4)
	if bleep.Count != 5 {
		t.Errorf("increase didn't work!")
	}
}

func BenchmarkBleepingCount(b *testing.B) {
	bleep := Bleep{Count: 5, Phrase: "hi"}

	for i := 0; i < b.N; i++ {
		bleep.SayStuff()
	}
}

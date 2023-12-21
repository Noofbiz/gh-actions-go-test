package bleep

import "fmt"

type Bleep struct {
	Count  int
	Phrase string
}

func (b *Bleep) Increase(n int) {
	b.Count += n
}

func (b *Bleep) SayStuff() {
	for i := 0; i < b.Count; i++ {
		fmt.Println(b.Phrase)
	}
}

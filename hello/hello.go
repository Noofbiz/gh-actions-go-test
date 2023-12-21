package hello

import bleep "github.com/Noofbiz/gh-actions-go-test"

type SpeakerBox struct {
	speaker bleep.Bleep
}

func NewSpeakerBox(Count int) *SpeakerBox {
	return &SpeakerBox{
		speaker: bleep.Bleep{
			Count:  Count,
			Phrase: "Hello!",
		},
	}
}

func (s *SpeakerBox) Speak() {
	s.speaker.SayStuff()
}

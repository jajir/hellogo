package lev3

import (
	"github.com/ev3go/ev3dev"
	"log"
	"time"
)

// SoundPath is the path to the ev3 sound events.
const soundPath = "/dev/input/by-path/platform-sound-event"

var Speaker = NewBeeper()

type Beeper struct {
	speaker *ev3dev.Speaker
}

func NewBeeper() Beeper {
	var out = *new(Beeper)
	out.speaker = ev3dev.NewSpeaker(soundPath)
	out.speaker.Init()
	return out
}

func (beeper *Beeper) Close() {
	beeper.speaker.Close()
}

func (beeper *Beeper) Beep() {
	// Play tone at 440Hz for 200ms...
	must(beeper.speaker.Tone(440), "Set tone to 440")
	time.Sleep(20 * time.Millisecond)

	// play tone at 220Hz for 200ms...
	must(beeper.speaker.Tone(220), "Set tone to 220")
	time.Sleep(20 * time.Millisecond)

	// then stop tone playback.
	must(beeper.speaker.Tone(0), "Set tone to 0")
}

func must(err error, s string) {
	if err != nil {
		log.Fatalf("%s %v", s, err)
	}
}

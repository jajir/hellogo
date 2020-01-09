package first

import (
	"log"
	"time"

	"github.com/ev3go/ev3dev"
)

// SoundPath is the path to the ev3 sound events.
const SoundPath = "/dev/input/by-path/platform-snd-legoev3-event"

var speaker = ev3dev.NewSpeaker(SoundPath)

func main() {
	must(speaker.Init())
	defer speaker.Close()

	// Play tone at 440Hz for 200ms...
	must(speaker.Tone(440))
	time.Sleep(200 * time.Millisecond)

	// play tone at 220Hz for 200ms...
	must(speaker.Tone(220))
	time.Sleep(200 * time.Millisecond)

	// then stop tone playback.
	must(speaker.Tone(0))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
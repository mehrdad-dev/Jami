package sound

import (
	"bufio"
	"log"
	"strings"
	"time"

	"github.com/dbatbold/beep"
)

// PlayNotes play a string of notes
func PlayNotes(notes string) {
	music := beep.NewMusic("")
	volume := 100

	if err := beep.OpenSoundDevice("default"); err != nil {
		log.Fatal(err)
	}

	if err := beep.InitSoundDevice(); err != nil {
		log.Fatal(err)
	}

	//beep.PrintSheet = true
	defer beep.CloseSoundDevice()

	reader := bufio.NewReader(strings.NewReader(notes))

	go music.Play(reader, volume)
	music.Wait()
	beep.FlushSoundBuffer()
}

// SaveNotes save a sound of notes
func SaveNotes(notes string) {
	name := time.Now().String()
	music := beep.NewMusic("Jami-" + name + ".wav")
	volume := 100

	if err := beep.OpenSoundDevice("default"); err != nil {
		log.Fatal(err)
	}

	if err := beep.InitSoundDevice(); err != nil {
		log.Fatal(err)
	}

	//beep.PrintSheet = true
	defer beep.CloseSoundDevice()

	reader := bufio.NewReader(strings.NewReader(notes))

	go music.Play(reader, volume)
	music.Wait()
	beep.FlushSoundBuffer()
}

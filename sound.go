package main

import (
	"embed"
	"fmt"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

//go:embed sound/time_end.mp3
var soundFS embed.FS

// InitSound decodes the sound file and inits the speaker with required sample rate
func InitSound() {
	fmt.Println("InitSound()")

	streamer, format := loadSoundFile()
	defer streamer.Close()

	err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		fmt.Printf("InitSound() Error while initializing the speaker: %v", err)
	}
}

// PlaySound plays the sound through the speaker
func PlaySound() {
	fmt.Println("PlaySound()")
	streamer, _ := loadSoundFile()
	defer streamer.Close()
	defer fmt.Println("PlaySound() Finished playing. deferred streamer closed")

	speaker.PlayAndWait(streamer)
}

// loadSoundFile loads file from fs, decodes it and returns its streamer and format
func loadSoundFile() (beep.StreamSeekCloser, *beep.Format) {
	fmt.Println("loadSoundFile()")
	f, err := soundFS.Open("sound/time_end.mp3")
	if err != nil {
		fmt.Printf("loadSoundFile() Error while opening the file: %v", err)
		return nil, nil
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Printf("loadSoundFile() Error while decoding the file: %v", err)
		return nil, nil
	}

	fmt.Printf("loadSoundFile() Successfully decoded, sample rate: %v \n", format.SampleRate)

	return streamer, &format
}

package engine

import (
	"bytes"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

func InitSoundPlayer(soundFile string, audioContext *audio.Context) (*audio.Player, error) {
	soundBytes, err := os.ReadFile(soundFile)
	if err != nil {
		return nil, err
	}

	streamer, err := mp3.Decode(audioContext, bytes.NewReader(soundBytes))
	if err != nil {
		return nil, err
	}

	audioPlayer, err := audio.NewPlayer(audioContext, streamer)
	if err != nil {
		return nil, err
	}

	return audioPlayer, nil
}

package engine

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type SoundFileType int

const (
	TypeWAV SoundFileType = iota
	TypeMP3
)

func InitSoundPlayer(soundFile string, soundFileType SoundFileType, audioContext *audio.Context) (*audio.Player, error) {
	soundBytes, err := os.ReadFile(soundFile)
	if err != nil {
		return nil, err
	}

	var streamer io.Reader

	switch soundFileType {
	case TypeMP3:
		streamer, err = mp3.Decode(audioContext, bytes.NewReader(soundBytes))
		if err != nil {
			return nil, err
		}

	case TypeWAV:
		streamer, err = wav.Decode(audioContext, bytes.NewReader(soundBytes))
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Unknown sound file type")
	}

	audioPlayer, err := audio.NewPlayer(audioContext, streamer)
	if err != nil {
		return nil, err
	}

	return audioPlayer, nil
}

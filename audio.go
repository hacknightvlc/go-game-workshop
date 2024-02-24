package main

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

const (
	sampleRate = 48000
)

type Audio struct {
	audioContext *audio.Context
	punchPlayer  *audio.Player
	musicPlayer  *audio.Player
}

func NewAudio() *Audio {
	audioContext := audio.NewContext(sampleRate)
	punch, err := mp3.DecodeWithoutResampling(bytes.NewReader(Punch_mp3))
	if err != nil {
		log.Fatal(err)
	}
	punchPlayer, err := audioContext.NewPlayer(punch)
	if err != nil {
		log.Fatal(err)
	}

	music, err := mp3.DecodeWithoutResampling(bytes.NewReader(Music_mp3))
	if err != nil {
		log.Fatal(err)
	}
	musicPlayer, err := audioContext.NewPlayer(music)
	if err != nil {
		log.Fatal(err)
	}
	return &Audio{
		audioContext: audioContext,
		punchPlayer:  punchPlayer,
		musicPlayer:  musicPlayer,
	}
}

func (a *Audio) PlaySound() error {
	if err := a.punchPlayer.Rewind(); err != nil {
		return err
	}
	a.punchPlayer.Play()
	return nil
}

func (a *Audio) PlayMusic() error {
	if err := a.musicPlayer.Rewind(); err != nil {
		return err
	}
	a.musicPlayer.Play()
	return nil
}

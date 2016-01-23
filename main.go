package main

import (
	"io"
	"os"

	"github.com/Aorioli/chopher/karplus.go"
	"github.com/Aorioli/chopher/note"
	"github.com/Aorioli/chopher/scale"
	"github.com/Aorioli/chopher/wave"
)

func main() {
	f, _ := os.Create("chopher.wav")
	w := wave.New(wave.Stereo, 44100)
	s := scale.Major.Scale(note.Note{Note: note.C, Octave: 2})
	for _, n := range s.Notes {
		w.Write(karplus.Sound(n.Frequency(), 1, 44100, 1000, 1))
	}
	io.Copy(f, w.Reader())
	f.Close()
}

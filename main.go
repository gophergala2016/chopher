package main

import (
	"io"
	"os"

	"github.com/gophergala2016/chopher/karplus"
	"github.com/gophergala2016/chopher/note"
	"github.com/gophergala2016/chopher/song"
	"github.com/gophergala2016/chopher/wave"
)

func main() {
	f, _ := os.Create("chopher.wav")
	w := wave.New(wave.Stereo, 44100)
	s := song.New(song.Medium)
	s.AddAfter(note.Note{Note: note.C, Octave: 2}, note.Half).
		AddAfter(note.Note{Note: note.DIS, Octave: 2}, note.Half).
		AddAfter(note.Note{Note: note.F, Octave: 2}, note.Half).
		AddAfter(note.Note{Note: note.C, Octave: 2}, note.Half).
		AddAfter(note.Note{Note: note.DIS, Octave: 2}, note.Half).
		AddAfter(note.Note{Note: note.FIS, Octave: 2}, note.Quarter).
		AddAfter(note.Note{Note: note.F, Octave: 2}, note.Half).
		AddAfter(note.Note{Note: note.C, Octave: 2}, note.Half).
		AddAfter(note.Note{Note: note.DIS, Octave: 2}, note.Half).
		AddAfter(note.Note{Note: note.F, Octave: 2}, note.Half).
		AddAfter(note.Note{Note: note.DIS, Octave: 2}, note.Half).
		AddAfter(note.Note{Note: note.C, Octave: 2}, note.Half)
	for _, n := range s.Notes {
		w.Write(karplus.Sound(n.Note.Frequency(), 0.996, 44100, 1000, float64(n.Duration)*2.0))
	}
	io.Copy(f, w.Reader())
	f.Close()
}

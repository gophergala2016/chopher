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
	s.AddAfter(note.Note{Note: note.C, Octave: 3}, note.Half).
		AddAfter(note.Note{Note: note.DIS, Octave: 3}, note.Half).
		AddAfter(note.Note{Note: note.F, Octave: 3}, note.Half).
		AddAfter(note.Note{Note: note.C, Octave: 3}, note.Half).
		AddAfter(note.Note{Note: note.DIS, Octave: 3}, note.Half).
		AddAfter(note.Note{Note: note.FIS, Octave: 3}, note.Quarter).
		AddAfter(note.Note{Note: note.F, Octave: 3}, note.Half).
		AddAfter(note.Note{Note: note.C, Octave: 3}, note.Half).
		AddAfter(note.Note{Note: note.DIS, Octave: 3}, note.Half).
		AddAfter(note.Note{Note: note.F, Octave: 3}, note.Half).
		AddAfter(note.Note{Note: note.DIS, Octave: 3}, note.Half).
		AddAfter(note.Note{Note: note.C, Octave: 3}, note.Half)

	ks := karplus.Song{
		Song:         s,
		SamplingRate: 44100,
	}

	data := ks.Sound()
	w.Write(data)
	io.Copy(f, w.Reader())
	f.Close()
}

package main

import (
	"io"
	"os"

	"github.com/gophergala2016/chopher/karplus"
	"github.com/gophergala2016/chopher/note"
	"github.com/gophergala2016/chopher/scale"
	"github.com/gophergala2016/chopher/song"
	"github.com/gophergala2016/chopher/wave"
)

func main() {
	f, _ := os.Create("chopher.wav")
	w := wave.New(wave.Stereo, 44100)
	s := song.New(song.Medium)
	s.Scale = scale.Major.New(note.Note{Note: note.C, Octave: 5})
	s.Add(note.Note{Note: note.C, Octave: 3}, note.Full).
		Add(note.Note{Note: note.E, Octave: 3}, note.Full*2.0).
		Add(note.Note{Note: note.G, Octave: 3}, note.Full).
		Add(note.Note{Note: note.CIS, Octave: 3}, note.Full).
		Add(note.Note{Note: note.F, Octave: 3}, note.Full).
		Add(note.Note{Note: note.GIS, Octave: 3}, note.Full).
		Add(note.Note{Note: note.D, Octave: 3}, note.Full).
		Add(note.Note{Note: note.FIS, Octave: 3}, note.Full).
		Add(note.Note{Note: note.A, Octave: 3}, note.Full)

	// s.AddAfter(note.Note{Note: note.C, Octave: 2}, note.Half).
	// 	AddWith(note.Note{Note: note.C, Octave: 2}.AddHalfSteps(5), note.Half).
	// 	AddAfter(note.Note{Note: note.DIS, Octave: 2}, note.Half).
	// 	AddWith(note.Note{Note: note.DIS, Octave: 2}.AddHalfSteps(5), note.Half).
	// 	AddAfter(note.Note{Note: note.F, Octave: 2}, note.Full).
	// 	AddWith(note.Note{Note: note.F, Octave: 2}.AddHalfSteps(5), note.Full).
	// 	AddAfter(note.Note{Note: note.C, Octave: 2}, note.Half).
	// 	AddWith(note.Note{Note: note.C, Octave: 2}.AddHalfSteps(5), note.Half).
	// 	AddAfter(note.Note{Note: note.DIS, Octave: 2}, note.Half).
	// 	AddWith(note.Note{Note: note.DIS, Octave: 2}.AddHalfSteps(5), note.Half).
	// 	AddAfter(note.Note{Note: note.FIS, Octave: 2}, note.Quarter).
	// 	AddWith(note.Note{Note: note.FIS, Octave: 2}.AddHalfSteps(5), note.Quarter).
	// 	AddAfter(note.Note{Note: note.F, Octave: 2}, note.Full).
	// 	AddWith(note.Note{Note: note.F, Octave: 2}.AddHalfSteps(5), note.Full).
	// 	AddAfter(note.Note{Note: note.C, Octave: 2}, note.Half).
	// 	AddWith(note.Note{Note: note.C, Octave: 2}.AddHalfSteps(5), note.Half).
	// 	AddAfter(note.Note{Note: note.DIS, Octave: 2}, note.Half).
	// 	AddWith(note.Note{Note: note.DIS, Octave: 2}.AddHalfSteps(5), note.Half).
	// 	AddAfter(note.Note{Note: note.F, Octave: 2}, note.Full).
	// 	AddWith(note.Note{Note: note.F, Octave: 2}.AddHalfSteps(5), note.Full).
	// 	AddWith(note.Note{Note: note.Rest}, note.Full).
	// 	AddAfter(note.Note{Note: note.DIS, Octave: 2}, note.Half).
	// 	AddWith(note.Note{Note: note.DIS, Octave: 2}.AddHalfSteps(5), note.Half).
	// 	AddAfter(note.Note{Note: note.C, Octave: 2}, note.Full).
	// 	AddWith(note.Note{Note: note.C, Octave: 2}.AddHalfSteps(5), note.Full)

	ks := karplus.Song{
		Song:         s,
		SamplingRate: 44100,
	}

	data := ks.Sound()
	w.Write(data)
	io.Copy(f, w.Reader())
	f.Close()
}

package main

import (
	"io"
	"os"

	"github.com/gophergala2016/chopher/hasher"
	"github.com/gophergala2016/chopher/karplus"
	"github.com/gophergala2016/chopher/wave"
)

func main() {
	fRead, _ := os.Open("main.go")
	fWrite, _ := os.Create("choper.wav")
	w := wave.New(wave.Stereo, 44000)

	// s.Add(note.Note{Note: note.C, Octave: 3}, note.Full).
	// 	Add(note.Note{Note: note.E, Octave: 3}, note.Full*2.0).
	// 	Add(note.Note{Note: note.G, Octave: 3}, note.Full).
	// 	Add(note.Note{Note: note.CIS, Octave: 3}, note.Full).
	// 	Add(note.Note{Note: note.F, Octave: 3}, note.Full).
	// 	Add(note.Note{Note: note.GIS, Octave: 3}, note.Full).
	// 	Add(note.Note{Note: note.D, Octave: 3}, note.Full).
	// 	Add(note.Note{Note: note.FIS, Octave: 3}, note.Full).
	// 	Add(note.Note{Note: note.A, Octave: 3}, note.Full)
	//
	// // s.AddAfter(note.Note{Note: note.C, Octave: 2}, note.Half).
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
	h := hasher.New(fRead)
	defer fRead.Close()

	sng := h.Hash()
	ks := karplus.Song{
		Song:         *sng,
		SamplingRate: 44000,
	}
	ks.Sound(&w)

	io.Copy(fWrite, w.Reader())
	fWrite.Close()

}

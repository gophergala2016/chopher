package note

import (
	"encoding/binary"
	"fmt"
	"math"
)

const (
	a4Frequency float64 = 440.0
	magic       float64 = 1.0594630943592952645618252949463417007792043174941856
)

// Musical scale halfstep values
const (
	C = iota
	CIS
	D
	DIS
	E
	F
	FIS
	G
	GIS
	A
	AIS
	B
)

var (
	a4 = Note{
		Note:   A,
		Octave: 4,
	}
)

// Note represents a piano key note with a combination of a note and it's Octave
type Note struct {
	Note   int
	Octave int
}

// AddHalfSteps adds a number of half steps and returns a new note
func (n Note) AddHalfSteps(hs int) Note {
	t := n.Note + hs
	n.Note = t % 12
	n.Octave = n.Octave + t/12
	return n
}

// Frequency of the note
//
// https://en.wikipedia.org/wiki/Piano_key_frequencies
func (n Note) Frequency() float64 {
	return a4Frequency * math.Pow(magic, float64(halfStepDistance(a4, n)))
}

func halfStepDistance(from, to Note) int {
	return (to.Octave-from.Octave)*12 + (to.Note - from.Note)
}

var notes = [...]string{
	"C", "C#", "D", "D#",
	"E", "F", "F#", "G",
	"G#", "A", "A#", "B",
}

// Sound plays the note into a byte slice
func (n Note) Sound(samplesPerSecond, duration, volume int) []byte {
	numberOfSamples := samplesPerSecond * duration
	phase := 0.0

	var frequencyRadiansPerSample = n.Frequency() * 2 * math.Pi / float64(samplesPerSecond)
	ret := make([]byte, 0, numberOfSamples*2)
	buf := make([]byte, 2, 2)
	for sample := 0; sample < numberOfSamples; sample++ {
		phase += frequencyRadiansPerSample
		sampleValue := float64(volume) * math.Sin(phase)
		binary.LittleEndian.PutUint16(buf, uint16(sampleValue))
		ret = append(ret, buf...)
	}
	return ret
}

func (n Note) String() string {
	return fmt.Sprintf("%s%d", notes[n.Note], n.Octave)
}

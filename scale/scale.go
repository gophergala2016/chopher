package scale

import (
	"strings"

	"github.com/gophergala2016/chopher/note"
)

type Chord []int

func (c Chord) NotesInChord(base note.Note, check note.Note, pos int) bool {
	return pos < len(c) && note.HalfstepDistance(base, check) == c[pos]
}

// Pattern is an array of halfstep jumps
type Pattern struct {
	Name   string
	Scale  []int
	Chords []Chord
}

var (
	// Major scale pattern
	Major = Pattern{
		Name:  "Major",
		Scale: []int{2, 2, 1, 2, 2, 2, 1},
		Chords: []Chord{
			Chord{0, 4, 7, 11}, // seventh
			Chord{0, 4, 7},     // major
			Chord{0, 7},        // power
		},
	}
	// Minor scale pattern
	Minor = Pattern{
		Name:  "Minor",
		Scale: []int{2, 1, 2, 2, 1, 2, 2},
		Chords: []Chord{
			Chord{0, 3, 7, 10}, // seventh
			Chord{0, 3, 7},     // minor
			Chord{0, 3, 6},     // diminished
		},
	}
)

// Scale is defined by the notes and the pattern they form
type Scale struct {
	Notes  []note.Note
	Chords []Chord
}

// Scale turn a pattern to a scale using a key note
func (p Pattern) New(key note.Note) Scale {
	n := make([]note.Note, len(p.Scale)+1)
	n[0] = key
	for i, v := range p.Scale {
		n[i+1] = n[i].AddHalfSteps(v)
	}
	return Scale{
		Notes:  n,
		Chords: p.Chords,
	}
}

func (s Scale) String() string {
	st := make([]string, len(s.Notes))
	for i, n := range s.Notes {
		st[i] = n.String()
	}
	return strings.Join(st, "-")
}

package scale

import (
	"strings"

	"github.com/gophergala2016/chopher/note"
)

// Pattern is an array of halfstep jumps
type Pattern []int

var (
	// Major scale pattern
	Major Pattern = []int{2, 2, 1, 2, 2, 2, 1}
	// Minor scale pattern
	Minor Pattern = []int{2, 1, 2, 2, 1, 2, 2}
)

// Scale is defined by the notes and the pattern they form
type Scale struct {
	Notes   []note.Note
	pattern Pattern
}

// Scale turn a pattern to a scale using a key note
func (p Pattern) New(key note.Note) Scale {
	n := make([]note.Note, len(p)+1)
	n[0] = key
	for i, v := range p {
		n[i+1] = n[i].AddHalfSteps(v)
	}
	return Scale{
		Notes:   n,
		pattern: p,
	}
}

func (s Scale) String() string {
	st := make([]string, len(s.Notes))
	for i, n := range s.Notes {
		st[i] = n.String()
	}
	return strings.Join(st, "-")
}

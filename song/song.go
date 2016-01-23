package song

import "github.com/gophergala2016/chopher/note"

const (
	Slow   Tempo = 2
	Medium Tempo = 1
	Fast   Tempo = 0.5
)

// Tempo is the duration of the Full note in seconds
//
// Not my tempo!!
type Tempo float64

type SongNote struct {
	Note     note.Note
	Duration note.Duration
	Start    float64
}

type Song struct {
	Tempo Tempo
	Notes []SongNote
}

func New(t Tempo) Song {
	return Song{
		Tempo: t,
	}
}

func (s *Song) add(note note.Note, duration note.Duration, start float64) {
	s.Notes = append(s.Notes, SongNote{
		Note:     note,
		Duration: duration,
		Start:    start,
	})
}

func (s *Song) AddAfter(note note.Note, duration note.Duration) *Song {
	lastNote := SongNote{}
	if len(s.Notes) > 0 {
		lastNote = s.Notes[len(s.Notes)-1]
	}
	s.add(note, duration, lastNote.Start+float64(lastNote.Duration))
	return s
}

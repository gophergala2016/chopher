package song

import "github.com/gophergala2016/chopher/note"

const (
	Fast   Tempo = 1.2
	Medium Tempo = 0.8
	Slow   Tempo = 0.5
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

func (s *SongNote) IsValid(time float64) bool {
	// log.Println(s.Note, time, s.Start, s.Start+float64(s.Duration))
	if time < s.Start {
		return false
	}
	return time < (s.Start + float64(s.Duration))
}

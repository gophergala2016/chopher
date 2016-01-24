package karplus

import (
	"testing"

	"github.com/gophergala2016/chopher/note"
	"github.com/gophergala2016/chopher/song"
)

type FakeWriter struct{}

func (f FakeWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

func BenchmarkSong(b *testing.B) {
	s := song.New(song.Tempo(1.0))
	s.AddAfter(note.Note{Note: note.E, Octave: 4}, note.Full)
	ks := Song{
		Song:         s,
		SamplingRate: 44100,
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ks.Sound(FakeWriter{})
	}
}

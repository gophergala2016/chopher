package karplus

import (
	"bytes"
	"encoding/binary"
	"math"
	"math/rand"

	"github.com/gophergala2016/chopher/song"
)

// Sound plays a sound generated with the Karplus-Strong algorithm
func Sound(frequency float64, damping float64, samplesPerSecond, volume int, duration float64) []byte {
	buf := make([]float64, int(
		math.Ceil(
			float64(samplesPerSecond)/frequency,
		),
	))
	r := rand.New(rand.NewSource(1))

	for i := 0; i < len(buf); i++ {
		buf[i] = r.Float64() - 0.5
	}

	ret := make([]byte, 0, len(buf)*2)
	j := 0
	for i := 0; i < int(float64(samplesPerSecond)*duration); i++ {
		var temp2 bytes.Buffer
		sampleValue := int16(buf[j] * 32767)
		binary.Write(&temp2, binary.LittleEndian, sampleValue)
		ret = append(ret, temp2.Bytes()...)

		nextJ := (j + 1) % len(buf)
		buf[j] = (buf[j] + buf[nextJ]) * 0.5 * damping
		j = nextJ
	}
	return ret
}

type Note struct {
	Note   song.SongNote
	Buffer []float64
}

func NewNote(n song.SongNote, samplingRate int) *Note {
	buf := make([]float64, int(
		math.Ceil(
			float64(samplingRate)/n.Note.Frequency(),
		),
	))
	r := rand.New(rand.NewSource(1))

	for i := 0; i < len(buf); i++ {
		buf[i] = r.Float64() - 0.5
	}
	return &Note{
		Note:   n,
		Buffer: buf,
	}
}

// Sound pops the current buffer value and appends the new one
func (n *Note) Sound() float64 {
	sampleValue := n.Buffer[0]

	v := (n.Buffer[0] + n.Buffer[1]) * 0.5 * 0.99996
	n.Buffer = append(n.Buffer[1:], v)

	return sampleValue
}

type Song struct {
	Song         song.Song
	SamplingRate int
	CurrentNotes []*Note
}

func (s *Song) Sound() []byte {
	var lastNote int
	for i, n := range s.Song.Notes {
		if n.IsValid(0) {
			s.CurrentNotes = append(s.CurrentNotes, NewNote(n, s.SamplingRate))
			lastNote = i
		}
	}

	var (
		time      float64
		increment = float64(s.Song.Tempo) / float64(s.SamplingRate)
	)

	ret := make([]byte, 0, len(s.Song.Notes)*s.SamplingRate)
	var orderBuffer bytes.Buffer
	for len(s.CurrentNotes) > 0 {
		var sample float64
		for _, n := range s.CurrentNotes {
			sample += n.Sound()
		}

		time += increment
		temp := make([]*Note, 0, len(s.CurrentNotes))
		for _, n := range s.CurrentNotes {
			if n.Note.IsValid(time) {
				temp = append(temp, n)
			}
		}

		for i := lastNote + 1; i < len(s.Song.Notes); i++ {
			n := s.Song.Notes[i]
			if n.IsValid(time) {
				temp = append(s.CurrentNotes, NewNote(n, s.SamplingRate))
				lastNote = i
			}
		}

		s.CurrentNotes = temp

		orderBuffer = bytes.Buffer{}
		sampleValue := int16(sample * 32767)
		binary.Write(&orderBuffer, binary.LittleEndian, sampleValue)
		ret = append(ret, orderBuffer.Bytes()...)
	}
	return ret
}

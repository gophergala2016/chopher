package scale

import (
	"testing"

	"github.com/gophergala2016/chopher/note"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPatternNew(t *testing.T) {
	Convey("Given pattern", t, func() {
		Convey("Major", func() {
			actual := Major.New(note.Note{Note: note.C, Octave: 4}, false)
			expected := Scale{
				Notes: []note.Note{
					note.Note{Note: note.C, Octave: 4},
					note.Note{Note: note.D, Octave: 4},
					note.Note{Note: note.E, Octave: 4},
					note.Note{Note: note.F, Octave: 4},
					note.Note{Note: note.G, Octave: 4},
					note.Note{Note: note.A, Octave: 4},
					note.Note{Note: note.B, Octave: 4},
					note.Note{Note: note.C, Octave: 5},
				},
				Chords: Major.Chords,
			}
			So(actual, ShouldResemble, expected)
		})
	})
}

func TestString(t *testing.T) {
	Convey("Given value", t, func() {
		actual := Major.New(note.Note{Note: note.C, Octave: 4}, false).String()
		expected := "C4-D4-E4-F4-G4-A4-B4-C5"
		So(actual, ShouldEqual, expected)
	})
}

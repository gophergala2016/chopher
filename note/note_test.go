package note

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestAddHalfStep(t *testing.T) {
	convey.Convey("Given value", t, func() {
		n := Note{Note: A, Octave: 4}
		expected := Note{Note: C, Octave: 5}

		convey.So(n.AddHalfSteps(3), convey.ShouldResemble, expected)
	})
}

func TestHalfStepDistance(t *testing.T) {
	convey.Convey("Given value", t, func() {
		expected := 3

		convey.So(halfStepDistance(a4, Note{Note: C, Octave: 5}), convey.ShouldEqual, expected)
	})
}

func TestFrequency(t *testing.T) {
	convey.Convey("Given value", t, func() {
		actual := Note{Note: C, Octave: 5}.Frequency()

		convey.So(actual, convey.ShouldAlmostEqual, 523.251, 0.001)
	})
}

func TestString(t *testing.T) {
	convey.Convey("Given value", t, func() {
		expected := "C5"
		convey.So(Note{Note: C, Octave: 5}.String(), convey.ShouldEqual, expected)
	})
}

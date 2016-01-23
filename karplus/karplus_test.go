package karplus

import "testing"

func BenchmarkSound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sound(440, 0.996, 44200, 1, 60)
	}
}

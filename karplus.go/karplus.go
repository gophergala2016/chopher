package karplus

import (
	"bytes"
	"encoding/binary"
	"math"
	"math/rand"
)

// Sound plays a sound generated with the Karplus-Strong algorithm
func Sound(frequency float64, damping float64, samplesPerSecond, volume, duration int) []byte {
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
	for i := 0; i < samplesPerSecond*duration; i++ {
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

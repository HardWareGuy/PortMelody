package portmelody

import (
	"code.google.com/p/portaudio-go/portaudio"
	"math"
	"time"
)

const sampleRate = 44100

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()
	outstream := newStereoSine(320, 320, sampleRate)
	defer outstream.Close()
	chk(outstream.Start())
	time.Sleep(2 * time.Second)
	chk(outstream.Stop())
}

type stereoSine struct {
	*portaudio.Stream
	stepL, phaseL float64
	stepR, phaseR float64
}

func newStereoSine(freqL, freqR, sampleRate float64) *stereoSine {
	outstream := &stereoSine{nil, freqL / sampleRate, 0, freqR / sampleRate, 0}
	var err error
	outstream.Stream, err = portaudio.OpenDefaultStream(0, 2, sampleRate, 0, outstream.processAudio)
	chk(err)
	return outstream
}

func (g *stereoSine) processAudio(out [][]float32) {
	for i := range out[0] {
		out[0][i] = float32(math.Sin(2*math.Pi*g.phaseL) * 0.5)
		_, g.phaseL = math.Modf(g.phaseL + g.stepL)
		out[1][i] = float32(math.Sin(2*math.Pi*g.phaseR) * 0.5)
		_, g.phaseR = math.Modf(g.phaseR + g.stepR)
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

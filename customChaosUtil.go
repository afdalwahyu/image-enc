package main

// ChaosKey Key that will be as initial value
import (
	"errors"
	"image"
	"math"
)

type ChaosKey struct {
	X0 float64
	u  float64
	k  float64
	N0 int
	lp int
}

// ChaosPixel Type value of chaos with pixel
type ChaosPixel struct {
	pixel byte
	chaos float64
	index int
}

// ChaosPixelSequence Type of sequence that wll be generated from logistic2 map sequence
type ChaosPixelSequence []ChaosPixel

// ByChaos sort based on chaos value
type ByChaos []ChaosPixel

func (c ByChaos) Len() int           { return len(c) }
func (c ByChaos) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByChaos) Less(i, j int) bool { return c[i].chaos < c[j].chaos }

// NewChaosKey create chaos key
// TODO: k still unknown?
func NewChaosKey(bounds *image.Rectangle, N0 int, X0 float64, u float64, k float64, lp int) (*ChaosKey, error) {

	maxLen := 3 * bounds.Max.X * bounds.Max.Y

	if N0 < 0 {
		return nil, errors.New("Cannot use N0 below 0")
	}

	if u < 0 || u > 10 {
		return nil, errors.New("U value error, must be above zero and below 10")
	}

	if X0 < 0 {
		return nil, errors.New("X0 as population cannot be below 10")
	}

	if lp < 1 || lp > maxLen {
		return nil, errors.New("lp key length chaos sequence invalid")
	}

	return &ChaosKey{
		X0: X0,
		u:  u,
		k:  k,
		N0: N0,
		lp: lp,
	}, nil
}

func splitChiperToRGB(buf []byte, alphaChannel []byte) arrayColor {
	lim := len(buf) / 3

	return arrayColor{
		red:   buf[:lim],
		green: buf[lim : 2*lim],
		blue:  buf[2*lim:],
		alpha: alphaChannel,
	}
}

func (key *ChaosKey) generateLogisticLogisticMapSequence(length int) []float64 {
	maxLength := length + key.N0
	logisticSequence := make([]float64, maxLength)

	logisticSequence[0] = key.X0
	for i := 1; i < maxLength; i++ {
		firstLogistic := key.u * logisticSequence[i-1] * (1 - logisticSequence[i-1]) * math.Pow(2, 14)
		secondLogistic := math.Floor(firstLogistic)
		logisticSequence[i] = firstLogistic - secondLogistic
	}

	return logisticSequence
}

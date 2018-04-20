package chaos

// Key Key that will be as initial value
import (
	"errors"
	"image"
	"math"
	"skripsi/util"
)

// Key variable that use to encrypt / decrypt
type Key struct {
	X0       float64
	u        float64
	k        float64
	N0       int
	lp       int
	sequence string
}

// Pixel Type value of chaos with pixel
type Pixel struct {
	pixel byte
	chaos float64
	index int
}

// PixelSequence Type of sequence that wll be generated from logistic2 map sequence
type PixelSequence []Pixel

// ByChaos sort based on chaos value
type ByChaos []Pixel

func (c ByChaos) Len() int           { return len(c) }
func (c ByChaos) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByChaos) Less(i, j int) bool { return c[i].chaos < c[j].chaos }

// NewChaosKey create chaos 2048key
func NewChaosKey(bounds *image.Rectangle, N0 int, X0 float64, u float64, k int64, lp int, sequence string) (*Key, error) {
	maxLen := 3 * bounds.Max.X * bounds.Max.Y

	if k < 8 || k > 20 {
		return nil, errors.New("k must be higher than 8 and below 20")
	}

	if N0 < 0 {
		return nil, errors.New("cannot use N0 below 0")
	}

	if u < 0 || u > 10 {
		return nil, errors.New("U value error, must be above zero and below 10")
	}

	if X0 < 0 {
		return nil, errors.New("x0 as population cannot be below 10")
	}

	if lp < 1 || lp > maxLen {
		return nil, errors.New("lp 2048key length chaos sequence invalid")
	}

	return &Key{
		X0:       X0,
		u:        u,
		k:        float64(k),
		N0:       N0,
		lp:       lp,
		sequence: sequence,
	}, nil
}

func splitCipherToRGB(buf []byte, alphaChannel []byte) util.ArrayColor {
	lim := len(buf) / 3

	return util.ArrayColor{
		Red:   buf[:lim],
		Green: buf[lim : 2*lim],
		Blue:  buf[2*lim:],
		Alpha: alphaChannel,
	}
}

func (key *Key) generateLogisticLogisticMapSequence(length int) []float64 {
	maxLength := length + key.N0
	logisticSequence := make([]float64, maxLength)

	logisticSequence[0] = key.X0
	for i := 1; i < maxLength; i++ {
		firstLogistic := key.u * logisticSequence[i-1] * (1 - logisticSequence[i-1]) * math.Pow(2, key.k)
		secondLogistic := math.Floor(firstLogistic)
		logisticSequence[i] = firstLogistic - secondLogistic
	}

	return logisticSequence[key.N0:]
}

func (key *Key) generateSineSineMapSequence(length int) []float64 {
	maxLength := length + key.N0
	sineSequence := make([]float64, maxLength)

	sineSequence[0] = key.X0
	for i := 1; i < maxLength; i++ {
		firstSine := key.u * math.Sin(math.Pi*sineSequence[i-1]) * math.Pow(2, key.k)
		secondSine := math.Floor(firstSine)
		sineSequence[i] = firstSine - secondSine
	}

	return sineSequence[key.N0:]
}

func (key *Key) generateChebyshevChebyshevMapSequence(length int) []float64 {
	maxLength := length + key.N0
	chebyshevSequence := make([]float64, maxLength)

	chebyshevSequence[0] = key.X0
	for i := 1; i < maxLength; i++ {
		firstChebyshev := math.Cos((key.u+1)*math.Acos(chebyshevSequence[i-1])) *
			math.Pow(2, key.k)
		secondChebyshev := math.Floor(firstChebyshev)
		chebyshevSequence[i] = firstChebyshev - secondChebyshev
	}

	return chebyshevSequence[key.N0:]
}

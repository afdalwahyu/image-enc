package main

import (
	"errors"
	"fmt"
	"image"
	"log"
	"math"
	"sort"
)

// ChaosKey Key that will be as initial value
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

// ChaosEncrypt Encrypt using chaos and logistic map
func ChaosEncrypt(bounds *image.Rectangle, c *arrayColor) {
	// Step 1: Get 3 color
	// Step 2: Concat 3 color to 1d array
	plainPixels := make([]byte, 0)
	plainPixels = append(plainPixels, c.red...)
	plainPixels = append(plainPixels, c.green...)
	plainPixels = append(plainPixels, c.blue...)

	// Step 3: Generate chaotic sequence
	//         key that used to generate: x0, u, k
	//         iterate sequence len(Step2)+N0 times
	//         get sequence from N0 to last as chaos sequence
	key, err := NewChaosKey(bounds, 1000, 0.5, 3.6, 0.0, 5)
	if err != nil {
		log.Fatal(err)
	}

	chaosSequence := key.generateLogisticLogisticMapSequence(len(plainPixels))[key.N0:]

	chaosPixels := make(ChaosPixelSequence, 0)
	for index, chaos := range chaosSequence {
		chaosPixels = append(chaosPixels, ChaosPixel{
			pixel: plainPixels[index],
			chaos: chaos,
		})
	}

	// Step 4: sort the chaos sequence
	// Step 5: move position byte image based on sorted sequence
	sort.Sort(ByChaos(chaosPixels))

	// Step 6: Generate diffusion matrix := mod(floor(Xi)*10^14, 256)
	diffusion := make([]byte, 0)

	for _, chaos := range chaosSequence {
		d := math.Mod(math.Floor(chaos)*math.Pow10(14), 256)
		diffusion = append(diffusion, uint8(d))
	}

	// Step 7: Generate chiper1 sequence by formula: C[i] = mod(P'[i]+D'[i], 256) xor C[i-1]
	chiper := make([]byte, 0)

	for index, cp := range chaosPixels {
		if index == 0 {
			c1 := uint8(math.Mod(float64(cp.pixel)+float64(diffusion[index]), 256))
			chiper = append(chiper, c1)
			continue
		}
		c1 := uint8(math.Mod(float64(cp.pixel)+float64(diffusion[index]), 256)) ^ chiper[index-1]
		chiper = append(chiper, c1)
	}

	// Step 8: Generate chiper2 sequence by rotating chiper1 to left
	tmp := chiper[:key.lp]
	chiper = chiper[key.lp:]
	chiper = append(chiper, tmp...)

	fmt.Println(len(chiper), len(diffusion), len(chaosPixels))
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

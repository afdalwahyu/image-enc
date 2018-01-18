package main

import (
	"image"
	"log"
	"math"
	"sort"
)

// chaosEncrypt Encrypt using chaos and logistic map
func chaosEncrypt(bounds *image.Rectangle, c *arrayColor) arrayColor {
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
		d := math.Mod(math.Floor(chaos*math.Pow10(14)), 256)
		diffusion = append(diffusion, uint8(d))
	}

	// Step 7: Generate chiper1 sequence by formula: C[i] = mod(P'[i]+D'[i], 256) xor C[i-1]
	chiper := make([]byte, 0)

	for index, cp := range chaosPixels {
		var c1 uint8
		if index == 0 {
			c1 = uint8(math.Mod(float64(cp.pixel+diffusion[index]), 256))
		} else {
			c1 = uint8(math.Mod(float64(cp.pixel+diffusion[index]), 256)) ^ chiper[index-1]
		}
		chiper = append(chiper, c1)
	}

	// Step 8: Generate chiper2 sequence by rotating chiper1 to left
	tmp := chiper[:key.lp]
	chiper = chiper[key.lp:]
	chiper = append(chiper, tmp...)

	encryptedColors := splitChiperToRGB(chiper, c.alpha)

	return encryptedColors
}

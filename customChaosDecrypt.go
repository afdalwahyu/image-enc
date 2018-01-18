package main

import (
	"image"
	"log"
	"math"
	"sort"
)

func chaosDecrypt(bounds *image.Rectangle, c *arrayColor) arrayColor {
	chiperPixels := make([]byte, 0)
	chiperPixels = append(chiperPixels, c.red...)
	chiperPixels = append(chiperPixels, c.green...)
	chiperPixels = append(chiperPixels, c.blue...)

	key, err := NewChaosKey(bounds, 1000, 0.5, 3.6, 0.0, 5)
	if err != nil {
		log.Fatal(err)
	}

	tmp2 := chiperPixels[:len(chiperPixels)-key.lp]
	chiperPixels = chiperPixels[len(chiperPixels)-key.lp:]
	chiperPixels = append(chiperPixels, tmp2...)

	chaosSequence := key.generateLogisticLogisticMapSequence(len(chiperPixels))[key.N0:]

	diffusion := make([]byte, 0)
	for _, chaos := range chaosSequence {
		d := math.Mod(math.Floor(chaos*math.Pow10(14)), 256)
		diffusion = append(diffusion, uint8(d))
	}

	chaosPlainPixels := make(ChaosPixelSequence, 0)
	for index, chaos := range chaosSequence {
		var p uint8
		if index == 0 {
			p = uint8(math.Mod(float64(chiperPixels[index]-diffusion[index]), 256))
		} else {
			p = uint8(math.Mod(float64(chiperPixels[index]^chiperPixels[index-1]-diffusion[index]), 256))
		}
		chaosPlainPixels = append(chaosPlainPixels, ChaosPixel{
			pixel: p,
			chaos: chaos,
			index: index,
		})
	}

	tmpPixel := make([]byte, 0)
	for _, p := range chaosPlainPixels {
		tmpPixel = append(tmpPixel, p.pixel)
	}

	sort.Sort(ByChaos(chaosPlainPixels))

	decryptedPixels := make([]byte, len(chaosPlainPixels))
	for index, pixelC := range chaosPlainPixels {
		decryptedPixels[pixelC.index] = tmpPixel[index]
	}

	decryptedColors := splitChiperToRGB(decryptedPixels, c.alpha)
	return decryptedColors
}

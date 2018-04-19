package chaos

import (
	"image"
	"math"
	"skripsi/util"
	"sort"
)

// ChaosEncrypt Encrypt using chaos and logistic map
func (key *Key) ChaosEncrypt(bounds *image.Rectangle, c *util.ArrayColor) util.ArrayColor {
	// Step 1: Get 3 color
	// Step 2: Concat 3 color to 1d array
	plainPixels := make([]byte, 0)
	plainPixels = append(plainPixels, c.Red...)
	plainPixels = append(plainPixels, c.Green...)
	plainPixels = append(plainPixels, c.Blue...)

	// Step 3: Generate chaotic sequence
	//         iterate sequence len(Step2)+N0 times
	//         get sequence from N0 to last as chaos sequence

	chaosSequence := key.generateLogisticLogisticMapSequence(len(plainPixels))

	chaosPixels := make(PixelSequence, 0)
	for index, chaos := range chaosSequence {
		chaosPixels = append(chaosPixels, Pixel{
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
	cipher := make([]byte, 0)

	for index, cp := range chaosPixels {
		var c1 uint8
		if index == 0 {
			c1 = uint8(math.Mod(float64(cp.pixel+diffusion[index]), 256))
		} else {
			c1 = uint8(math.Mod(float64(cp.pixel+diffusion[index]), 256)) ^ cipher[index-1]
		}
		cipher = append(cipher, c1)
	}

	// Step 8: Generate chiper2 sequence by rotating chiper1 to left
	tmp := cipher[:key.lp]
	cipher = cipher[key.lp:]
	cipher = append(cipher, tmp...)

	encryptedColors := splitCipherToRGB(cipher, c.Alpha)

	return encryptedColors
}

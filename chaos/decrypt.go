package chaos

import (
    "image"
    "math"
    "sort"
    "1-learn/util"
)

// ChaosDecrypt decrypt using chaos algorithm
func (key *Key) ChaosDecrypt(bounds *image.Rectangle, c *util.ArrayColor) util.ArrayColor {
    cipherPixels := make([]byte, 0)
    cipherPixels = append(cipherPixels, c.Red...)
    cipherPixels = append(cipherPixels, c.Green...)
    cipherPixels = append(cipherPixels, c.Blue...)

    //key, err := NewChaosKey(bounds, 1000, 0.5, 3.6, 0.0, 5)
    //if err != nil {
    //	log.Fatal(err)
    //}

    tmp2 := cipherPixels[:len(cipherPixels)-key.lp]
    cipherPixels = cipherPixels[len(cipherPixels)-key.lp:]
    cipherPixels = append(cipherPixels, tmp2...)

    chaosSequence := key.generateLogisticLogisticMapSequence(len(cipherPixels))[key.N0:]

    diffusion := make([]byte, 0)
    for _, chaos := range chaosSequence {
        d := math.Mod(math.Floor(chaos*math.Pow10(14)), 256)
        diffusion = append(diffusion, uint8(d))
    }

    chaosPlainPixels := make(PixelSequence, 0)
    for index, chaos := range chaosSequence {
        var p uint8
        if index == 0 {
            p = uint8(math.Mod(float64(cipherPixels[index]-diffusion[index]), 256))
        } else {
            p = uint8(math.Mod(float64(cipherPixels[index]^cipherPixels[index-1]-diffusion[index]), 256))
        }
        chaosPlainPixels = append(chaosPlainPixels, Pixel{
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

    decryptedColors := splitCipherToRGB(decryptedPixels, c.Alpha)
    return decryptedColors
}

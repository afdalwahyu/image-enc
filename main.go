package main

import (
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"log"
	"os"
)

func main() {

	bounds, redColor, greenColor, blueColor, alphaChannel := openImage("./image/eye.png")
	Red, Green, Blue, Alpha := RSAEncryptImage(redColor, greenColor, blueColor, alphaChannel)
	generateImage("./image/encrypted.png", &bounds, Red, Green, Blue, Alpha)

	// bounds, redColor, greenColor, blueColor, alphaChannel := openImage("./image/encrypted.png")
	// Red, Green, Blue, Alpha := RSADecryptImage(redColor, greenColor, blueColor, alphaChannel)
	// generateImage("./image/decrypted.png", &bounds, Red, Green, Blue, Alpha)
}

func openImage(imageFile string) (bounds image.Rectangle, redColor []byte, greenColor []byte, blueColor []byte, alphaChannel []byte) {
	reader, err := os.Open(imageFile)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	bounds = m.Bounds()

	// change color
	maxHeight := bounds.Max.Y
	maxWidth := bounds.Max.X
	maxLengthColor := maxHeight * maxWidth

	redColor = make([]byte, maxLengthColor)
	greenColor = make([]byte, maxLengthColor)
	blueColor = make([]byte, maxLengthColor)
	alphaChannel = make([]byte, maxLengthColor)

	index := 0
	for y := bounds.Min.Y; y < maxHeight; y++ {
		for x := bounds.Min.X; x < maxWidth; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			redColor[index] = byte(r)
			greenColor[index] = byte(g)
			blueColor[index] = byte(b)
			alphaChannel[index] = byte(a)
			index++
		}
	}

	return
}

func generateImage(imageName string, bounds *image.Rectangle, Red []byte, Green []byte, Blue []byte, Alpha []byte) {
	newImg := image.NewNRGBA(image.Rect(0, 0, bounds.Max.X, bounds.Max.Y))
	maxHeight := bounds.Max.Y
	maxWidth := bounds.Max.X

	for y := bounds.Min.Y; y < maxHeight; y++ {
		for x := bounds.Min.X; x < maxWidth; x++ {
			index := ((maxWidth - 1) * y) + (x + y)
			newImg.Set(x, y, color.NRGBA{
				R: Red[index],
				G: Green[index],
				B: Blue[index],
				A: Alpha[index],
			})
		}
	}

	f, err := os.Create(imageName)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, newImg); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

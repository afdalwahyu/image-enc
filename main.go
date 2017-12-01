package main

import (
	"image"
	"image/color"

	"golang.org/x/image/bmp"
	// "image/jpeg"
	// "image/png"
	"log"
	"os"
)

func main() {

	// bounds, redColor, greenColor, blueColor, alphaChannel := openImage("./image/eye.png")
	// Red, Green, Blue, Alpha := RSAEncryptImage(redColor, greenColor, blueColor, alphaChannel)
	// generateImage("./image/encrypted.png", &bounds, Red, Green, Blue, Alpha)

	// bounds, redColor, greenColor, blueColor, alphaChannel := openImage("./image/encrypted.png")
	// Red, Green, Blue, Alpha := RSADecryptImage(redColor, greenColor, blueColor, alphaChannel)
	// generateImage("./image/decrypted.png", &bounds, Red, Green, Blue, Alpha)

	bounds, inputPlainColor := openImage("./image/eye.bmp")
	encryptedColor := AESEncryptImage(inputPlainColor)
	generateImage("./image/EncryptedAES.bmp", &bounds, encryptedColor)

	bounds, inputColor2 := openImage("./image/EncryptedAES.bmp")
	decryptedColor := AESDecryptImage(inputColor2)
	generateImage("./image/DecryptedAES.png", &bounds, decryptedColor)
}

func openImage(imageFile string) (bounds image.Rectangle, c arrayColor) {
	reader, err := os.Open(imageFile)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, err := bmp.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	bounds = m.Bounds()

	// change color
	maxHeight := bounds.Max.Y
	maxWidth := bounds.Max.X
	maxLengthColor := maxHeight * maxWidth

	redColor := make([]byte, maxLengthColor)
	greenColor := make([]byte, maxLengthColor)
	blueColor := make([]byte, maxLengthColor)
	alphaChannel := make([]byte, maxLengthColor)

	index := 0
	for y := bounds.Min.Y; y < maxHeight; y++ {
		for x := bounds.Min.X; x < maxWidth; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			redColor[index] = uint8(r >> 8)
			greenColor[index] = uint8(g >> 8)
			blueColor[index] = uint8(b >> 8)
			alphaChannel[index] = uint8(a >> 8)
			index++
		}
	}

	c = arrayColor{
		red:   redColor,
		green: greenColor,
		blue:  blueColor,
		alpha: alphaChannel,
	}

	return
}

func generateImage(imageName string, bounds *image.Rectangle, c arrayColor) {
	newImg := image.NewNRGBA(image.Rect(0, 0, bounds.Max.X, bounds.Max.Y))
	maxHeight := bounds.Max.Y
	maxWidth := bounds.Max.X

	for y := bounds.Min.Y; y < maxHeight; y++ {
		for x := bounds.Min.X; x < maxWidth; x++ {
			index := ((maxWidth - 1) * y) + (x + y)
			newImg.Set(x, y, color.RGBA{
				R: c.red[index],
				G: c.green[index],
				B: c.blue[index],
				A: c.alpha[index],
			})
		}
	}

	f, err := os.Create(imageName)
	if err != nil {
		log.Fatal(err)
	}

	if err := bmp.Encode(f, newImg); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

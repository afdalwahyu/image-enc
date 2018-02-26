package util

import (
    "image"
    "os"
    "log"
    "golang.org/x/image/bmp"
    "image/color"
)

// ArrayColor Color of image
type ArrayColor struct {
    Red, Green, Blue, Alpha []byte
}

type ImageFile struct {
    Path string
}

// OpenImage open image file
func (img *ImageFile) OpenImage() (bounds image.Rectangle, c ArrayColor) {
    reader, err := os.Open(img.Path)
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
    for x := bounds.Min.X; x < maxWidth; x++ {
        for y := bounds.Min.Y; y < maxHeight; y++ {
            r, g, b, a := m.At(x, y).RGBA()
            redColor[index] = uint8(r >> 8)
            greenColor[index] = uint8(g >> 8)
            blueColor[index] = uint8(b >> 8)
            alphaChannel[index] = uint8(a >> 8)
            index++
        }
    }

    c = ArrayColor{
        Red:   redColor,
        Green: greenColor,
        Blue:  blueColor,
        Alpha: alphaChannel,
    }

    return
}

func (img ImageFile) WriteImage(bounds *image.Rectangle, c ArrayColor) {
    newImg := image.NewNRGBA(image.Rect(0, 0, bounds.Max.X, bounds.Max.Y))
    maxHeight := bounds.Max.Y
    maxWidth := bounds.Max.X

    for x := bounds.Min.X; x < maxWidth; x++ {
        for y := bounds.Min.Y; y < maxHeight; y++ {
            index := ((maxHeight - 1) * x) + (x + y)
            newImg.Set(x, y, color.RGBA{
                R: c.Red[index],
                G: c.Green[index],
                B: c.Blue[index],
                A: c.Alpha[index],
            })
        }
    }

    f, err := os.Create(img.Path)
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

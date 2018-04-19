package rsa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"golang.org/x/image/bmp"
)

// Key that used for rsa encryption
type Key struct {
	Private *rsa.PrivateKey
	Public  *rsa.PublicKey
}

// LoadKey from a file
func LoadKey(privateKeyPath string) (*Key, error) {
	data, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)

	decode, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return &Key{
		Private: decode,
		Public:  &decode.PublicKey,
	}, nil
}

// LoadPublicKey from file
func LoadPublicKey(publicKeyPath string) (*Key, error) {
	data, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)

	decode, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return &Key{
		Private: nil,
		Public:  decode,
	}, nil
}

// GenerateRSAKey based on 16 bit
func GenerateRSAKey() (*big.Int, *big.Int, *big.Int) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 16)
	if err != nil {
		log.Fatal(err)
	}

	e := big.NewInt(int64(privateKey.E))

	return privateKey.D, e, privateKey.N
}

// GenerateConcatColor based on limit
func GenerateConcatColor(color []uint8, lim int) [][]byte {
	var chunk []byte

	var buffer bytes.Buffer
	for _, el := range color {
		buffer.WriteByte(el)
	}

	buf := buffer.Bytes()

	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	// Append the rest buffer
	if len(buf) > 0 {
		chunks = append(chunks, buf)
	}

	return chunks
}

func fillPad(src, dest []byte) {
	remainingPad := len(dest) - len(src)
	for i := 0; i < remainingPad; i++ {
		dest[i] = 0
	}
	copy(dest[remainingPad:], src)
}

func fillPadForDecrypt(src, dest []byte) {
	remainingPad := len(dest) - len(src)
	if remainingPad < 0 {
		return
	}
	for i := 0; i < remainingPad; i++ {
		dest[i] = 0
	}
	copy(dest[remainingPad:], src)
}

func leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}

// LoadImage from a file
func (key *Key) LoadImage(path string) {
	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, err := bmp.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	bounds := m.Bounds()

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
}

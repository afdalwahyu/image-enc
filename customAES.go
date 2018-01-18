package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// arrayColor Color of image
type arrayColor struct {
	red, green, blue, alpha []byte
}

func AESEncryptImage(c arrayColor) arrayColor {
	key := []byte("example key 1234")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	colorSize := block.BlockSize()

	splittedRed := splitColor(c.red, colorSize)
	splittedGreen := splitColor(c.green, colorSize)
	splittedBlue := splitColor(c.blue, colorSize)
	// splittedAlpha := splitColor(c.alpha, colorSize)

	encryptedRed := encryptColor(splittedRed, colorSize, block)
	encryptedGreen := encryptColor(splittedGreen, colorSize, block)
	encryptedBlue := encryptColor(splittedBlue, colorSize, block)
	// encryptedAlpha := encryptColor(splittedAlpha, colorSize, block)

	encryptedColor := arrayColor{
		red:   encryptedRed,
		green: encryptedGreen,
		blue:  encryptedBlue,
		alpha: c.alpha,
	}

	return encryptedColor
}

func encryptColor(splittedColor [][]byte, colorSize int, block cipher.Block) []byte {
	var buffer bytes.Buffer
	for _, a := range splittedColor {
		chiper := make([]byte, colorSize, colorSize)
		block.Encrypt(chiper, a)
		buffer.Write(chiper)
	}

	return buffer.Bytes()
}

func AESDecryptImage(c arrayColor) arrayColor {
	key := []byte("example key 1234")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	colorSize := block.BlockSize()

	splittedRed := splitColor(c.red, colorSize)
	splittedGreen := splitColor(c.green, colorSize)
	splittedBlue := splitColor(c.blue, colorSize)
	// splittedAlpha := splitColor(c.alpha, colorSize)

	decryptedRed := decryptColor(splittedRed, colorSize, block)
	decryptedGreen := decryptColor(splittedGreen, colorSize, block)
	decryptedBlue := decryptColor(splittedBlue, colorSize, block)
	// decryptedAlpha := decryptColor(splittedAlpha, colorSize, block)

	encryptedColor := arrayColor{
		red:   decryptedRed,
		green: decryptedGreen,
		blue:  decryptedBlue,
		alpha: c.alpha,
	}

	return encryptedColor
}

func decryptColor(splittedColor [][]byte, colorSize int, block cipher.Block) []byte {
	var buffer bytes.Buffer
	for _, c := range splittedColor {
		message := make([]byte, colorSize, colorSize)
		block.Decrypt(message, c)
		buffer.Write(message)
	}

	return buffer.Bytes()
}

func splitColor(color []byte, lim int) [][]byte {
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
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}

	return chunks
}

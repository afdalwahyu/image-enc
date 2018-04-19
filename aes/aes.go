package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"io"
	"log"
	"skripsi/util"
	"time"
)

type Key struct {
	Password string
	Block    cipher.Block
}

// NewKey generate key using md5 as 128 bit key
func NewKey(password string) (*Key, error) {
	h := md5.New()
	io.WriteString(h, password)

	block, err := aes.NewCipher(h.Sum(nil))
	if err != nil {
		return nil, err
	}
	return &Key{
		Password: password,
		Block:    block,
	}, nil
}

// AESEncryptImage encrypt using aes
func (key *Key) AESEncryptImage(c util.ArrayColor) util.ArrayColor {
	start := time.Now()
	colorSize := key.Block.BlockSize()

	splitRed := splitColor(c.Red, colorSize)
	splitGreen := splitColor(c.Green, colorSize)
	splitBlue := splitColor(c.Blue, colorSize)

	encryptedRed := encryptColor(splitRed, colorSize, key.Block)
	encryptedGreen := encryptColor(splitGreen, colorSize, key.Block)
	encryptedBlue := encryptColor(splitBlue, colorSize, key.Block)

	elapsed := time.Since(start)
	log.Printf("AES Encryption took %s", elapsed)

	encryptedColor := util.ArrayColor{
		Red:   encryptedRed,
		Green: encryptedGreen,
		Blue:  encryptedBlue,
		Alpha: c.Alpha,
	}

	return encryptedColor
}

func encryptColor(splitColor [][]byte, colorSize int, block cipher.Block) []byte {
	var buffer bytes.Buffer
	for _, a := range splitColor {
		cipher := make([]byte, colorSize, colorSize)
		block.Encrypt(cipher, a)
		buffer.Write(cipher)
	}

	return buffer.Bytes()
}

// AESDecryptImage decrypt image using aes algorithm
func (key *Key) AESDecryptImage(c util.ArrayColor) util.ArrayColor {
	start := time.Now()
	colorSize := key.Block.BlockSize()

	splitRed := splitColor(c.Red, colorSize)
	splitGreen := splitColor(c.Green, colorSize)
	splitBlue := splitColor(c.Blue, colorSize)

	decryptedRed := decryptColor(splitRed, colorSize, key.Block)
	decryptedGreen := decryptColor(splitGreen, colorSize, key.Block)
	decryptedBlue := decryptColor(splitBlue, colorSize, key.Block)

	elapsed := time.Since(start)
	log.Printf("AES Decryption took %s", elapsed)

	encryptedColor := util.ArrayColor{
		Red:   decryptedRed,
		Green: decryptedGreen,
		Blue:  decryptedBlue,
		Alpha: c.Alpha,
	}

	return encryptedColor
}

func decryptColor(splitColor [][]byte, colorSize int, block cipher.Block) []byte {
	var buffer bytes.Buffer
	for _, c := range splitColor {
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
		chunks = append(chunks, buf)
	}

	return chunks
}

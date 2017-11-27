package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"time"
)

func loadKeyFromFile(privateKeyPath string) *rsa.PrivateKey {
	data, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	block, _ := pem.Decode(data)

	decode, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	return decode
}

// TODO: must spread array evenly
func RSAEncryptImage(red []uint8, green []uint8, blue []uint8, alpha []uint8) ([]uint8, []uint8, []uint8, []uint8) {
	start := time.Now()

	privateKey := loadKeyFromFile("./key_file/key")
	publicKey := privateKey.PublicKey

	k := (publicKey.N.BitLen()) / 8
	maxSize := k
	fmt.Println(maxSize)

	redConcat := generateConcatColor(red, maxSize)
	greenConcat := generateConcatColor(green, maxSize)
	blueConcat := generateConcatColor(blue, maxSize)
	alphaConcat := generateConcatColor(alpha, maxSize)

	EncRed := encryptConcattedColor(&publicKey, redConcat)
	EncGreen := encryptConcattedColor(&publicKey, greenConcat)
	EncBlue := encryptConcattedColor(&publicKey, blueConcat)
	EncAlpha := encryptConcattedColor(&publicKey, alphaConcat)

	fmt.Println(len(red), len(EncRed), len(redConcat))
	fmt.Println(len(green), len(EncGreen), len(greenConcat))
	fmt.Println(len(blue), len(EncBlue), len(blueConcat))
	fmt.Println(len(alpha), len(EncAlpha), len(alphaConcat))

	elapsed := time.Since(start)
	log.Printf("RSA Encryption took %s", elapsed)

	return EncRed, EncGreen, EncBlue, EncAlpha
}

func RSADecryptImage(red []uint8, green []uint8, blue []uint8, alpha []uint8) ([]uint8, []uint8, []uint8, []uint8) {
	start := time.Now()

	privateKey := loadKeyFromFile("./key_file/key")
	publicKey := privateKey.PublicKey

	k := (publicKey.N.BitLen()) / 8
	maxSize := k

	redConcat := generateConcatColor(red, maxSize)
	greenConcat := generateConcatColor(green, maxSize)
	blueConcat := generateConcatColor(blue, maxSize)
	alphaConcat := generateConcatColor(alpha, maxSize)

	DecRed := decryptConcattedColor(privateKey, redConcat)
	DecGreen := decryptConcattedColor(privateKey, greenConcat)
	DecBlue := decryptConcattedColor(privateKey, blueConcat)
	DecAlpha := decryptConcattedColor(privateKey, alphaConcat)

	fmt.Println(len(red), len(DecRed), len(redConcat))
	fmt.Println(len(green), len(DecGreen), len(greenConcat))
	fmt.Println(len(blue), len(DecBlue), len(blueConcat))
	fmt.Println(len(alpha), len(DecAlpha), len(alphaConcat))

	elapsed := time.Since(start)
	log.Printf("RSA Decryption took %s", elapsed)

	return DecRed, DecGreen, DecBlue, DecAlpha
}

func encryptConcattedColor(publicKey *rsa.PublicKey, concattedColor [][]byte) []byte {
	var buffer bytes.Buffer
	for _, el := range concattedColor {
		// chiperText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, el)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		chiperText := encrypt(new(big.Int), publicKey, new(big.Int).SetBytes(el))
		buffer.Write(chiperText)
	}
	return buffer.Bytes()
}

func testEncrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) []byte {
	return c.Add(m, big.NewInt(1)).Bytes()
}

func encrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) []byte {
	e := big.NewInt(int64(pub.E))
	c.Exp(m, e, pub.N)
	return c.Bytes()
}

func decryptConcattedColor(privateKey *rsa.PrivateKey, concattedColor [][]byte) []byte {
	var buffer bytes.Buffer
	for _, el := range concattedColor {
		// message, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, el)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		message := decrypt(privateKey, new(big.Int).SetBytes(el))
		buffer.Write(message)
	}

	return buffer.Bytes()
}

func testDecrypt(c *big.Int, privateKey *rsa.PrivateKey, m *big.Int) []byte {
	return c.Add(m, big.NewInt(-1)).Bytes()
}

func decrypt(privateKey *rsa.PrivateKey, c *big.Int) []byte {
	m := new(big.Int).Exp(c, privateKey.D, privateKey.N)

	return m.Bytes()
}

func genRSAKey() (*big.Int, *big.Int, *big.Int) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 16)
	if err != nil {
		log.Fatal(err)
	}

	e := big.NewInt(int64(privateKey.E))

	return privateKey.D, e, privateKey.N
}

func generateConcatColor(color []uint8, lim int) [][]byte {
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

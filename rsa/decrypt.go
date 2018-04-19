package rsa

import (
	"bytes"
	"crypto/rsa"
	"fmt"
	"log"
	"math/big"
	"skripsi/util"
	"time"
)

func (key *Key) DecryptImage(c *util.ArrayColor) util.ArrayColor {
	start := time.Now()

	k := (key.Public.N.BitLen() + 7) / 8
	maxSize := k

	redConcat := GenerateConcatColor(c.Red, maxSize)
	greenConcat := GenerateConcatColor(c.Green, maxSize)
	blueConcat := GenerateConcatColor(c.Blue, maxSize)

	fmt.Println("red")
	DecRed := decryptConcatColor(key.Private, redConcat)
	fmt.Println("green")
	DecGreen := decryptConcatColor(key.Private, greenConcat)
	fmt.Println("blue")
	DecBlue := decryptConcatColor(key.Private, blueConcat)

	// fmt.Println(len(c.Red), len(DecRed), len(redConcat))
	// fmt.Println(len(c.Green), len(DecGreen), len(greenConcat))
	// fmt.Println(len(c.Blue), len(DecBlue), len(blueConcat))

	elapsed := time.Since(start)
	log.Printf("RSA Decryption took %s", elapsed)

	decryptedColor := util.ArrayColor{
		Red:   DecRed,
		Green: DecGreen,
		Blue:  DecBlue,
		Alpha: c.Alpha,
	}

	return decryptedColor
}

func decryptConcatColor(privateKey *rsa.PrivateKey, concatColor [][]byte) []byte {
	var buffer bytes.Buffer
	for _, el := range concatColor {
		// message, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, el)
		// if err != nil {
		// 	buffer.Write(el)
		// 	continue
		// }
		// buffer.Write(message)

		k := (privateKey.N.BitLen()+7)/8 - 1
		dst := make([]byte, k)
		src := Decrypt(privateKey, new(big.Int).SetBytes(el))
		fillPadForDecrypt(src, dst)
		buffer.Write(dst)
	}

	return buffer.Bytes()
}

func testDecrypt(privateKey *rsa.PrivateKey, c *big.Int) []byte {
	return new(big.Int).Add(c, big.NewInt(-10)).Bytes()
}

func Decrypt(privateKey *rsa.PrivateKey, c *big.Int) []byte {
	return new(big.Int).Exp(c, privateKey.D, privateKey.N).Bytes()
}

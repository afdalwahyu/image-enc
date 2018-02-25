package rsa

import (
	"bytes"
	"crypto/rsa"
	"fmt"
	"log"
	"math/big"
	"time"
    "1-learn/util"
)

func (key *Key) EncryptImage(c *util.ArrayColor) util.ArrayColor{
	start := time.Now()

	k := (key.Public.N.BitLen()) / 8
	maxSize := k - 1
	fmt.Println(maxSize)

	redConcat := generateConcatColor(c.Red, maxSize)
	greenConcat := generateConcatColor(c.Green, maxSize)
	blueConcat := generateConcatColor(c.Blue, maxSize)

	EncRed := encryptConcatColor(key.Public, redConcat)
	EncGreen := encryptConcatColor(key.Public, greenConcat)
	EncBlue := encryptConcatColor(key.Public, blueConcat)

	fmt.Println(len(c.Red), len(EncRed), len(redConcat))
	fmt.Println(len(c.Green), len(EncGreen), len(greenConcat))
	fmt.Println(len(c.Blue), len(EncBlue), len(blueConcat))

	elapsed := time.Since(start)
	log.Printf("RSA Encryption took %s", elapsed)

	encryptedColor := util.ArrayColor{
		Red:   EncRed,
		Green: EncGreen,
		Blue:  EncBlue,
		Alpha: c.Alpha,
	}

	return encryptedColor
}

func encryptConcatColor(publicKey *rsa.PublicKey, concattedColor [][]byte) []byte {
	var buffer bytes.Buffer
	for _, el := range concattedColor {
		// cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, el)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		cipherText := encrypt(new(big.Int), publicKey, new(big.Int).SetBytes(el))
		buffer.Write(cipherText)
	}
	return buffer.Bytes()
}

func testEncrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) []byte {
	return c.Add(m, big.NewInt(10)).Bytes()
}

func encrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) []byte {
	if m.Cmp(pub.N) == 1 {
		fmt.Println("m larger than N")
	}
	e := big.NewInt(int64(pub.E))
	return c.Exp(m, e, pub.N).Bytes()
}
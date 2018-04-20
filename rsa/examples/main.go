package main

import (
	"fmt"
	"log"
	"skripsi/rsa"
	"skripsi/util"
	"time"
)

func main() {
	privKey, _ := rsa.LoadKey("./3072_priv")
	pubKey, err := rsa.LoadPublicKey("./3072_pub")
	if err != nil {
		log.Fatalln(err)
	}

	eyeOriginal := &util.ImageFile{Path: "eye.bmp"}
	eyeEncrypted := &util.ImageFile{Path: "eye-encrypted.bmp"}
	eyeDecrypted := &util.ImageFile{Path: "eye-decrypted.bmp"}

	bounds, inputPlain := eyeOriginal.OpenImage()

	start := time.Now()
	encryptedColor := pubKey.EncryptImage(&inputPlain)
	elapsed := time.Since(start)
	log.Printf("Encryption took %s", elapsed)

	eyeEncrypted.WriteImage(&bounds, encryptedColor)

	fmt.Println("decrypting....")

	bounds2, inputEncrypted := eyeEncrypted.OpenImage()
	decryptedColor := privKey.DecryptImage(&inputEncrypted)
	eyeDecrypted.WriteImage(&bounds2, decryptedColor)
}

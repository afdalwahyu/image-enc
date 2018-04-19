package main

import (
	"fmt"
	"log"
	"skripsi/rsa"
	"skripsi/util"
	"time"
)

func main() {
	key, _ := rsa.LoadKey("./key_file/1024key")

	eyeOriginal := &util.ImageFile{Path: "eye.bmp"}
	eyeEncrypted := &util.ImageFile{Path: "eye-encrypted.bmp"}
	eyeDecrypted := &util.ImageFile{Path: "eye-decrypted.bmp"}

	bounds, inputPlain := eyeOriginal.OpenImage()

	start := time.Now()
	encryptedColor := key.EncryptImage(&inputPlain)
	elapsed := time.Since(start)
	log.Printf("Encryption took %s", elapsed)

	eyeEncrypted.WriteImage(&bounds, encryptedColor)

	fmt.Println("decrypting....")

	bounds2, inputEncrypted := eyeEncrypted.OpenImage()
	decryptedColor := key.DecryptImage(&inputEncrypted)
	eyeDecrypted.WriteImage(&bounds2, decryptedColor)
}

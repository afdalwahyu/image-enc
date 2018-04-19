package main

import (
	"skripsi/aes"
	"skripsi/util"
)

func main() {
	eyeOriginal := &util.ImageFile{Path: "sipi1.bmp"}
	eyeEncrypted := &util.ImageFile{Path: "sipi1-encrypted.bmp"}
	eyeDecrypted := &util.ImageFile{Path: "sipi1-decrypted.bmp"}

	key, _ := aes.NewKey("afdal")

	bounds, inputPlainColor := eyeOriginal.OpenImage()
	encryptedColor := key.AESEncryptImage(inputPlainColor)
	eyeEncrypted.WriteImage(&bounds, encryptedColor)

	bounds2, inputEncryptedColor := eyeEncrypted.OpenImage()
	inputDecryptedColor := key.AESDecryptImage(inputEncryptedColor)
	eyeDecrypted.WriteImage(&bounds2, inputDecryptedColor)
}

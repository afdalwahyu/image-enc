package main

import (
	"skripsi/rsa"
	"skripsi/util"
)

func main() {
	// Kunci 4096, hasil encrypt -> 512 bit, kunci -> 511 <- chunk untuk decrypt
	// Kunci 2048, hasil encrypt -> 256 bit, kunci -> 255 <- chunk untuk decrypt
	// Kunci 2048key, hasil encrypt -> 128 bit, kunci -> 127 <- chunk untuk decrypt
	// Kunci 512,  hasil encrypt -> 64 bit,  kunci -> 63  <- chunk untuk decrypt
	// Kunci 256,  hasil encrypt -> 32 bit,  kunci -> 30  <- chunk untuk decrypt
	key, _ := rsa.LoadKey("./key_file/2048key")
	// Ternyata permasalahannya adalah bagaimana cara split arraynya
	// harus di cek lagi splitnya berapa aja, next mungkin idenya adalah
	// liat berapa panjang result encrypt (count) jika dibandingkan decrypt
	// Decrypt harus mengikuti panjang count dari encrypted message

	eyeOriginal := &util.ImageFile{Path: "eye.bmp"}
	eyeEncrypted := &util.ImageFile{Path: "eye-encrypted.bmp"}
	eyeDecrypted := &util.ImageFile{Path: "eye-decrypted.bmp"}

	bounds, inputPlain := eyeOriginal.OpenImage()
	encryptedColor := key.EncryptImage(&inputPlain)
	eyeEncrypted.WriteImage(&bounds, encryptedColor)

	bounds2, inputEncrypted := eyeEncrypted.OpenImage()
	decryptedColor := key.DecryptImage(&inputEncrypted)
	eyeDecrypted.WriteImage(&bounds2, decryptedColor)
}

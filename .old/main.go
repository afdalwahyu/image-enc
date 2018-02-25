package examples

import (
	"image"
	"image/color"

	"golang.org/x/image/bmp"
	// "image/jpeg"
	// "image/png"
	"log"
	"os"
)

func main() {

	bounds, inputPlainColor := openImage("./image/eye.bmp")
	encryptedColor := rsaEncryptImage(&inputPlainColor)
	generateImage("./image/eye-rsa-encrypted.bmp", &bounds, encryptedColor)

	bounds, inputPlainColor2 := openImage("./image/eye-rsa-encrypted.bmp")
	decryptedColor := rsaDecryptImage(&inputPlainColor2)
	generateImage("./image/eye-rsa-decrypted.bmp", &bounds, decryptedColor)

	// ----- HALF DONE AES ------

	// bounds, inputPlainColor := openImage("./image/eye.bmp")
	// encryptedColor := AESEncryptImage(inputPlainColor)
	// generateImage("./image/EncryptedAES.bmp", &bounds, encryptedColor)

	// bounds, inputColor2 := openImage("./image/EncryptedAES.bmp")
	// decryptedColor := AESDecryptImage(inputColor2)
	// generateImage("./image/DecryptedAES.bmp", &bounds, decryptedColor)

	// ----- DONE CHAOS --------

	// bounds, inputPlainColor := openImage("./image/4x4pixel.bmp")
	// encryptedColor := chaosEncrypt(&bounds, &inputPlainColor)
	// generateImage("./image/4x4pixel-enc.bmp", &bounds, encryptedColor)

	// bounds2, inputPlainColor2 := openImage("./image/4x4pixel-enc.bmp")
	// decryptedColor := chaosDecrypt(&bounds2, &inputPlainColor2)
	// generateImage("./image/4x4pixel-dec.bmp", &bounds2, decryptedColor)

	// bounds, inputPlainColor := openImage("./image/eye.bmp")
	// encryptedColor := chaosEncrypt(&bounds, &inputPlainColor)
	// generateImage("./image/eye-enc.bmp", &bounds, encryptedColor)

	// bounds2, inputPlainColor2 := openImage("./image/eye-enc.bmp")
	// decryptedColor := chaosDecrypt(&bounds2, &inputPlainColor2)
	// generateImage("./image/eye-dec.bmp", &bounds2, decryptedColor)
}

package main

import (
    "1-learn/util"
    "1-learn/aes"
)

func main() {
    eyeOriginal := &util.ImageFile{Path: "eye.bmp"}
    eyeEncrypted := &util.ImageFile{Path: "eye-encrypted.bmp"}
    eyeDecrypted := &util.ImageFile{Path: "eye-decrypted.bmp"}

    key, _ := aes.NewKey("example 2048key 1234")

    bounds, inputPlainColor := eyeOriginal.OpenImage()
    encryptedColor := key.AESEncryptImage(inputPlainColor)
    eyeEncrypted.WriteImage(&bounds, encryptedColor)

    bounds2, inputEncryptedColor := eyeEncrypted.OpenImage()
    inputDecryptedColor := key.AESDecryptImage(inputEncryptedColor)
    eyeDecrypted.WriteImage(&bounds2, inputDecryptedColor)
}

package main

import (
    "1-learn/util"
    "1-learn/rsa"
)

func main() {
    key, _ := rsa.LoadKey("./key_file/1024key")

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

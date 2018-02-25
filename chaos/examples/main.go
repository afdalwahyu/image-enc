package main

import (
    "1-learn/util"
    "1-learn/chaos"
)

func main() {
    eyeOriginal := &util.ImageFile{Path: "eye.bmp"}
    eyeEncrypted := &util.ImageFile{Path: "eye-encrypted.bmp"}
    eyeDecrypted := &util.ImageFile{Path: "eye-decrypted.bmp"}

    bounds, inputPlainColor := eyeOriginal.OpenImage()
    key, _ := chaos.NewChaosKey(&bounds, 1000, 0.5, 3.6, 0.0, 5)

    encryptedColor := key.ChaosEncrypt(&bounds, &inputPlainColor)
    eyeEncrypted.WriteImage(&bounds, encryptedColor)

    bounds2, inputEncryptedPlainColor := eyeEncrypted.OpenImage()

    decryptedColor := key.ChaosDecrypt(&bounds2, &inputEncryptedPlainColor)
    eyeDecrypted.WriteImage(&bounds2, decryptedColor)
}

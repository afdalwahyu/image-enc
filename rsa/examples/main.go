package main

import (
    "1-learn/rsa"
    "fmt"
    "math/big"
    "bytes"
)

func main() {
    // Ternyata permasalahannya adalah bagaimana cara split arraynya
    // harus di cek lagi splitnya berapa aja, next mungkin idenya adalah
    // liat berapa panjang result encrypt (count) jika dibandingkan decrypt
    // Decrypt harus mengikuti panjang count dari encrypted message
    key, _ := rsa.LoadKey("./key_file/256key")
    listColor := []uint8{
        0, 8, 5, 2, 5, 0, 0, 1, 3, 2, 4, 6,
        0, 8, 1, 2, 5, 5, 2, 6, 6, 3, 6, 1,
        0, 8, 5, 2, 5, 0, 3, 6, 2, 1, 7, 8,
        0, 8, 5, 2, 4, 6, 1, 7, 0, 8, 0, 7,
        0, 8, 5, 2, 5, 0, 0, 1, 3, 2, 4, 6,
        0, 8, 1, 2, 5, 5, 2, 6, 6, 3, 6, 1,
        0, 8, 5, 2, 5, 0, 3, 6, 2, 1, 7, 8,
        0, 8, 5, 2, 4, 6, 1, 7, 0, 8, 0, 7,
    }
    limit := key.Private.N.BitLen()/8-1
    colors := rsa.GenerateConcatColor(listColor, limit)

    fmt.Println("Key len", limit)
    //fmt.Println("\n\nPlain", colors, len(listColor))
    var buffer bytes.Buffer
    for _, el := range colors {
        cipher := rsa.Encrypt(new(big.Int), key.Public, new(big.Int).SetBytes(el))
        fmt.Println("Crypted......", cipher, len(cipher))
        buffer.Write(cipher)
    }
    encrypted := buffer.Bytes()
    //fmt.Println("\nPlain Encrypted", encrypted)

    encryptedColors := rsa.GenerateConcatColor(encrypted, limit+2)
    //fmt.Println("\n\nEncrypted", encryptedColors)
    var buffer2 bytes.Buffer
    for _, el := range encryptedColors {
        fmt.Println("Decrypting...", el, len(el))
        message := rsa.Decrypt(key.Private, new(big.Int).SetBytes(el))
        buffer2.Write(message)
    }
    decrypted := buffer2.Bytes()
    fmt.Println("\n\nDecrypted:", decrypted)

    //eyeOriginal := &util.ImageFile{Path: "eye.bmp"}
    //eyeEncrypted := &util.ImageFile{Path: "eye-encrypted.bmp"}
    //eyeDecrypted := &util.ImageFile{Path: "eye-decrypted.bmp"}

    //
    //bounds, inputPlain := eyeOriginal.OpenImage()
    //encryptedColor := key.EncryptImage(&inputPlain)
    //eyeEncrypted.WriteImage(&bounds, encryptedColor)
    //
    //bounds2, inputEncrypted := eyeEncrypted.OpenImage()
    //decryptedColor := key.DecryptImage(&inputEncrypted)
    //eyeDecrypted.WriteImage(&bounds2, decryptedColor)
}

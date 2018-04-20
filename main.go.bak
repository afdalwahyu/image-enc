package main

import (
	"fmt"
	"skripsi/aes"
	"skripsi/chaos"
	"skripsi/cmd"
	"skripsi/rsa"
	"skripsi/util"
)

func main() {
	mode, err := cmd.SelectMode()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mode)

	algo, err := cmd.SelectType()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(algo)

	inputImage, _ := cmd.SelectInputImage()
	outputImage, _ := cmd.OutputImageName()

	plainImage := &util.ImageFile{Path: inputImage}
	resultImage := &util.ImageFile{Path: outputImage}

	bounds, inputPlainColor := plainImage.OpenImage()

	if mode == "encrypt" {
		if algo == "AES" {
			password, _ := cmd.InputPassword()
			fmt.Println(password)
			key, _ := aes.NewKey(password)
			encryptedColor := key.AESEncryptImage(inputPlainColor)
			resultImage.WriteImage(&bounds, encryptedColor)
		} else if algo == "RSA" {
			pathKey, _ := cmd.InputPublicKey()

			key, _ := rsa.LoadPublicKey(pathKey)
			encryptedColor := key.EncryptImage(&inputPlainColor)
			resultImage.WriteImage(&bounds, encryptedColor)
		} else if algo == "Chaos" {
			cmd.SelectSequence()
			k, _ := cmd.InputK()
			u, _ := cmd.InputU()
			x0, _ := cmd.InputX0()

			key, _ := chaos.NewChaosKey(&bounds, 200, x0, u, k, 5)
			encryptedColor := key.ChaosEncrypt(&bounds, &inputPlainColor)
			resultImage.WriteImage(&bounds, encryptedColor)
		}
	} else if mode == "decrypt" {
		if algo == "AES" {
			password, _ := cmd.InputPassword()

			key, _ := aes.NewKey(password)
			encryptedColor := key.AESDecryptImage(inputPlainColor)
			resultImage.WriteImage(&bounds, encryptedColor)
		} else if algo == "RSA" {
			pathKey, _ := cmd.InputPrivateKey()

			key, _ := rsa.LoadKey(pathKey)
			encryptedColor := key.EncryptImage(&inputPlainColor)
			resultImage.WriteImage(&bounds, encryptedColor)
		} else if algo == "Chaos" {
			cmd.SelectSequence()
			k, _ := cmd.InputK()
			u, _ := cmd.InputU()
			x0, _ := cmd.InputX0()

			key, _ := chaos.NewChaosKey(&bounds, 200, x0, u, k, 5)
			encryptedColor := key.ChaosDecrypt(&bounds, &inputPlainColor)
			resultImage.WriteImage(&bounds, encryptedColor)
		}
	}
}

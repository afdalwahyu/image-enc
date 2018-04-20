package main

import (
	"log"
	"os"
	"skripsi/aes"
	"skripsi/chaos"
	"skripsi/rsa"
	"skripsi/util"
	"strconv"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	var inputImage, outputImage, typeAlg, sequence, aesPassword, pub, priv, k, u, x0 string

	app.Name = "Skripsi Encryption & Decryption"
	app.Usage = "Encrypt an image using AES, RSA, or Chaos"
	app.Version = "1.0.0"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:        "type, t",
			Value:       "AES",
			Usage:       "Encryption mode, use AES, RSA, or Chaos",
			Destination: &typeAlg,
		},
		cli.StringFlag{
			Name:        "in, i",
			Value:       "eye.bmp",
			Usage:       "input image",
			Destination: &inputImage,
		},
		cli.StringFlag{
			Name:        "out, o",
			Value:       "eye-end.bmp",
			Usage:       "output image",
			Destination: &outputImage,
		},
		cli.StringFlag{
			Name:        "aespwd, ap",
			Value:       "afdal",
			Usage:       "AES password",
			Destination: &aesPassword,
		},
		cli.StringFlag{
			Name:        "public, pub",
			Value:       "",
			Usage:       "Public key path file",
			Destination: &pub,
		},
		cli.StringFlag{
			Name:        "private, priv",
			Value:       "",
			Usage:       "Private key path file",
			Destination: &priv,
		},
		cli.StringFlag{
			Name:        "chaos_seq, cs",
			Value:       "LLM",
			Usage:       "chaotic sequence, only LLM, SSM, CCM",
			Destination: &sequence,
		},
		cli.StringFlag{
			Name:        "chaos_k, ck",
			Value:       "",
			Usage:       "chaos k key",
			Destination: &k,
		},
		cli.StringFlag{
			Name:        "chaos_u, cu",
			Value:       "",
			Usage:       "chaos u key",
			Destination: &u,
		},
		cli.StringFlag{
			Name:        "chaos_x0, cx0",
			Value:       "",
			Usage:       "chaos x0 key",
			Destination: &x0,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "encrypt",
			Aliases: []string{"e"},
			Usage:   "Encrypt mode",
			Flags:   flags,
			Action: func(c *cli.Context) error {
				plainImage := &util.ImageFile{Path: inputImage}
				resultImage := &util.ImageFile{Path: outputImage}
				bounds, inputPlainColor := plainImage.OpenImage()

				if typeAlg == "AES" {
					key, _ := aes.NewKey(aesPassword)
					encryptedColor := key.AESEncryptImage(inputPlainColor)
					resultImage.WriteImage(&bounds, encryptedColor)
				} else if typeAlg == "RSA" {
					key, _ := rsa.LoadPublicKey(pub)
					encryptedColor := key.EncryptImage(&inputPlainColor)
					resultImage.WriteImage(&bounds, encryptedColor)
				} else if typeAlg == "Chaos" {
					finalX0, _ := strconv.ParseFloat(x0, 64)
					finalU, _ := strconv.ParseFloat(u, 64)
					finalK, _ := strconv.ParseInt(u, 10, 64)

					key, _ := chaos.NewChaosKey(&bounds, 200, finalX0, finalU, finalK, 5, sequence)
					encryptedColor := key.ChaosEncrypt(&bounds, &inputPlainColor)
					resultImage.WriteImage(&bounds, encryptedColor)
				}

				return nil
			},
		},
		{
			Name:    "decrypt",
			Aliases: []string{"d"},
			Usage:   "Decrypt mode",
			Flags:   flags,
			Action: func(c *cli.Context) error {
				plainImage := &util.ImageFile{Path: inputImage}
				resultImage := &util.ImageFile{Path: outputImage}
				bounds, inputPlainColor := plainImage.OpenImage()

				if typeAlg == "AES" {
					key, _ := aes.NewKey(aesPassword)
					encryptedColor := key.AESDecryptImage(inputPlainColor)
					resultImage.WriteImage(&bounds, encryptedColor)
				} else if typeAlg == "RSA" {
					key, _ := rsa.LoadKey(priv)
					encryptedColor := key.EncryptImage(&inputPlainColor)
					resultImage.WriteImage(&bounds, encryptedColor)
				} else if typeAlg == "Chaos" {
					finalX0, _ := strconv.ParseFloat(x0, 64)
					finalU, _ := strconv.ParseFloat(u, 64)
					finalK, _ := strconv.ParseInt(u, 10, 64)

					key, _ := chaos.NewChaosKey(&bounds, 200, finalX0, finalU, finalK, 5, sequence)
					encryptedColor := key.ChaosDecrypt(&bounds, &inputPlainColor)
					resultImage.WriteImage(&bounds, encryptedColor)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

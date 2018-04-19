package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	var inputImage, outputImage, typeEncrypt, sequence string

	app.Name = "Skripsi Encryption & Decryption"
	app.Usage = "Encrypt an image using AES, RSA, or Chaos"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "type, t",
			Value:       "AES",
			Usage:       "Encryption mode, use AES, RSA, or Chaos",
			Destination: &typeEncrypt,
		},
		cli.StringFlag{
			Name:        "in, i",
			Value:       ".",
			Usage:       "input image",
			Destination: &inputImage,
		},
		cli.StringFlag{
			Name:        "out, o",
			Value:       ".",
			Usage:       "output image",
			Destination: &outputImage,
		},
		cli.StringFlag{
			Name:        "sequence, s",
			Value:       "LLM",
			Usage:       "chaotic sequence, only LLM, SSM, CCM",
			Destination: &sequence,
		},
		cli.StringFlag{
			Name:        "aespwd, ap",
			Value:       "afdal",
			Usage:       "AES password",
			Destination: &sequence,
		},
		cli.StringFlag{
			Name:        "public, pub",
			Value:       "",
			Usage:       "Public key path file",
			Destination: &sequence,
		},
		cli.StringFlag{
			Name:        "private, priv",
			Value:       "",
			Usage:       "Private key path file",
			Destination: &sequence,
		},
		cli.StringFlag{
			Name:        "k",
			Value:       "",
			Usage:       "chaos k key",
			Destination: &sequence,
		},
		cli.StringFlag{
			Name:        "u",
			Value:       "",
			Usage:       "chaos u key",
			Destination: &sequence,
		},
		cli.StringFlag{
			Name:        "x0",
			Value:       "",
			Usage:       "chaos x0 key",
			Destination: &sequence,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "encrypt",
			Aliases: []string{"e"},
			Usage:   "Encrypt mode",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "decrypt",
			Aliases: []string{"d"},
			Usage:   "Decrypt mode",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

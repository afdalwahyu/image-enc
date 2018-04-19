package cmd

import (
	"errors"

	"github.com/urfave/cli"
)

func encrypt(c *cli.Context, typeEnc, input, output, sequence string) error {
	if typeEnc == "AES" {
		return nil
	} else if typeEnc == "RSA" {
		return nil
	} else if typeEnc == "Chaos" {
		if sequence == "LLM" {
			return nil
		} else if sequence == "SSM" {
			return nil
		} else if sequence == "CCM" {
			return nil
		}
	} else {
		return errors.New("Only AES, RSA, and Chaos supported")
	}
	return nil
}

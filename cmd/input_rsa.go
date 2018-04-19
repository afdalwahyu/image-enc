package cmd

import "github.com/manifoldco/promptui"

func InputPublicKey() (string, error) {
	prompt := promptui.Prompt{
		Label: "public key path file",
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func InputPrivateKey() (string, error) {
	prompt := promptui.Prompt{
		Label: "private key path file",
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

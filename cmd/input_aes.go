package cmd

import "github.com/manifoldco/promptui"

func InputPassword() (string, error) {
	prompt := promptui.Prompt{
		Label: "AES password",
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

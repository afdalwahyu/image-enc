package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/manifoldco/promptui"
)

func SelectMode() (string, error) {
	prompt := promptui.Select{
		Label: "Select program mode",
		Items: []string{"Encrypt", "Decrypt"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func SelectType() (string, error) {
	prompt := promptui.Select{
		Label: "Select algorithm",
		Items: []string{"AES", "RSA", "Chaos"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println(result)

	return result, nil
}

func SelectInputImage() (string, error) {
	var listFileName []string
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		listFileName = append(listFileName, file.Name())
	}

	prompt := promptui.Select{
		Label: "Select image file",
		Items: listFileName,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func OutputImageName() (string, error) {
	prompt := promptui.Prompt{
		Label: "Output image name",
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

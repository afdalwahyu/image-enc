package cmd

import (
	"errors"
	"strconv"

	"github.com/manifoldco/promptui"
)

func InputX0() (float64, error) {
	var x0 float64
	validateFloat := func(input string) error {
		number, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid float number")
		}

		if number >= 1 || number <= -1 || number == 0 {
			return errors.New("X0 must zero float number")
		}

		x0 = number
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "X0",
		Validate: validateFloat,
	}

	_, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	return x0, nil
}

func SelectSequence() (string, error) {
	prompt := promptui.Select{
		Label: "Select chaotic sequence",
		Items: []string{"LLM", "SSM", "CCM"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func InputK() (int64, error) {
	var k int64

	validateK := func(input string) error {
		number, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return errors.New("Invalid float number")
		}

		if number < 8 || number > 20 {
			return errors.New("k must be higher than 8 and below 20")
		}

		k = number
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "K",
		Validate: validateK,
	}

	_, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	return k, nil
}

func InputU() (float64, error) {
	var u float64
	validateFloat := func(input string) error {
		number, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid float number")
		}

		if number < 0 || number > 10 {
			return errors.New("U must in between 0 to 10")
		}

		u = number
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "U",
		Validate: validateFloat,
	}

	_, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	return u, nil
}

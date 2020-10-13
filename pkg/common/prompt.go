package common

import (
	"errors"
	"github.com/manifoldco/promptui"
	"strconv"
	"strings"
)

func PromptString(s string) (string,error){
	prompt := promptui.Prompt{
		Label: s,
		Validate: ValidateEmptyInput,
	}
	return prompt.Run()
}
func ValidateEmptyInput(input string) error {
	if len(strings.TrimSpace(input)) < 1 {
		return errors.New("this input must not be empty")
	}
	return nil
}

func ValidateIntegerNumberInput(input string) error {
	_, err := strconv.ParseInt(input, 0, 64)
	if err != nil {
		return errors.New("invalid number")
	}
	return nil
}
func PromptInteger(name string) (int64, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: ValidateIntegerNumberInput,
	}

	promptResult, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	parseInt, _ := strconv.ParseInt(promptResult, 0, 64)
	return parseInt, nil
}
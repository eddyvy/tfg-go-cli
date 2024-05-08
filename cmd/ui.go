package cmd

import (
	"github.com/manifoldco/promptui"
)

func selectOption(label string, options []string) (string, error) {
	opts := append([]string{"All tables"}, options...)
	prompt := promptui.Select{
		Label: label,
		Items: opts,
	}

	_, result, err := prompt.Run()

	return result, err
}

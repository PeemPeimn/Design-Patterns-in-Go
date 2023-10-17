package factory_method

import "fmt"

func getDialog(dialogType string) (Dialog, error) {
	if dialogType == "Windows" {
		return &windowsDialog{}, nil
	}

	if dialogType == "Web" {
		return &webDialog{}, nil
	}

	return nil, fmt.Errorf("unknown dialog type")
}

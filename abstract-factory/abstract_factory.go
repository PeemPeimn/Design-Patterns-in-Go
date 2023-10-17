package abstract_factory

import "fmt"

func NewFactory(factoryType string) (GUIFactory, error) {
	if factoryType == "Windows" {
		return &windowsFactory{}, nil
	}

	if factoryType == "Mac" {
		return &macFactory{}, nil
	}

	return nil, fmt.Errorf("unsupported factory type")
}

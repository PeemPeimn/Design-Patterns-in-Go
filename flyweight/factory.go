package flyweight

import "fmt"

// constant facory
var treeFactory = &TreeFactory{make(map[string]*TreeType)}

// A flyweight factory
type TreeFactory struct {
	treeTypes map[string]*TreeType
}

func (f *TreeFactory) getTreeTypes(name, color, texture string) *TreeType {
	key := fmt.Sprintf("%s %s %s", name, color, texture)

	if treeType, exists := f.treeTypes[key]; exists {
		return treeType
	}

	f.treeTypes[key] = &TreeType{name, color, texture}

	return f.treeTypes[key]
}

func getFactoryInstance() *TreeFactory {
	return treeFactory
}

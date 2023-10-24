package flyweight

type Forest struct {
	trees []*Tree
}

func (f *Forest) PlantTree(x, y int, name, color, texture string) {
	factory := getFactoryInstance()
	treeTypes := factory.getTreeTypes(name, color, texture)
	f.trees = append(f.trees, &Tree{x, y, treeTypes})
}

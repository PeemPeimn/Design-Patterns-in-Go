package flyweight

type Tree struct {
	x        int
	y        int
	treeType *TreeType
}

type TreeType struct {
	name    string
	color   string
	texture string
}

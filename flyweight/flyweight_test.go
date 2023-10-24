package flyweight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlyweight(t *testing.T) {
	forest := &Forest{}

	forest.PlantTree(0, 0, "oak", "oak", "smooth")
	forest.PlantTree(1, 0, "oak", "oak", "smooth")

	// check number of flyweights
	assert.Equal(t, 1, len(treeFactory.treeTypes)) // flyweights
	assert.Equal(t, 2, len(forest.trees))

	forest.PlantTree(0, 1, "oak", "dark oak", "smooth")
	forest.PlantTree(1, 1, "oak", "dark oak", "rough")
	forest.PlantTree(1, 4, "oak", "oak", "smooth")

	// check number of flyweights
	assert.Equal(t, 3, len(treeFactory.treeTypes)) // flyweights
	assert.Equal(t, 5, len(forest.trees))
}

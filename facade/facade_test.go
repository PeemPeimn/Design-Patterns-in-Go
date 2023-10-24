package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFacade(t *testing.T) {
	facade := &Facade{}

	assert.Equal(t, []int{1, 3, 5}, facade.ExecuteOddSubSystems())
	assert.Equal(t, []int{2, 4}, facade.ExecuteEvenSubSystems())
}

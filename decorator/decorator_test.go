package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecorator(t *testing.T) {
	b := &BaseDecorator{&ConcreteComponent{}}
	assert.Equal(t, "execute_base.execute.", b.Execute())

	d1 := &ConcreteDecorator1{b}
	assert.Equal(t, "execute_dec1.execute_base.execute.", d1.Execute())

	d2 := &ConcreteDecorator2{d1}
	assert.Equal(t, "execute_dec2.execute_dec1.execute_base.execute.", d2.Execute())

}

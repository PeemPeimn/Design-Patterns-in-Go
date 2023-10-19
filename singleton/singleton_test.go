package singleton

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func pushInstance(c chan *Singleton, wg *sync.WaitGroup) {
	defer wg.Done()
	instance := getInstance()
	c <- instance
}

func TestSingleton(t *testing.T) {
	c := make(chan *Singleton, 100)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go pushInstance(c, &wg)
	}
	wg.Wait()
	expected := getInstance()
	close(c)

	for instance := range c {
		assert.Equal(t, instance, expected)
	}

}

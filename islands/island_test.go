package island

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	assert.Equal(t, 1, numIslands([][]byte{
		[]byte{49, 49, 49, 49, 48},
		[]byte{49, 49, 48, 49, 48},
		[]byte{49, 49, 48, 48, 48},
		[]byte{48, 48, 48, 48, 48}}))

}

func TestNumIslands2(t *testing.T) {
	assert.Equal(t, 3, numIslands([][]byte{
		[]byte{49, 49, 48, 48, 48},
		[]byte{49, 49, 48, 48, 48},
		[]byte{48, 48, 49, 48, 48},
		[]byte{48, 48, 48, 49, 49}}))
}

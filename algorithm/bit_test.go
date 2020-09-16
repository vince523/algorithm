package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfTwo(t *testing.T) {
	assert.Equal(t, true, IsPowerOfTwo(2))
	assert.Equal(t, false, IsPowerOfTwo(3))
	assert.Equal(t, true, IsPowerOfTwo(0))
	assert.Equal(t, false, IsPowerOfTwo(-1))
	assert.Equal(t, true, IsPowerOfTwo(8))
}

func TestHowManyOne(t *testing.T) {
	assert.Equal(t, 1, HowManyOne(2))
	assert.Equal(t, 2, HowManyOne(3))
}

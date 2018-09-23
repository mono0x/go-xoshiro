package splitmix64

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

var _ rand.Source64 = (*Source)(nil)

func TestNewSource(t *testing.T) {
	s := NewSource(1)
	assert.NotNil(t, s)
}

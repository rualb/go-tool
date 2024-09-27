package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNothing(t *testing.T) {

	a := "Hello"
	b := "Hello"

	assert.Equal(t, a, b, "The two words should be the same.")

}

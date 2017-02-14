package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUI1(e Engine, t *testing.T) {
	assert := assert.New(t)
	result := NewResult()
	e.GetPage("fred", result)
	assert.True(result.NotFound())
	return
}

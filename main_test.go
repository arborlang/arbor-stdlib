package main

import (
	"testing"

	arbor "github.com/arborlang/arbor-dev"
	"github.com/stretchr/testify/assert"
)

func TestEnvConfromsToModule(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*arbor.Module)(nil), Env)
}

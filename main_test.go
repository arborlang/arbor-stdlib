package main

import (
	"github.com/radding/arbor-dev"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnvConfromsToModule(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*arbor.Module)(nil), Env)
}

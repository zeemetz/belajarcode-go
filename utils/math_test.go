package utils_test

import (
	"testing"
	"testong/utils"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 2, utils.Add(1, 1))
}

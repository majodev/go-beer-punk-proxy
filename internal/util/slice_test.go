package util_test

import (
	"testing"

	"github.com/majodev/go-beer-punk-proxy/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestContainsString(t *testing.T) {
	test := []string{"a", "b", "d"}
	assert.True(t, util.ContainsString(test, "a"))
	assert.True(t, util.ContainsString(test, "b"))
	assert.False(t, util.ContainsString(test, "c"))
	assert.True(t, util.ContainsString(test, "d"))
}

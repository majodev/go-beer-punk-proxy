package util_test

import (
	"math"
	"testing"

	"github.com/majodev/go-beer-punk-proxy/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestMinAndMapInt(t *testing.T) {
	t.Parallel()

	max := math.MaxInt32
	min := math.MinInt32
	assert.Equal(t, max, util.MaxInt(max, min))
	assert.Equal(t, max, util.MaxInt(min, max))
	assert.Equal(t, min, util.MinInt(max, min))
	assert.Equal(t, min, util.MinInt(min, max))
	assert.Equal(t, min, util.MaxInt(min, min))
	assert.Equal(t, max, util.MinInt(max, max))
}

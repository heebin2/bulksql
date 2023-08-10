package placeholder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgument(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Sequence(3), []any{1, 2, 3})
}

package placeholder

import (
	"testing"

	"github.com/heebin2/bulksql/internal/dsn"
	"github.com/stretchr/testify/assert"
)

func TestSprintPlaceholder(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(NewPlaceholder(dsn.MySQL))
	assert.NotNil(NewPlaceholder(dsn.PostgreSQL))
	assert.NotNil(NewPlaceholder(dsn.Oracle))
	assert.Nil(NewPlaceholder(dsn.None))
}

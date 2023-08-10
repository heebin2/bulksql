package method

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMethod(t *testing.T) {
	assert := assert.New(t)

	method, err := FindMethod("INSERT INTO a VALUES ((?, ?), (?, ?))")
	assert.Equal(method, Insert)
	assert.Nil(err)
}

func TestRemoveMethodAfterText(t *testing.T) {
	assert := assert.New(t)

	sqlc := `-- name: CreateAuthor :execresult
	INSERT INTO authors (
	  name, bio
	) VALUES (
	  ?, ? 
	)
	`
	ret := `INSERT INTO authors (
	  name, bio
	) VALUES (
	  ?, ? 
	)
	`

	sqlc, err := RemoveMethodAfterText(sqlc)
	assert.Equal(sqlc, ret)
	assert.Nil(err)
}

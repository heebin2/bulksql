package dsn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDSN1(t *testing.T) {
	assert := assert.New(t)

	dsn, err := FindDSN("INSERT INTO a VALUES ((?, ?), (?, ?))")
	assert.Equal(dsn, MySQL)
	assert.Nil(err)
}

func TestDSN2(t *testing.T) {
	assert := assert.New(t)

	dsn, err := FindDSN("INSERT INTO a VALUES ($1, $2)")
	assert.Equal(dsn, PostgreSQL)
	assert.Nil(err)
}

func TestDSN3(t *testing.T) {
	assert := assert.New(t)

	dsn, err := FindDSN("INSERT INTO a VALUES (:1, :2)")
	assert.Equal(dsn, Oracle)
	assert.Nil(err)
}

func TestDSN4(t *testing.T) {
	assert := assert.New(t)

	dsn, err := FindDSN("INSERT INTO a VALUES ")
	assert.Equal(dsn, None)
	assert.NotNil(err)
}

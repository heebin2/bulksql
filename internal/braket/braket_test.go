package braket

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBraket(t *testing.T) {
	sql := "INSERT INTO a VALUES(?, ?)"
	assert := assert.New(t)
	begin, err := BeginBraket(sql)
	assert.Equal(begin, 20)
	assert.Nil(err)

	end, err := EndBraket(sql, begin)
	assert.Equal(end, 25)
	assert.Nil(err)
}
func TestBraket2(t *testing.T) {
	sql := "INSERT INTO a VALUES ((?, ?), (?, ?))"
	assert := assert.New(t)
	begin, err := BeginBraket(sql)
	assert.Equal(begin, 21)
	assert.Nil(err)

	end, err := EndBraket(sql, begin)
	assert.Equal(end, 36)
	assert.Nil(err)
}

func TestBraket3(t *testing.T) {
	sql := "INSERT INTO a values ((?, ?), (?, ?)"
	assert := assert.New(t)
	begin, err := BeginBraket(sql)
	assert.Equal(begin, 21)
	assert.Nil(err)

	end, err := EndBraket(sql, begin)
	assert.Equal(end, 0)
	assert.NotNil(err)
}

func TestBraket4(t *testing.T) {
	sql := "INSERT INTO a v ((?, ?), (?, ?)"
	assert := assert.New(t)
	begin, err := BeginBraket(sql)
	assert.Equal(begin, 0)
	assert.NotNil(err)

	end, err := EndBraket(sql, begin)
	assert.Equal(end, 0)
	assert.NotNil(err)
}

func TestBraket5(t *testing.T) {
	sql := "INSERT INTO a values )?, ?))"
	assert := assert.New(t)
	begin, err := BeginBraket(sql)
	assert.Equal(begin, 0)
	assert.NotNil(err)

	end, err := EndBraket(sql, begin)
	assert.Equal(end, 21)
	assert.Nil(err)
}

func TestBraket6(t *testing.T) {
	sql := "INSERT INTO a values "
	assert := assert.New(t)
	begin, err := BeginBraket(sql)
	assert.Equal(begin, 0)
	assert.NotNil(err)

	end, err := EndBraket(sql, begin)
	assert.Equal(end, 0)
	assert.NotNil(err)
}

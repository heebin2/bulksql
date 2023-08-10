package bulksql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMysql(t *testing.T) {
	assert := assert.New(t)

	bc, err := NewBatch("INSERT INTO table_name VALUES (?, ?, ?) ON~")
	assert.NotNil(bc)
	assert.Nil(err)
	assert.Equal(bc, bc.New())
	assert.Nil(bc.Push(1, 2, 3))
	assert.Nil(bc.Push(1, 2, 3))
	assert.Equal(bc.DSN(), "mysql")

	assert.Equal(bc.Len(), 6)
	assert.Equal(bc.DSN(), "mysql")

	query, err := bc.Query()
	assert.Nil(err)
	assert.Equal(query, "INSERT INTO table_name VALUES (?,?,?),(?,?,?) ON~")

	for {
		if err := bc.Push(1, 2, 3); err != nil {
			assert.Equal(err, ErrRangeOver)
			break
		}
	}
	assert.NotNil(bc.Datas())
	bc.Clear()

	bc2, err := NewBatch("INSERT INTO table_name VALUES (?) ON~")
	assert.Nil(err)
	assert.Nil(bc2.Push(1))
	assert.Nil(bc2.Push(2))
	assert.Nil(bc2.Push(3))
	assert.NotNil(bc2.Push(4, 5))
	query2, err := bc2.Query()
	assert.Nil(err)
	assert.Equal(query2, "INSERT INTO table_name VALUES (?,?,?) ON~")

	bc3 := bc.New()
	_, err = bc3.Query()
	assert.NotNil(err)

	_, err = NewBatch("123")
	assert.NotNil(err)

	bc5, _ := NewBatch("INSERT INTO table_name VALUES ?")
	bc5.Push(1)
	_, err = bc5.Query()
	assert.NotNil(err)

	bc6, _ := NewBatch("INSERT INTO table_name VALUES (?")
	bc6.Push(1)
	_, err = bc6.Query()
	assert.NotNil(err)

}

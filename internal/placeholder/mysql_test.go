package placeholder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMysql(t *testing.T) {
	assert := assert.New(t)

	m := Mysql{}

	sql := "(?,?),(?,?),(?,?)"

	assert.Equal(m.Generate(2, 3), sql)
	assert.Equal(m.QueryLen(2, 3), len(sql))
	assert.Equal(m.ArgsCount("(?, ?)"), 2)
	assert.Equal(m.MaxDatas("insert into table values (?, ?)"), (255-31-1)/2+1)
	assert.Equal(m.Generate(1, 4), "(?,?,?,?)")
}

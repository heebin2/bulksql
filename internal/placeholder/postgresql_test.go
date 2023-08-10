package placeholder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgresql(t *testing.T) {
	assert := assert.New(t)

	p := Postgresql{}

	sql := "($1,$2),($3,$4),($5,$6)"

	assert.Equal(p.Generate(2, 3), sql)
	assert.Equal(p.QueryLen(2, 3), len(sql))
	assert.Equal(p.ArgsCount("($1, $2)"), 2)
	assert.Equal(p.MaxDatas("insert into table values ($1, :$2)"), (255-33-1)/3+1)
	assert.Equal(p.Generate(1, 4), "($1,$2,$3,$4)")
}

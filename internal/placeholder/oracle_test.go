package placeholder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOracle(t *testing.T) {
	assert := assert.New(t)

	o := Oracle{}

	sql := "(:1,:2),(:3,:4),(:5,:6)"

	assert.Equal(o.Generate(2, 3), sql)
	assert.Equal(o.QueryLen(2, 3), len(sql))
	assert.Equal(o.ArgsCount("(:1, :2)"), 2)
	assert.Equal(o.MaxDatas("insert into table values (:1, :2)"), (255-33-1)/3+1)
	assert.Equal(o.Generate(1, 4), "(:1,:2,:3,:4)")
}

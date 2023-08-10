package bulksql

import (
	"fmt"

	"github.com/heebin2/bulksql/internal/braket"
	"github.com/heebin2/bulksql/internal/dsn"
	"github.com/heebin2/bulksql/internal/placeholder"
	"github.com/pkg/errors"
)

// Batch manages information for bulk operation.
type Batch struct {
	sql          string
	datas        []any
	argsCount    int
	dsn          dsn.DSN
	placeholder  placeholder.Placeholder
	maxArgsCount int
}

// NewBatch()
func NewBatch(sql string) (Batch, error) {
	dsnName, err := dsn.FindDSN(sql)
	if err != nil {
		return Batch{}, errors.Wrap(err, "find dsn fail")
	}

	ph := placeholder.NewPlaceholder(dsnName)

	return Batch{
		sql:          sql,
		datas:        make([]any, 0),
		placeholder:  ph,
		argsCount:    ph.ArgsCount(sql),
		dsn:          dsnName,
		maxArgsCount: ph.MaxDatas(sql),
	}, nil
}

// New() is a deep-copy excluding data.
func (b *Batch) New() Batch {
	return Batch{
		sql:          b.sql,
		datas:        make([]any, 0),
		argsCount:    b.argsCount,
		placeholder:  placeholder.NewPlaceholder(b.dsn),
		dsn:          b.dsn,
		maxArgsCount: b.maxArgsCount,
	}
}

// Push()
// The number of arguments you put in push() should be the same as the number of arguments in sql.
func (b *Batch) Push(param ...any) error {
	if len(param) != b.argsCount {
		return ErrUnmatchParam
	}

	if len(b.datas)+b.argsCount > b.maxArgsCount {
		return ErrRangeOver
	}

	b.datas = append(b.datas, param...)

	return nil
}

// Clear() is data clean
func (b *Batch) Clear() {
	b.datas = make([]any, 0)
}

// Query() generate SQL
func (b *Batch) Query() (string, error) {
	if len(b.datas) < b.argsCount {
		return "", fmt.Errorf("length of the data is too small")
	}

	begin, err := braket.BeginBraket(b.sql)
	if err != nil {
		return "", errors.Wrap(err, "begin braket fail")
	}

	end, err := braket.EndBraket(b.sql, begin)
	if err != nil {
		return "", errors.Wrap(err, "end braket fail")
	}

	return b.sql[:begin] +
		b.placeholder.Generate(b.argsCount, len(b.datas)/b.argsCount) +
		b.sql[end+1:], nil
}

// Len() len(data)
func (b *Batch) Len() int {
	return len(b.datas)
}

// DSN() postgresql, mysql, oracle
func (b *Batch) DSN() string {
	return string(b.dsn)
}

// Datas() data slice (do not edit)
func (b *Batch) Datas() []any {
	return b.datas
}

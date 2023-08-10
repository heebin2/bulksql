package placeholder

import "github.com/heebin2/bulksql/internal/dsn"

type Placeholder interface {
	ArgsCount(sql string) int
	Generate(atts, rows int) string
	QueryLen(atts, rows int) int
	MaxDatas(sql string) int
}

func NewPlaceholder(dsnName dsn.DSN) Placeholder {
	switch dsnName {
	case dsn.MySQL:
		return &Mysql{}
	case dsn.PostgreSQL:
		return &Postgresql{}
	case dsn.Oracle:
		return &Oracle{}
	default:
		return nil
	}
}

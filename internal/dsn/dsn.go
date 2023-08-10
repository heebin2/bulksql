package dsn

import (
	"fmt"
	"strings"
)

type DSN string

const (
	None       DSN = ""
	MySQL      DSN = "mysql"
	PostgreSQL DSN = "postgresql"
	Oracle     DSN = "oracle"
)

func FindDSN(sql string) (DSN, error) {
	if strings.Contains(sql, "?") {
		return MySQL, nil
	}

	if strings.Contains(sql, "$") {
		return PostgreSQL, nil
	}

	if strings.Contains(sql, ":") {
		return Oracle, nil
	}

	return "", fmt.Errorf("not found placeholder")
}

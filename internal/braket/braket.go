package braket

import (
	"fmt"
	"strings"
)

const BeginMinLength = 14 // "insert into a " 14

func findValuesIndex(sql string) (int, error) {
	valueIndex := strings.Index(sql, "values")
	if valueIndex >= BeginMinLength {
		return valueIndex, nil
	}

	valueIndex = strings.Index(sql, "VALUES")
	if valueIndex >= BeginMinLength {
		return valueIndex, nil
	}

	return 0, fmt.Errorf("not found values")
}

func BeginBraket(sql string) (int, error) {
	valuesIndex, err := findValuesIndex(sql)
	if err != nil {
		return 0, err
	}

	for i := valuesIndex + 6; i < len(sql); i++ {
		if sql[i] == '(' {
			return i, nil
		}

		if sql[i] == ')' {
			return 0, fmt.Errorf("not found \"(\"")
		}
	}

	return 0, fmt.Errorf("not found \"(\"")
}

func EndBraket(sql string, beginBraket int) (int, error) {
	open := 0
	for i := beginBraket + 1; i < len(sql); i++ {
		if sql[i] == ')' {
			if open != 0 {
				open -= 1
			} else {
				return i, nil
			}
		} else if sql[i] == '(' {
			open += 1
		}
	}

	return 0, fmt.Errorf("not found \")\"")
}

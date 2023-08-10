package placeholder

import "strings"

type Mysql struct {
}

// ArgsCount(sql string) int
// Generate(atts, rows int) string
// QueryLen(atts, rows int) int
var _ Placeholder = (*Mysql)(nil)

func (m *Mysql) Generate(argsLen, rowsLen int) string {
	element := "("
	if argsLen == 1 {
		for i := 0; i < rowsLen; i++ {
			if i != rowsLen-1 {
				element += "?,"
			} else {
				element += "?)"
			}
		}
		return element
	}

	for i := 0; i < argsLen; i++ {
		if i != argsLen-1 {
			element += "?,"
		} else {
			element += "?)"
		}
	}

	ret := element
	for i := 1; i < rowsLen; i++ {
		ret += "," + element
	}

	return ret
}

func (m *Mysql) QueryLen(argsLen, rowsLen int) int {
	return ((argsLen*2+1)+1)*rowsLen - 1
}

func (m *Mysql) ArgsCount(sql string) int {
	return strings.Count(sql, "?")
}

func (m *Mysql) MaxDatas(sql string) int {
	return (255-len(sql)-1)/2 + 1
}

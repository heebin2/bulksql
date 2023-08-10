package placeholder

import (
	"fmt"
	"strings"
)

type Postgresql struct {
}

var _ Placeholder = (*Postgresql)(nil)

func (p *Postgresql) Generate(argsLen, rowsLen int) string {
	element := "("
	if argsLen == 1 {
		for i := 0; i < rowsLen; i++ {
			if i != rowsLen-1 {
				element += "$%d,"
			} else {
				element += "$%d)"
			}
		}
		return fmt.Sprintf(element, Sequence(rowsLen)...)
	}

	for i := 0; i < argsLen; i++ {
		if i != argsLen-1 {
			element += "$%d,"
		} else {
			element += "$%d)"
		}
	}

	format := element
	for i := 1; i < rowsLen; i++ {
		format += "," + element
	}

	return fmt.Sprintf(format, Sequence(argsLen*rowsLen)...)
}

func (p *Postgresql) QueryLen(argsLen, rowsLen int) int {
	numlen := SequenceLen(argsLen * rowsLen)

	return numlen + ((argsLen*2+1)+1)*rowsLen - 1
}

func (p *Postgresql) ArgsCount(sql string) int {
	return strings.Count(sql, "$")
}

func (p *Postgresql) MaxDatas(sql string) int {
	return (255-len(sql)-1)/3 + 1
}

package placeholder

import "strconv"

var seq = []any{}
var seqLen = []int{}

func init() {
	str := ""
	for i := 1; i < 255; i++ {
		seq = append(seq, i)
		str += strconv.Itoa(i)
		seqLen = append(seqLen, len(str))
	}
}

func Sequence(len int) []any {
	return seq[:len]
}

func SequenceLen(len int) int {
	return seqLen[len-1]
}

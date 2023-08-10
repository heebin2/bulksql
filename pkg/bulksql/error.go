package bulksql

import "errors"

var ErrRangeOver = errors.New("bulksql: range over")
var ErrUnmatchParam = errors.New("bulksql: unmatched parameter")

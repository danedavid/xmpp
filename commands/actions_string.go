// Code generated by "stringer -type=Actions -linecomment"; DO NOT EDIT.

package commands

import "strconv"

const (
	_Actions_name_0 = "prevnext"
	_Actions_name_1 = "complete"
)

var (
	_Actions_index_0 = [...]uint8{0, 4, 8}
)

func (i Actions) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _Actions_name_0[_Actions_index_0[i]:_Actions_index_0[i+1]]
	case i == 4:
		return _Actions_name_1
	default:
		return "Actions(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
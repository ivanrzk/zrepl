// Code generated by "stringer -type=FrameType"; DO NOT EDIT.

package rpc

import "strconv"

const (
	_FrameType_name_0 = "FrameTypeHeaderFrameTypeDataFrameTypeTrailer"
	_FrameType_name_1 = "FrameTypeRST"
)

var (
	_FrameType_index_0 = [...]uint8{0, 15, 28, 44}
)

func (i FrameType) String() string {
	switch {
	case 1 <= i && i <= 3:
		i -= 1
		return _FrameType_name_0[_FrameType_index_0[i]:_FrameType_index_0[i+1]]
	case i == 255:
		return _FrameType_name_1
	default:
		return "FrameType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

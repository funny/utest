package unitest

import (
	"io"
	"testing"
)

func Test_All(t *testing.T) {
	Check(t, false, 1, 2, 3)
	CheckNotError(t, io.ErrClosedPipe)
	CheckByte(t, 1, "==", 2)
	CheckInt(t, 1, ">=", 2)
	CheckInt64(t, 3, ">", 4)
	CheckFloat32(t, 1.233, "==", 3.333)
	CheckFloat64(t, 1.233, "==", 3.333)
	CheckRune(t, '你', '好')
	CheckString(t, "sadkfjsl", "sdfs*\r\n")
	CheckBytes(t, []byte{1, 2, 3, 3}, []byte{3, 4, 5, 6})
}

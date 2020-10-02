package bytesconv

import (
	"strconv"
)

// Int8ToBytes returns the extended buffer of i in the given base,
func Int8ToBytes(i int8) []byte {
	dst := make([]byte, 0, 4)

	return strconv.AppendInt(dst, int64(i), 10)
}

func Int16ToBytes(i int32) []byte {
	dst := make([]byte, 0, 6)

	return strconv.AppendInt(dst, int64(i), 10)
}

func Itobytes(i int) []byte {
	if strconv.IntSize == 64 {
		return Int64ToBytes(int64(i))
	}
	return Int32ToBytes(int32(i))
}

func Int32ToBytes(i int32) []byte {
	dst := make([]byte, 0, 11)

	return strconv.AppendInt(dst, int64(i), 10)
}

func Int64ToBytes(i int64) []byte {
	dst := make([]byte, 0, 20)

	return strconv.AppendInt(dst, i, 10)
}

func Uint8ToBytes(i uint8) []byte {
	dst := make([]byte, 0, 4)

	return strconv.AppendUint(dst, uint64(i), 10)
}

func Uint16ToBytes(i uint16) []byte {
	dst := make([]byte, 0, 6)

	return strconv.AppendUint(dst, uint64(i), 10)
}

func Uint32ToBytes(i uint32) []byte {
	dst := make([]byte, 0, 11)

	return strconv.AppendUint(dst, uint64(i), 10)
}

func Uint64ToBytes(i uint64) []byte {
	dst := make([]byte, 0, 20)

	return strconv.AppendUint(dst, i, 10)
}

func UintToBytes(i uint) []byte {
	if strconv.IntSize == 64 {
		return Uint64ToBytes(uint64(i))
	}
	return Uint32ToBytes(uint32(i))
}
package bytesconv

import (
	"fmt"
	"strconv"
	"testing"
)

func TestInt8ToBytes(t *testing.T) {
	Int8ToBytes(1)
}

func TestInt16ToBytes(t *testing.T) {
	Int16ToBytes(1)
}

func TestInt32ToBytes(t *testing.T) {
	Int32ToBytes(1)
}

func TestInt64ToBytes(t *testing.T) {
	bytes := Int64ToBytes(-9223372036854775808)
	fmt.Println(len(bytes))
}

func TestIntToBytes(t *testing.T) {
	IntToBytes(1)
}

func TestUint8ToBytes(t *testing.T) {
	Uint8ToBytes(uint8(1))
}

func TestUint16ToBytes(t *testing.T) {
	Uint16ToBytes(uint16(1))
}

func TestUint32ToBytes(t *testing.T) {
	Uint32ToBytes(uint32(1))
}

func TestUint64ToBytes(t *testing.T) {
	fmt.Println(Uint64ToBytes(uint64(18446744073709551615)))
}

func TestUintToBytes(t *testing.T) {
	UintToBytes(uint(1))
}

func BenchmarkInt8ToBytes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Int8ToBytes(1)
	}
}

func BenchmarkInt8ToBytes_strconv(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = []byte(strconv.Itoa(1))
	}
}

func BenchmarkIntToBytes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Itobytes(i)
	}
}

func BenchmarkStrconvIntToBytes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = []byte(strconv.Itoa(i))
	}
}

func BenchmarkUintToBytes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UintToBytes(uint(i))
	}
}

func BenchmarkStrconvUintToBytes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = []byte(strconv.FormatInt(int64(i), 10))
	}
}
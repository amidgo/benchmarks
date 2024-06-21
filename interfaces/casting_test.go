package interfaces_test

import (
	"io"
	"testing"
)

func BenchmarkCloseWithCastingToStruct(b *testing.B) {
	for range b.N {
		var s any = newStruct()
		closeWithCastingToStruct(s)
	}
}

func BenchmarkCloseWithCastingToInterface(b *testing.B) {
	for range b.N {
		var s any = newStruct()
		closeWithCastingToInterface(s)
	}
}

func BenchmarkCloseWithoutCasting(b *testing.B) {
	for range b.N {
		var s *SampleStruct = newStruct()
		s.Close()
	}
}

//go:noinline
func closeWithCastingToStruct(s any) {
	ss := s.(*SampleStruct)

	ss.Close()
}

//go:noinline
func closeWithCastingToInterface(s any) {
	closer := s.(io.Closer)

	closer.Close()
}

//go:noinline
func newStruct() *SampleStruct {
	return &SampleStruct{}
}

type SampleStruct struct {
	UintArray  []uint64
	IntArray   []int64
	FloatArray []float32
	LongArray  []float64
}

func (s *SampleStruct) Close() error {
	return nil
}

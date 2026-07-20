package mat

import (
	"encoding/binary"
	"errors"
	"io"
)

const version uint32 = 0x1

const maxLen = int64(int(^uint(0) >> 1))

var (
	headerSize  = binary.Size(storage{})
	sizeFloat64 = binary.Size(float64(0))

	errWrongType = errors.New("mat: wrong data type")

	errTooBig    = errors.New("mat: resulting data slice too big")
	errTooSmall  = errors.New("mat: input slice too small")
	errBadBuffer = errors.New("mat: data buffer size mismatch")
	errBadSize   = errors.New("mat: invalid dimension")
)

func (m Dense) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m Dense) MarshalBinaryTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *Dense) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (m *Dense) UnmarshalBinaryFrom(r io.Reader) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (v VecDense) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (v VecDense) MarshalBinaryTo(w io.Writer) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (v *VecDense) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (v *VecDense) UnmarshalBinaryFrom(r io.Reader) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type storage struct {
	Version uint32
	Form    byte
	Packing byte
	Uplo    byte
	Unit    bool
	Rows    int64
	Cols    int64
	KU      int64
	KL      int64
}

func (s storage) marshalBinaryTo(w io.Writer) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *storage) unmarshalBinary(buf []byte) error { _ = "STUB: not implemented"; return nil }

func (s *storage) unmarshalBinaryFrom(r io.Reader) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func readFull(r io.Reader, buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

package gocodec

import (
	"fmt"
	"reflect"
	"unsafe"
)

type GocEncoder struct {
	cfg       *frozenConfig
	// there are two pointers being written to
	// buf + ptrOffset => the place where the pointer will be updated
	// buf + len(buf) => the place where actual content of pointer will be appended to
	buf       []byte
	ptrOffset uintptr
	Error     error
}

func (cfg *frozenConfig) NewGocEncoder(buf []byte) *GocEncoder {
	return &GocEncoder{cfg: cfg, buf: buf}
}

func (encoder *GocEncoder) Reset(buf []byte) {
	encoder.buf = buf
	encoder.ptrOffset = 0
}

// buf + ptrOffset
func (encoder *GocEncoder) ptr() unsafe.Pointer {
	buf := encoder.buf[encoder.ptrOffset:]
	return ptrOfSlice(unsafe.Pointer(&buf))
}

func (encoder *GocEncoder) EncodeVal(val interface{}) {
	typ := reflect.TypeOf(val)
	valEncoder, err := encoderOfType(encoder.cfg, typ)
	if err != nil {
		encoder.ReportError("EncodeVal", err)
		return
	}
	valEncoder.Encode(ptrOfEmptyInterface(val), encoder)
}

func (encoder *GocEncoder) Buffer() []byte {
	return encoder.buf
}

func (encoder *GocEncoder) ReportError(operation string, err error) {
	if encoder.Error != nil {
		return
	}
	encoder.Error = fmt.Errorf("%s: %s", operation, err)
}

func (encoder *GocEncoder) EncodeInt(val int) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[8]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeInt8(val int8) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[1]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeInt16(val int16) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[2]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeInt32(val int32) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[4]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeInt64(val int64) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[8]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeUint(val uint) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[8]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeUint8(val uint8) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[1]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeUint16(val uint16) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[2]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeUint32(val uint32) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[4]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeUint64(val uint64) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[8]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeUintptr(val uintptr) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[8]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeFloat32(val float32) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[4]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeFloat64(val float64) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[8]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}
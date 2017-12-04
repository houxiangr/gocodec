package gocodec

import (
	"reflect"
	"fmt"
	"unsafe"
)

type GocDecoder struct {
	cfg   *frozenConfig
	buf   []byte
	Error error
}

func (cfg *frozenConfig) NewGocDecoder(buf []byte) *GocDecoder {
	return &GocDecoder{cfg: cfg, buf: buf}
}

func (decoder *GocDecoder) DecodeVal(objPtr interface{}) {
	typ := reflect.TypeOf(objPtr)
	valDecoder, err := decoderOfType(decoder.cfg, typ.Elem())
	if err != nil {
		decoder.ReportError("DecodeVal", err)
		return
	}
	valDecoder.Decode(ptrOfEmptyInterface(objPtr), decoder)
}

func (decoder *GocDecoder) ReportError(operation string, err error) {
	if decoder.Error != nil {
		return
	}
	decoder.Error = fmt.Errorf("%s: %s", operation, err)
}

func (decoder *GocDecoder) DecodeInt() int {
	bufPtr := ptrOfSlice(unsafe.Pointer(&decoder.buf))
	val := *(*int)(bufPtr)
	decoder.buf = decoder.buf[8:]
	return val
}

func (decoder *GocDecoder) DecodeInt8() int8 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&decoder.buf))
	val := *(*int8)(bufPtr)
	decoder.buf = decoder.buf[1:]
	return val
}

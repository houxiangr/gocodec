package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_int64(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(int64(100))
	should.Nil(err)
	should.Equal([]byte{100, 0, 0, 0, 0, 0, 0, 0}, encoded)
	var val int64
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal(int64(100), val)
	encoder := gocodec.DefaultConfig.NewGocEncoder(encoded)
	encoder.EncodeInt64(-1)
	should.Nil(encoder.Error)
	encoded = encoder.Buffer()
	should.Equal([]byte{100, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, encoded)
	decoder := gocodec.DefaultConfig.NewGocDecoder(encoded)
	should.Equal(int64(100), decoder.DecodeInt64())
	should.Equal(int64(-1), decoder.DecodeInt64())
	should.Nil(decoder.Error)
}
package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/esdb/gocodec"
)

func Test_byte_slice(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal([]byte("hello"))
	should.Nil(err)
	should.Equal([]byte{
		0x18, 0, 0, 0, 0, 0, 0, 0,
		5, 0, 0, 0, 0, 0, 0, 0,
		5, 0, 0, 0, 0, 0, 0, 0,
		'h', 'e', 'l', 'l', 'o'}, encoded[8:])
	decoded, err := gocodec.ReadonlyConfig.Unmarshal(encoded, (*[]byte)(nil))
	should.Nil(err)
	should.Equal([]byte("hello"), *decoded.(*[]byte))
	decoded, err = gocodec.Unmarshal(encoded, (*[]byte)(nil))
	should.Nil(err)
	should.Equal([]byte("hello"), *decoded.(*[]byte))
}
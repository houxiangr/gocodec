package test

import (
	"testing"
	"github.com/willf/bloom"
	"github.com/stretchr/testify/require"
	"github.com/esdb/gocodec"
	"io/ioutil"
	"github.com/edsrzf/mmap-go"
	"os"
	"github.com/json-iterator/go"
)

func Test_bloomfilter(t *testing.T) {
	should := require.New(t)
	f := bloom.New(1000*1024, 4)
	f.Add([]byte("hello"))
	f.Add([]byte("world"))
	should.True(f.Test([]byte("hello")))
	should.False(f.Test([]byte("hi")))
	encoded, err := gocodec.Marshal(*f)
	should.Nil(err)
	should.NotNil(encoded)
	ioutil.WriteFile("/tmp/bloomfilter.bin", encoded, 0666)
	var f2 bloom.BloomFilter
	should.Nil(gocodec.Unmarshal(encoded, &f2))
	should.True(f2.Test([]byte("hello")))
	should.False(f2.Test([]byte("hi")))
}

func Test_mmap(t *testing.T) {
	should := require.New(t)
	f, err := os.Open("/tmp/bloomfilter.bin")
	should.Nil(err)
	mem, err := mmap.Map(f, mmap.COPY, 0)
	should.Nil(err)
	var f2 bloom.BloomFilter
	should.Nil(gocodec.Unmarshal(mem, &f2))
	should.True(f2.Test([]byte("hello")))
	should.False(f2.Test([]byte("hi")))
	mem.Unmap()
}

func Test_json(t *testing.T) {
	should := require.New(t)
	f := bloom.New(1000*1024, 4)
	f.Add([]byte("hello"))
	f.Add([]byte("world"))
	encoded, err := jsoniter.Marshal(f)
	should.Nil(err)
	ioutil.WriteFile("/tmp/bloomfilter.json", encoded, 0666)
}

func Benchmark(b *testing.B) {
	b.Run("gocodec", func(b *testing.B) {
		f, _ := os.OpenFile("/tmp/bloomfilter.bin", os.O_RDONLY, 0)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var f2 bloom.BloomFilter
			mem, _ := mmap.Map(f, mmap.COPY, 0)
			err := gocodec.Unmarshal(mem, &f2)
			if err != nil {
				b.Error(err)
			}
			mem.Unmap()
		}
	})
	b.Run("json", func(b *testing.B) {
		f, _ := os.Open("/tmp/bloomfilter.json")
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var f2 bloom.BloomFilter
			//bytes, _ := ioutil.ReadFile("/tmp/bloomfilter.bin")
			mem, _ := mmap.Map(f, mmap.COPY, 0)
			err := jsoniter.Unmarshal(mem, &f2)
			if err != nil {
				b.Error(err)
			}
			mem.Unmap()
		}
	})
}
package lzsc

import (
	"bytes"
	"math/rand"
	"testing"
)

const (
	PkgSize = 2048
	MaxMTU  = 1500
)

func TestLZS(t *testing.T) {
	var n int
	var err error
	for i := 1; i < MaxMTU; i++ {
		pkgBuf := randBytes(i)
		comprBuf := make([]byte, PkgSize)
		n, err = Compress(pkgBuf, comprBuf)
		if err != nil {
			t.Errorf("Compress failed: %d %s", i, err.Error())
		}
		unprBuf := make([]byte, i)
		n, err = Uncompress(comprBuf[:n], unprBuf)
		if err != nil {
			t.Errorf("Uncompress failed: %d %s", i, err.Error())
		}
		if !bytes.Equal(pkgBuf[:i], unprBuf[:n]) {
			t.Errorf("Compress and uncompress data not equal")
		}
	}
}

func BenchmarkCompress(b *testing.B) {
	buf := randBytes(1500)
	comprBuf := make([]byte, len(buf))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Compress(buf, comprBuf)
	}
}

func BenchmarkUncompress(b *testing.B) {
	buf := randBytes(1500)
	comprBuf := make([]byte, len(buf))
	Compress(buf, comprBuf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Uncompress(comprBuf, buf)
	}
}

func randBytes(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte(rand.Intn(256))
	}
	return b
}

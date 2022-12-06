package lzsc

/*
#include <c/lzs.h>
#include <c/lzs.c>
*/
import "C"
import (
	"errors"
	"unsafe"
)

const (
	ErrEFBIG   = "content too large"
	ErrEINVAL  = "invalid argument"
	ErrZero    = "result length is zero"
	ErrUnknown = "unknown error"
)

func Compress(src []byte, dst []byte) (int, error) {
	n := int(C.lzs_compress((*C.uchar)(unsafe.Pointer(&dst[0])), (C.int)(cap(dst)), (*C.uchar)(unsafe.Pointer(&src[0])), (C.int)(len(src))))
	return n, parseErr(n)
}

func Uncompress(src []byte, dst []byte) (int, error) {
	n := int(C.lzs_decompress((*C.uchar)(unsafe.Pointer(&dst[0])), (C.int)(cap(dst)), (*C.uchar)(unsafe.Pointer(&src[0])), (C.int)(len(src))))
	return n, parseErr(n)
}

func parseErr(n int) error {
	switch {
	case n > 0:
		return nil
	case n == -C.EFBIG:
		return errors.New(ErrEFBIG)
	case n == -C.EINVAL:
		return errors.New(ErrEINVAL)
	case n == 0:
		return errors.New(ErrZero)
	default:
		return errors.New(ErrUnknown)
	}
}

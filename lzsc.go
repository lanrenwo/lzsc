package lzsc

/*
#include <./c/lzs.h>
#include <./c/lzs.c>
*/
import "C"
import (
	"unsafe"
)

func Compress(src []byte, dst []byte) int {
	ret := C.lzs_compress((*C.uchar)(unsafe.Pointer(&dst[0])), (C.int)(cap(dst)), (*C.uchar)(unsafe.Pointer(&src[0])), (C.int)(len(src)))
	return (int)(ret)
}

func Uncompress(src []byte, dst []byte) int {
	ret := C.lzs_decompress((*C.uchar)(unsafe.Pointer(&dst[0])), (C.int)(cap(dst)), (*C.uchar)(unsafe.Pointer(&src[0])), (C.int)(len(src)))
	return (int)(ret)
}

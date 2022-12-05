# lzsc
lzsc comes from the OpenConnect's LZS library. 

Using CGO, which makes it easier for you to call in Golang, and it is very fast.

# Installation
```
go get github.com/lanrenwo/lzsc
```
# How to use
```
package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/lanrenwo/lzsc"
)

func main() {
	s := "hello world"
	src := []byte(strings.Repeat(s, 50))

	comprBuf := make([]byte, 2048)
	ret := lzsc.Compress(src, comprBuf)
	if ret <= 0 {
		fmt.Printf("Compress failed: %d", ret)
		return
	}
  
	unprBuf := make([]byte, 2048)
	ret = lzsc.Uncompress(comprBuf[:ret], unprBuf)
	if ret <= 0 {
		fmt.Printf("Uncompress failed: %d", ret)
		return
	}
  
	if !bytes.Equal(src, unprBuf[:ret]) {
		fmt.Printf("Uncompress failed: %d", ret)
		return
	}
  
	fmt.Println("ok")
}

```
# Benchmarks
* lzsgo: translated by c2go, and optimized, the pressure test effect is good, the fluctuation is large in the actual environment
* lzsc: In the actual environment, it is very stable and fast

<img width="858" alt="image" src="https://user-images.githubusercontent.com/3632406/205533834-765b6cfc-7150-47d5-b305-ab63180d6f61.png">


# Thanks
[OpenConnect](https://gitlab.com/openconnect/)

package geo

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
)

func buildGeoDB() ([]byte, []byte, []byte) {
	var lens = []int{1447004, 2026092, 0}
	var db2 = []byte("")
	return decompress(db0, lens[0]), decompress(db1, lens[1]), db2
}


func decompress(b []byte, lens int) []byte {
	zr, e := zlib.NewReader(bytes.NewReader(b))
	if e != nil {
		panic(e)
	}
	var nw int
	w := make([]byte, lens)
	for nw < lens {
		n, e := zr.Read(w[nw:])
		nw += n
		if e != nil {
			if e == io.EOF {
				break
			}
			panic(e)
		}
	}
	if nw != lens {
		panic(fmt.Errorf("expected len=%d but read len=%d", lens, nw))
	}
	return w
}
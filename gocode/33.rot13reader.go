package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// 转换byte  前进13位/后退13位
func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i := range b {
		switch {
		case b[i] >= 'A' && b[i] <= 'Z':
			b[i] = ((b[i] - 'A' + 13) % 26) + 'A'
		case b[i] >= 'a' && b[i] <= 'z':
			b[i] = ((b[i] - 'a' + 13) % 26) + 'a'
		default:
		}
	}
	return
}

func main() {
	// s := strings.NewReader("H jnag gb cynl. Yrg'f cynl!")
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

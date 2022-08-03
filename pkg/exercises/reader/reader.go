package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello, Reader!")

	bit := make([]byte, 8)
	for {
		n, err := reader.Read(bit)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, bit)
		fmt.Printf("b[:n] = %q\n", bit[:n])
		if err == io.EOF {
			break
		}
	}
}
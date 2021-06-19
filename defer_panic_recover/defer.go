package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	a() // Prints 0

	b() // Prints 3 2 1 0
	fmt.Println()

	fmt.Printf("C() : %d\n", c()) // Prints 2
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(strconv.Itoa(i) + " ")
	}
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	written, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}

func CopyFileWithDefer(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

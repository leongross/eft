package main

import (
	"debug/elf"
	"fmt"
	"io"
	"os"
)

func ioReader(filePath string) (io.ReaderAt, error) {
	r, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func byte_to_escaped(bytes []byte) {
	for _, b := range bytes {
		fmt.Printf("\\x%02x", b)
	}
	fmt.Println()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <elf>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := ioReader(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	elf, err := elf.NewFile(f)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	textSection := elf.Section(".text")
	if textSection == nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	textBytes, err := textSection.Data()
	if err != nil {
		fmt.Printf("%v", err)
	}

	byte_to_escaped(textBytes)
}

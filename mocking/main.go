package main

import (
	"io"
	"os"
)

func Countdown(writer io.Writer) {
	writer.Write([]byte("3"))
}

func main() {
	Countdown(os.Stdout)
}

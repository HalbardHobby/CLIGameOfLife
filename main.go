package main

import (
	"fmt"
	"time"

	"bufio"
	"bytes"
	"os"
)

var nWorldWidth, nWorldHeight int = 120, 40

var Output *bufio.Writer = bufio.NewWriter(os.Stdout)

var Screen *bytes.Buffer = new(bytes.Buffer)

func main() {
	Clear() // Clear current screen

	for {
		fmt.Fprintln(Screen, "Current Time:", time.Now().Format(time.RFC1123))
		WriteConsoleOutputFrame()
	}
}

func Clear() {
	Output.WriteString("\033[2J")
}

func WriteConsoleOutputFrame() {
	fmt.Fprintf(Screen, "\033[%d;%dH", 1, 1)

	Output.WriteString(Screen.String())

	Output.Flush()
	Screen.Reset()
}

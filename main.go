package main

import (
	"fmt"

	"bufio"
	"bytes"
	"os"
)

var nWorldWidth, nWorldHeight int = 50, 25

var Output *bufio.Writer = bufio.NewWriter(os.Stdout)

var Screen *bytes.Buffer = new(bytes.Buffer)

func main() {
	w := NewWorld(nWorldWidth, nWorldHeight)

	for {
		Clear()
		fmt.Fprint(w, "Holi, 😸")
		fmt.Fprintln(Screen, w)
		WriteConsoleOutputFrame()
	}
}

func Clear() {
	Output.WriteString("\033c")
}

func WriteConsoleOutputFrame() {
	Output.WriteString(Screen.String())

	Output.Flush()
	Screen.Reset()
}

package main

import (
	"fmt"

	"bufio"
	"bytes"
	"os"
)

const nWorldWidth, nWorldHeight int = 15, 10
const empty, alive rune = '\u3000', '😸'

var Output *bufio.Writer = bufio.NewWriter(os.Stdout)
var Screen *bytes.Buffer = new(bytes.Buffer)

func main() {
	w := NewWorld(nWorldWidth, nWorldHeight)
	s := InitState()
	s.SetRandomState()

	for {
		fmt.Fprintln(Screen, w)
		WriteConsoleOutputFrame()
	}
}

func Clear() {
	Output.WriteString("\033c")
}

func WriteConsoleOutputFrame() {
	Clear()
	Output.WriteString(Screen.String())

	Output.Flush()
	Screen.Reset()
}

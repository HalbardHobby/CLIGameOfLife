package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"time"
)

const nWorldWidth, nWorldHeight int = 85, 40

var Output *bufio.Writer = bufio.NewWriter(os.Stdout)
var Screen *bytes.Buffer = new(bytes.Buffer)

func main() {
	w := NewWorld(nWorldWidth, nWorldHeight)
	s := InitState()
	s.SetRandomState()

	for {
		fmt.Fprint(w, s.Render())
		fmt.Fprintln(Screen, w)
		WriteConsoleOutputFrame()

		s.IterateState()
		time.Sleep(time.Second / 10)
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

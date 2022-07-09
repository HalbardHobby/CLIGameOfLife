package main

import (
	"bytes"
	"strings"
)

const DEFAULT_BORDER = "-- │ ┌ ┐ └ ┘"

type World struct {
	Buf *bytes.Buffer

	Width  int
	Height int
	Border string
}

func NewWorld(width, height int) *World {
	w := new(World)

	w.Buf = new(bytes.Buffer)
	w.Width = width
	w.Height = height
	w.Border = DEFAULT_BORDER

	return w
}

func (w *World) Write(p []byte) (int, error) {
	w.Buf.Reset()
	return w.Buf.Write(p)
}

func (w *World) String() (out string) {
	borders := strings.Split(w.Border, " ")
	lines := strings.Split(w.Buf.String(), "\n")

	prefix, suffix := borders[1], borders[1]

	out += borders[2] + strings.Repeat(borders[0], w.Width) + borders[3] + "\n"
	for y := 0; y < w.Height; y++ {
		var line string

		if y < len(lines) {
			line = prefix + lines[y] + strings.Repeat("\u3000", w.Width-len([]rune(lines[y]))) + suffix
		} else {
			line = prefix + strings.Repeat("\u3000", w.Width) + suffix
		}

		out += line + "\n"
	}

	out += borders[4] + strings.Repeat(borders[0], w.Width) + borders[5]

	return out
}

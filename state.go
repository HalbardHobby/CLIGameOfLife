package main

import (
	"math/rand"
	"time"
)

const empty, alive rune = '\u3000', '😸'

type State struct {
	State_0, State_1        []bool
	CurrentState, NextState []bool
}

func InitState() *State {
	s := new(State)
	s.State_0 = make([]bool, nWorldHeight*nWorldWidth)
	s.State_1 = make([]bool, nWorldHeight*nWorldWidth)

	s.CurrentState = s.State_0[:]
	s.NextState = s.State_1[:]

	return s
}

func (s *State) GetTotalNeighbors(x, y int) (total int) {
	for j := x - 1; j <= x+1; j++ {
		for i := y - 1; i <= y+1; i++ {
			// Check if current coordinate is beyond border
			if !(j < 0 || j >= nWorldWidth || i < 0 || i >= nWorldHeight) {
				// Check if not given coordinates
				if s.CurrentState[(i*nWorldWidth)+j] && !(i == y && j == x) {
					total += 1
				}
			}
		}
	}
	return
}

func (s *State) IterateState() {
	for j := 0; j < nWorldHeight; j++ {
		for i := 0; i < nWorldWidth; i++ {
			neighbors := s.GetTotalNeighbors(i, j)
			if s.CurrentState[(j*nWorldWidth)+i] {
				s.NextState[(j*nWorldWidth)+i] = neighbors >= 2 && neighbors <= 3
			} else {
				s.NextState[(j*nWorldWidth)+i] = neighbors == 3
			}
		}
	}

	s.CurrentState, s.NextState = s.NextState, s.CurrentState
}

func (s *State) SetRandomState() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nWorldWidth*nWorldHeight; i++ {
		s.CurrentState[i] = rand.Intn(5) == 1
	}
}

func (s *State) Render() (out string) {
	for j := 0; j < nWorldHeight; j++ {
		for i := 0; i < nWorldWidth; i++ {
			if s.CurrentState[(j*nWorldWidth)+i] {
				out += string(alive)
			} else {
				out += string(empty)
			}
		}
		out += "\n"
	}
	return
}

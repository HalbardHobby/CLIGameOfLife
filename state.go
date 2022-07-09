package main

import (
	"math/rand"
	"time"
)

type State struct {
	State_0, State_1        [nWorldHeight * nWorldWidth]bool
	CurrentState, NextState []bool
}

func InitState() *State {
	s := new(State)

	for i := 0; i < nWorldHeight*nWorldWidth; i++ {
		s.State_0[i] = false
		s.State_1[i] = false
	}

	s.CurrentState = s.State_0[:]
	s.NextState = s.State_1[:]

	return s
}

func (s *State) GetTotalNeighbors(x, y int) (total int) {
	for j := x - 1; j < x+1; j++ { // check neighbouring cells
		for i := y - 1; i < y+1; i++ {
			// Check if current coordinate is beyond border
			if !(j < 0 || i < 0 || j >= nWorldWidth || i >= nWorldHeight) {
				// Check if not given coordinates
				if s.CurrentState[j*nWorldWidth+i] && !(i == x && j == y) {
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
			if s.CurrentState[j*nWorldWidth+i] {
				s.NextState[j*nWorldWidth+i] = neighbors >= 2 && neighbors <= 3
			} else {
				s.NextState[j*nWorldWidth+i] = neighbors == 3
			}
		}
	}

	s.CurrentState, s.NextState = s.NextState, s.CurrentState
}

func (s *State) SetRandomState() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nWorldWidth*nWorldHeight; i++ {
		s.CurrentState[i] = rand.Intn(3) == 1
	}
}

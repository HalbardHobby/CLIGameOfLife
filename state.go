package main

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
	for i := x - 1; i < x+1; i++ { // check neighbouring cells
		for j := y - 1; j < y+1; j++ {
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

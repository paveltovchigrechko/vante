package league

import "github.com/paveltovchigrechko/vante/internal/match"

type tour struct {
	matches   []*match.Match
	is_played bool
}

type Schedule struct {
	Tours       []tour
	CurrentTour int
}

func (s *Schedule) GetCurrentTour() *tour {
	return &s.Tours[s.CurrentTour]
}

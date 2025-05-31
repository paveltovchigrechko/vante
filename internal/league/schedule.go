package league

import "github.com/paveltovchigrechko/vante/internal/match"

type tour []*match.Match

type Schedule struct {
	Tours       []tour
	CurrentTour int
}

func (s *Schedule) GetCurrentTour() *tour {
	return &s.Tours[s.CurrentTour]
}

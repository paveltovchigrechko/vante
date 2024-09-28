package league

import (
	"fmt"
	"slices"

	"github.com/paveltovchigrechko/vante/internal/match"
	"github.com/paveltovchigrechko/vante/internal/team"
)

type Season struct {
	Teams    []*team.Team
	Schedule *Schedule
}

func (s *Season) GenerateSchedule(rounds int) {
	for round := 1; round <= rounds; round++ {
		s.generateRound()
		slices.Reverse(s.Teams)
	}
}

func (s *Season) generateRound() {
	lastIndex := len(s.Teams) - 1
	mid := len(s.Teams) / 2

	for i := 0; i <= lastIndex-1; i++ {
		var tour tour
		for j := 0; j < mid; j++ {
			host := s.Teams[j]
			guest := s.Teams[lastIndex-j]
			m, _ := match.New(host, guest)
			tour = append(tour, m)
		}

		s.Schedule.Tours = append(s.Schedule.Tours, tour)
		rotate(s.Teams)
	}
}

func (s *Season) PrintSchedule() {
	for tour, matches := range s.Schedule.Tours {
		fmt.Printf("Tour %d\n", tour+1)
		for _, match := range matches {
			match.PrintResult()
		}
	}
}

func (s *Season) Simulate() {
	for _, matches := range s.Schedule.Tours {
		for _, match := range matches {
			match.Simulate()
		}
	}
}

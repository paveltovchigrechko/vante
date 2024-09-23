package league

import (
	"fmt"

	"github.com/paveltovchigrechko/vante/internal/match"
	"github.com/paveltovchigrechko/vante/internal/team"
)

type League struct {
	Name  string
	Teams []*team.Team

	Schedule *Schedule
}

func New(name string, teams []*team.Team) *League {
	return &League{
		Name:  name,
		Teams: teams,
	}
}

// Make this function create schedule for N rounds
func (l *League) GenerateSchedule() error {
	lastIndex := len(l.Teams) - 1
	mid := len(l.Teams) / 2
	tours := &Schedule{}

	for i := 0; i <= lastIndex-1; i++ {
		var tour tour
		for j := 0; j < mid; j++ {
			host := l.Teams[j]
			guest := l.Teams[lastIndex-j]
			m, _ := match.New(host, guest)
			tour = append(tour, m)
		}

		rotate(l.Teams)
		tours.Tours = append(tours.Tours, tour)
	}

	l.Schedule = tours

	return nil
}

func (l *League) PrintSchedule() error {
	for tour, matches := range l.Schedule.Tours {
		fmt.Printf("Tour %d\n", tour+1)
		for _, match := range matches {
			match.PrintResult()
		}
	}

	return nil
}

func rotate(arr []*team.Team) {
	lastIndex := len(arr) - 1
	tail := append(arr[lastIndex:], arr[1:lastIndex]...)
	arr = append(arr[:1], tail...)
}

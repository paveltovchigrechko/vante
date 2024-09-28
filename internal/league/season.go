package league

import (
	"fmt"
	"slices"
	"sort"

	"github.com/paveltovchigrechko/vante/internal/match"
	"github.com/paveltovchigrechko/vante/internal/team"
)

type Season struct {
	Teams      []*team.Team
	Schedule   *Schedule
	Statistics *Statistics
}

func (s *Season) GenerateSchedule(rounds int) {
	for round := 1; round <= rounds; round++ {
		s.generateRound()
		slices.Reverse(s.Teams)
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
			s.addStatistics(match)
		}
	}
	s.makeTable()
}

func (s *Season) addStatistics(m *match.Match) {
	hostStats := s.Statistics.Team[m.Host]
	guestStats := s.Statistics.Team[m.Guest]

	hostStats.team = m.Host.Name
	guestStats.team = m.Guest.Name

	hostStats.Games += 1
	guestStats.Games += 1

	switch {
	case m.HostScore > m.GuestScore:
		hostStats.Wins += 1
		hostStats.Points += 3
		guestStats.Losses += 1
	case m.HostScore < m.GuestScore:
		hostStats.Losses += 1
		guestStats.Wins += 1
		guestStats.Points += 3
	default:
		hostStats.Draws += 1
		hostStats.Points += 1
		guestStats.Draws += 1
		guestStats.Points += 1
	}

	hostStats.GoalsFor += m.HostScore
	hostStats.GoalsAgainst += m.GuestScore
	hostStats.GoalDifference = hostStats.GoalsFor - hostStats.GoalsAgainst

	guestStats.GoalsFor += m.GuestScore
	guestStats.GoalsAgainst += m.HostScore
	guestStats.GoalDifference = guestStats.GoalsFor - guestStats.GoalsAgainst

	s.Statistics.Team[m.Host] = hostStats
	s.Statistics.Team[m.Guest] = guestStats
}

func (s *Season) makeTable() {
	for _, stats := range s.Statistics.Team {
		s.Statistics.Table = append(s.Statistics.Table, stats)
		sort.Slice(s.Statistics.Table, s.Statistics.Table.ByPoints)
		slices.Reverse(s.Statistics.Table)
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

func rotate(arr []*team.Team) {
	lastIndex := len(arr) - 1
	tail := append(arr[lastIndex:], arr[1:lastIndex]...)
	arr = append(arr[:1], tail...)
}

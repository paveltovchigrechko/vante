package league

import (
	"github.com/paveltovchigrechko/vante/internal/team"
)

type League struct {
	Name         string
	CurrentTeams []*team.Team

	Seasons []*Season
}

func New(name string, teams []*team.Team) *League {
	return &League{
		Name:         name,
		CurrentTeams: teams,
		Seasons:      []*Season{},
	}
}

func (l *League) CreateNewSeason() {
	s := &Season{
		Teams:    l.CurrentTeams,
		Schedule: &Schedule{},
	}
	l.Seasons = append(l.Seasons, s)
}

func (l *League) SimulateSeason() {
	curSeason := len(l.Seasons) - 1
	l.Seasons[curSeason].GenerateSchedule(2)
	l.Seasons[curSeason].Simulate()
}

func (l *League) PrintCurrentSeason() {
	curSeason := len(l.Seasons) - 1
	l.Seasons[curSeason].PrintSchedule()
}

func rotate(arr []*team.Team) {
	lastIndex := len(arr) - 1
	tail := append(arr[lastIndex:], arr[1:lastIndex]...)
	arr = append(arr[:1], tail...)
}

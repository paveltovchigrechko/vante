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
		Statistics: &Statistics{
			Team:  make(map[*team.Team]TeamStats, len(l.CurrentTeams)),
			Table: Table{},
		},
	}
	l.Seasons = append(l.Seasons, s)
}

func (l *League) SimulateSeason() {
	curSeason := len(l.Seasons) - 1
	l.Seasons[curSeason].GenerateSchedule(2)
	l.Seasons[curSeason].Simulate()
}

func (l *League) PrintCurrentSchedule() {
	curSeason := len(l.Seasons) - 1
	l.Seasons[curSeason].PrintSchedule()
}

func (l *League) PrintCurrentTable() {
	curSeason := len(l.Seasons) - 1
	l.Seasons[curSeason].Statistics.Table.Print()
}

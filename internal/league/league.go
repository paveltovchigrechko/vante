package league

import (
	"slices"
	"strings"

	"github.com/paveltovchigrechko/vante/internal/team"
)

type League struct {
	Name         string
	CurrentTeams []*team.Team
	Seasons      []*Season
	Rules        Rules
}

func New(name string, teams []*team.Team) *League {
	return &League{
		Name:         name,
		CurrentTeams: teams,
		Seasons:      []*Season{},
		Rules: Rules{
			Rotation{
				principle:       swap,
				relegationTeams: 2,
				promotionTeams:  2,
			},
		},
	}
}

func (l *League) CreateNewSeason() {
	teams := make([]team.Team, 0, len(l.CurrentTeams))
	for _, team := range l.CurrentTeams {
		teams = append(teams, *team)
	}

	s := &Season{
		Teams:    teams,
		Schedule: &Schedule{},
		Statistics: &Statistics{
			Team:  make(map[string]TeamStats, len(l.CurrentTeams)),
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

func (l *League) getRelegated() []*team.Team {
	curSeason := len(l.Seasons) - 1
	offset := len(l.Seasons[curSeason].Teams) - l.Rules.rotation.relegationTeams
	relegated := []*team.Team{}
	for _, stat := range l.Seasons[curSeason].Statistics.Table[offset:] {
		for _, t := range l.CurrentTeams {
			if stat.team == t.Name {
				relegated = append(relegated, t)
			}
		}

	}
	return relegated
}

func (l *League) RemoveRelegated() []*team.Team {
	relegated := l.getRelegated()
	for _, team := range relegated {
		l.removeTeam(team.Name)
	}
	return relegated
}

func (l *League) getPromoted() []*team.Team {
	curSeason := len(l.Seasons) - 1
	promoted := []*team.Team{}
	for _, stat := range l.Seasons[curSeason].Statistics.Table[:l.Rules.rotation.promotionTeams] {
		for _, team := range l.CurrentTeams {
			if stat.team == team.Name {
				promoted = append(promoted, team)
				break
			}
		}
	}

	return promoted
}

func (l *League) RemovePromoted() []*team.Team {
	promoted := l.getPromoted()
	for _, team := range promoted {
		l.removeTeam(team.Name)
	}
	return promoted
}

func (l *League) AddTeams(teams []*team.Team) {
	l.CurrentTeams = append(l.CurrentTeams, teams...)
}

func (l *League) removeTeam(name string) {
	for _, t := range l.CurrentTeams {
		if t == nil {
			continue
		}

		if strings.Compare(t.Name, name) == 0 {
			l.CurrentTeams = slices.DeleteFunc(l.CurrentTeams, func(team *team.Team) bool {
				return team.Name == name
			})
		}
	}
}

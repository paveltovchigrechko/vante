package league

import (
	"fmt"
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
			2,
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
	l.Seasons[curSeason].GenerateSchedule(l.Rules.roundNumber)
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

// Returns a list of the last N teams in the current season (i.e. relegated teams that should go to the lower league). N equals `l.Rules.Potation.relegationTeams`.
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

// Removes from the current league teams the last N teams in the current season and returns them as a list. N equals `l.Rules.Potation.relegationTeams`. Does not remove teams from the current season.
func (l *League) RemoveRelegatedTeams() []*team.Team {
	relegated := l.getRelegated()
	for _, team := range relegated {
		l.removeTeam(team.Name)
	}
	return relegated
}

// Returns a list of the first N teams in the current season (i.e. promoted teams that should go to the upper league). N equals `l.Rules.Potation.promotionTeams`.
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

// Removes from the current league teams the first N teams in the current season and returns them as a list. N equals `l.Rules.Potation.promotionTeams`. Does not remove teams from the current season.
func (l *League) RemovePromotedTeams() []*team.Team {
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

func (l *League) ListTeams() {
	for _, t := range l.CurrentTeams {
		fmt.Println(t.Name)
	}
}

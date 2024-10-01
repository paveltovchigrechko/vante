package league

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/paveltovchigrechko/vante/internal/team"
)

type TeamStats struct {
	team           string
	Games          int
	Points         int
	Wins           int
	Draws          int
	Losses         int
	GoalsFor       int
	GoalsAgainst   int
	GoalDifference int
}

type Statistics struct {
	Team  map[*team.Team]TeamStats
	Table Table
}
type Table []TeamStats

func (t Table) Len() int {
	return len(t)
}

func (t Table) ByPoints(i, j int) bool {
	return t[i].Points < t[j].Points
}

func (t Table) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Table) Print() {
	w := tabwriter.NewWriter(os.Stdout, 1, 4, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintf(w, "#\tTeam\tPlayed\tPoints\tWon\tDrawn\tLost\tGF\tGA\tGD\t\n")
	for i, stat := range t {
		fmt.Fprintf(w, "%d\t%s\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t\n", i+1, stat.team, stat.Games, stat.Points, stat.Wins, stat.Draws, stat.Losses, stat.GoalsFor, stat.GoalsAgainst, stat.GoalDifference)
	}
	w.Flush()
}

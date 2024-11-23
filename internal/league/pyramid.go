package league

// Pyramid represents a hierarchy of leagues.
// The league order dictates the position of a league in the pyramid.
type Pyramid struct {
	Leagues []*League
}

func NewPyramid(leagues []*League) *Pyramid {
	return &Pyramid{
		Leagues: leagues,
	}
}

func (p *Pyramid) SwapTeams() {
	leaguePairs := len(p.Leagues) - 1
	for i := 0; i < leaguePairs; i++ {
		upperLeague := p.Leagues[i]
		lowerLeague := p.Leagues[i+1]
		relegatedTeams := upperLeague.RemoveRelegatedTeams()
		promotedTeams := lowerLeague.RemovePromotedTeams()
		upperLeague.AddTeams(promotedTeams)
		lowerLeague.AddTeams(relegatedTeams)
	}
}

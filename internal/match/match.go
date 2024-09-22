package match

import (
	"fmt"
	"math/rand"

	"github.com/paveltovchigrechko/vante/internal/team"
)

type Match struct {
	Host  *team.Team
	Guest *team.Team

	HostScore  int
	GuestScore int
}

func New(host, guest *team.Team) (*Match, error) {
	if host == guest {
		return nil, fmt.Errorf("The match has the same team %q as host and guest", host.Name)
	}

	return &Match{
		Host:  host,
		Guest: guest,
	}, nil
}

func (m *Match) Simulate() {
	m.HostScore = rand.Intn(5)
	m.GuestScore = rand.Intn(5)
}

func (m *Match) PrintResult() string {
	return fmt.Sprintf("%s - %s %d:%d\n", m.Host.Name, m.Guest.Name, m.HostScore, m.GuestScore)
}

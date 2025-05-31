package team

import "fmt"

type Team struct {
	Name string
	city string
}

func New(name, city string) (*Team, error) {
	if name == "" {
		return nil, fmt.Errorf("the team name must have at least 1 symbol")
	}

	return &Team{
		Name: name,
		city: city,
	}, nil
}

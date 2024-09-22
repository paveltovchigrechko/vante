package team

import "fmt"

type Team struct {
	Name string
}

func New(name string) (*Team, error) {
	if name == "" {
		return nil, fmt.Errorf("The team name must have at least 1 symbol")
	}
	return &Team{
		Name: name,
	}, nil
}

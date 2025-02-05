package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type Player struct {
	Name string
	Wins int
}

type League []Player

func NewLeague(rdr io.Reader) (League, error) {
	var league League

	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		err = fmt.Errorf("problems parsing league, %v", err)
	}

	return league, err
}

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}

	return nil
}

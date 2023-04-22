package main

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	playersWin map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.playersWin[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.playersWin[name]++
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.playersWin { // iteration over map
		league = append(league, Player{name, wins})
	}
	return league
}

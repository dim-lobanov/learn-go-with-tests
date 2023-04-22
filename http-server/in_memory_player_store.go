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
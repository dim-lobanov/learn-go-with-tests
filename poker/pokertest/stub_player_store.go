package pokertest

import "learn-go-with-tests/poker"

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []poker.Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() poker.League {
	return s.league
}

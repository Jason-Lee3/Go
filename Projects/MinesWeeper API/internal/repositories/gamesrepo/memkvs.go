package gamesrepo

import (
	domain "command-line-arguments/Users/jasonlee/Code/Go/Projects/MinesWeeper API/internal/core/domain/game.go"
	"encoding/json"
	"errors"
)

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS() *memkvs {
	return &memkvs{kvs: map[string][]byte{}}
}

func (repo *memkvs) Get(id string) (domain.Game, error) {
	if value, ok := repo.kvs[id]; ok {
		game := domain.Game{}
		err := json.Unmarshal(value, &game)
		if err != nil {
			return domain.Game{}, errors.New("fail to get value from kvs")
		}

		return game, nil
	}

	return domain.Game{}, errors.New("game not found in kvs")
}

package gamesrv

import (
	domain "command-line-arguments/Users/jasonlee/Code/Go/Projects/MinesWeeper API/internal/core/domain/game.go"
	"command-line-arguments/Users/jasonlee/Code/Go/Projects/MinesWeeper API/internal/core/ports/ports.go"
	"errors"
)

type service struct {
	gamesRepository ports.GamesRepository
	uidGen          uidgen.UIDGen
}

func New(gamesRepository ports.GamesRepository, uidGen uidgen.UIDGen) *service {
	return &service{
		gamesRepository: gamesRepository,
		uidGen:          uidGen,
	}
}

func (srv *service) Get(id string) (domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		return domain.Game{}, errors.New("get game from repository has failed")
	}
	return game, nil
}

func (srv *service) Create(name string, size uint, bombs uint) (domain.Game, error) {
	if bombs >= size*size {
		return domain.Game{}, errors.New("the number of bombs is invalid")
	}

	game := domain.NewGame(srv.uidGen.New(), name, size, bombs)
	if err := srv.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New("create game into repository has failed")
	}

	return game, nil
}

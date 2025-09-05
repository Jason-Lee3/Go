package main

import (
	gamesrv "command-line-arguments/Users/jasonlee/Code/Go/Projects/MinesWeeper API/internal/core/services/gamesrv/service.go"
	gamehdl "command-line-arguments/Users/jasonlee/Code/Go/Projects/MinesWeeper API/internal/handlers/gamehdl/http.go"
	gamesrepo "command-line-arguments/Users/jasonlee/Code/Go/Projects/MinesWeeper API/internal/repositories/gamesrepo/memkvs.go"
)

func main() {
	gamesRepository := gamesrepo.NewMemKVS()
	gamesService := gamesrv.New(gamesRepository, uidgen.New())
	gamesHandler := gamehdl.NewHTTPHandler(gamesService)

	router := gin.New()
	router.GET("/games/:id", gamesHandler.Get)
	router.POST("/games", gamesHandler.Create)

	router.Run(":8080")
}

package main

import (
	"github.com/acmacalister/helm"
	"github.com/blainsmith/battleship-api/handlers"
)

func main() {
	router := helm.New(handlers.RootHandler)
	router.Handle("DELETE", "/battleship/game/:gameId", handlers.DeleteGameHandler)
	router.Handle("POST", "/battleship/game/:gameId/shot", handlers.SendShotHandler)
	router.Handle("POST", "/battleship/game/:gameId/shot-result/:shotResult", handlers.SendShotResultHandler)
	router.Handle("POST", "/battleship/game/:gameId/receive-shot/:letter/:number", handlers.ReceiveShotHandler)
	router.Handle("POST", "/battleship/game", handlers.CreateGameHandler)
	router.Run(":8080")
}

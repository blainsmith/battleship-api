package main

import (
	"github.com/acmacalister/helm"
	"github.com/blainsmith/battleship/handlers"
)

func init() {
	// Connect to DB
}

func main() {
	router := helm.New(handlers.RootHandler)
	router.Handle("POST", "/battleship/game", handlers.CreateGameHandler)
	router.Handle("DELETE", "/battleship/game/:gameId", handlers.DeleteGameHandler)
	router.Handle("POST", "/battleship/game/:gameId/shot", handlers.SendShotHandler)
	router.Handle("POST", "/battleship/game/:gameId/shot-result/:shotResult", handlers.SendShotResultHandler)
	router.Handle("POST", "/battleship/game/:gameId/receive-shot/:letter/:number", handlers.ReceiveShotHandler)
	router.Run(":8080")
}

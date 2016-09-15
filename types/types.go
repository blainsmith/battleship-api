package types

type Game struct {
	GameID string `json:"gameId"`
	Grid   string `json:"grid"`
}

type Shot struct {
	Letter string `json:"letter"`
	Number string `json:"number"`
}

type ShotResult struct {
	Result int16 `json:"result"`
}

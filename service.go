package main

import "tennis_player/config"

func UpdatePlayerContainer(id int, playerId int) ([]string,string) {
	const (
		failed  = "failed"
		success = "success"
		full = "full"
	)
	conn := config.ConnectPGDB()
	playerData, stat := UpdateContainerBall(conn, id, playerId)

	if stat == failed {
		return nil, failed
	}
	if stat == full {
		return playerData, full
	}

	return nil, success
}
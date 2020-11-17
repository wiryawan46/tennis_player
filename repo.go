package main

import (
	"github.com/go-pg/pg/v10"
	"log"
	"tennis_player/model"
)

func UpdateContainerBall(conn *pg.DB, id int, playerId int) ([]string,string) {
	const (
		failed  = "failed"
		success = "success"
		full = "full"
	)
	container := model.ContainerBalls{}
	player := model.TennisPlayer{}
	var playerName []string
	data, res := CheckIfContainerIsFull(conn, id)
	if res == "full" {
		for in,_ := range data {
			_, err := conn.QueryOne(&player, `SELECT player_name FROM tennis_players WHERE id = ?`, data[in])
			if err != nil {
				log.Print(err.Error())
				return nil,failed
			}
			playerName = append(playerName, player.PlayerName)
		}
		return playerName, full
	}
	_, err := conn.QueryOne(container, `UPDATE container_balls SET player_id = array_append(player_id, ?) WHERE id = ?`, playerId, id)
	if err != nil {
		log.Print(err.Error())
		return nil, failed
	}

	return nil, success
}

func CheckIfContainerIsFull(conn *pg.DB, id int) ([]int,string) {
	const (
		failed  = "failed"
		full = "full"
	)
	container := model.ContainerBalls{}
	_, err := conn.QueryOne(&container, `SELECT size, player_id FROM container_balls WHERE id = ?`, id)
	if err != nil {
		log.Print(err.Error())
		return nil,failed
	}
	if len(container.PlayerId) == container.Size {
		log.Print("Container already full!")
		return container.PlayerId, full
	}
	return nil,""
}
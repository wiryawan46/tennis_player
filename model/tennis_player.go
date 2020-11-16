package model

type TennisPlayer struct {
	tableName interface{} `pg:"public.tennis_players"`
	Id int `pg:"id"`
	PlayerName string `pg:"player_name"`
}
package model

type ContainerBalls struct {
	tableName interface{} `pg:"public.container_balls"`
	Id int `pg:"id"`
	Size int `pg:"size"`
	PlayerId []int `pg:"player_id,array"`
}
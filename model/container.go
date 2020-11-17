package model

type ContainerBalls struct {
	tableName interface{} `pg:"public.container_balls"`
	Id        int         `pg:"id"`
	Size      int         `pg:"size"`
	PlayerId  []int       `pg:"player_id,array"`
	IsReady   bool        `pg:"is_ready,use_zero"`
}

// ResponseDetail data structure
type ResponseDetail struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

package footstats

type Coach struct {
	FootstatsId int64
	Name        string
}

type footstatsCoach struct {
	FootstatsId string `json:"@Id"`
	Name        string `json:"@Nome"`
}

package footstats

type Referee struct {
	FootstatsId int64
	Name        string
}

type footstatsReferee struct {
	FootstatsId string `json:"@Id"`
	Name        string `json:"@Nome"`
}

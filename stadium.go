package footstats

type Stadium struct {
	FootstatsId   int64
	Name          string
	City          string
	State         string
	IsPlaceholder bool
}

type footstatsStadium struct {
	FootstatsId string `json:"@Id"`
	Name        string `json:"@Nome"`
	City        string `json:"@Cidade"`
	State       string `json:"@Estado"`
}

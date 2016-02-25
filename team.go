package footstats

type Team struct {
	FootstatsId   int64
	Name          string
	Initials      string
	IsPlaceholder bool
}

type footstatsTeam struct {
	FootstatsId   string `json:"@Id"`
	Name          string `json:"@Nome"`
	Initials      string `json:"@Sigla"`
	IsPlaceholder string `json:"@EquipeFantasia"`
}

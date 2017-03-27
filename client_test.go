package footstats

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"

	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	token  = "footstats-token"
	client *Client
)

func init() {
	mux = http.NewServeMux()

	server = httptest.NewServer(mux)
	client = &Client{
		baseURL: server.URL,
		token:   token,
	}
}

func copyFileToResponse(w http.ResponseWriter, filename string) {
	f, _ := os.Open(filename)
	io.Copy(w, f)
}

func TestURLBuilding(t *testing.T) {
	u := client.buildURL("test-endpoint", nil)

	if !strings.Contains(u, "token="+token) {
		t.Error("missing token")
	}

	params := &url.Values{}
	params.Set("key", "value")
	u = client.buildURL("test-endpoint", params)

	if !strings.Contains(u, "key=value") {
		t.Error("params were not included")
	}
}

func TestChampionships(t *testing.T) {
	mux.HandleFunc("/V2/api/Campeonato/ListarCampeonatos", func(w http.ResponseWriter, r *http.Request) {
		copyFileToResponse(w, "api-samples/championships.json")
	})

	championships, err := client.Championships()
	if err != nil {
		t.Fatal("unable to retrieve championships:", err)
	}

	if clen := len(championships); clen != 3 {
		t.Error("championship count mismatch: expected", 3, "got", clen)
	}
}

func TestPlayersByTeam(t *testing.T) {
	mux.HandleFunc("/V2/api/Jogador/JogadoresEquipe", func(w http.ResponseWriter, r *http.Request) {
		copyFileToResponse(w, "api-samples/players-by-team.json")
	})

	players, err := client.PlayersByTeam(1006)
	if err != nil {
		t.Fatal("unable to retrieve players by team:", err)
	}

	if plen := len(players); plen != 20 {
		t.Error("player count mismatch: expected", 20, "got", plen)
	}
}

func TestTeamsByChampionship(t *testing.T) {
	mux.HandleFunc("/V2/api/Equipe/EquipesCampeonato", func(w http.ResponseWriter, r *http.Request) {
		copyFileToResponse(w, "api-samples/teams.json")
	})

	teams, err := client.TeamsByChampionship(515)
	if err != nil {
		t.Fatal("unable to retrieve teams by championship:", err)
	}

	if tlen := len(teams); tlen != 4 {
		t.Error("team count mismatch: expected", 4, "got", tlen)
	}

}

func TestMatchesByChampionship(t *testing.T) {
	mux.HandleFunc("/V2/api/Partida/PartidasCampeonato", func(w http.ResponseWriter, r *http.Request) {
		copyFileToResponse(w, "api-samples/matches.json")
	})

	matches, err := client.MatchesByChampionship(515)
	if err != nil {
		t.Fatal("unable to retrieve matches by championship:", err)
	}

	if mlen := len(matches); mlen != 5 {
		t.Error("match count mismatch: expected", 5, "got", mlen)
	}
}

func TestOnGoingMatches(t *testing.T) {
	mux.HandleFunc("/V2/api/Partida/PartidasAndamento", func(w http.ResponseWriter, r *http.Request) {
		copyFileToResponse(w, "api-samples/ongoing-matches.json")
	})

	matches, err := client.OnGoingMatches()
	if err != nil {
		t.Fatal("unable to retrieve ongoing matches:", err)
	}

	if mlen := len(matches); mlen != 2 {
		t.Error("match count mismatch: expected", 2, "got", mlen)
	}
}

func TestMatchFeed(t *testing.T) {
	mux.HandleFunc("/V2/api/Partida/NarracaoMinMin/", func(w http.ResponseWriter, r *http.Request) {
		copyFileToResponse(w, "api-samples/match-feed.json")
	})

	feed, err := client.MatchFeed(125718)
	if err != nil {
		t.Fatal("unable to retrieve match feed:", err)
	}

	if flen := len(feed); flen != 55 {
		t.Error("feed count mismatch: expected", 55, "got", flen)
	}
}

func TestMatchLineup(t *testing.T) {
	mux.HandleFunc("/V2/api/Partida/Escalacao", func(w http.ResponseWriter, r *http.Request) {
		copyFileToResponse(w, "api-samples/match-lineup.json")
	})

	lineup, err := client.MatchLineup(125738)
	if err != nil {
		t.Fatal("unable to retrieve match lineup:", err)
	}

	if htlen := len(lineup.HomeTeamLineup); htlen != 14 {
		t.Error("home team lineup count mismatch: expected", htlen, "got", htlen)
	}

	if vtlen := len(lineup.VisitingTeamLineup); vtlen != 14 {
		t.Error("visiting team lineup count mismatch: expected", 14, "got", vtlen)
	}
}

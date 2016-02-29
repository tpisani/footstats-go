package footstats

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"testing"
)

var (
	mux      *http.ServeMux
	client   *Client
	username string = "username"
	password string = "password"
	server   *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(server.URL, username, password)
}

func teardown() {
	server.Close()
}

func writeFileToResponse(w http.ResponseWriter, filename string) {
	f, _ := os.Open(filename)
	b, _ := ioutil.ReadAll(f)
	w.Write(b)
}

func TestAuth(t *testing.T) {
	setup()
	defer teardown()

	url := client.buildURL("test-endpoint", nil)

	if !strings.Contains(url, "usuario="+username) {
		t.Error("Missing username")
	}

	if !strings.Contains(url, "senha="+password) {
		t.Error("Missing password")
	}
}

func TestChampionships(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/ListaCampeonatos", func(w http.ResponseWriter, r *http.Request) {
		writeFileToResponse(w, "api-samples/championships.xml")
	})

	championships, err := client.Championships()
	if err != nil {
		t.Error("Unable to retrieve championships:", err)
	}

	if clen := len(championships); clen != 2 {
		t.Error("Expected 2 championships, got", clen)
	}
}

func TestMatches(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/ListaPartidas", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("campeonato")
		if id != "434" {
			t.Error("Expected championship ID 434, got", id)
		}

		writeFileToResponse(w, "api-samples/matches.xml")
	})

	matches, err := client.Matches(434)
	if err != nil {
		t.Error("Unable to retrieve matches:", err)
	}

	if mlen := len(matches); mlen != 4 {
		t.Error("Expected 4 matches, got", mlen)
	}
}

func TestEntities(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/ListaEntidades", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("campeonato")
		if id != "434" {
			t.Error("Expected championship ID 434, got", id)
		}

		writeFileToResponse(w, "api-samples/entities.xml")
	})

	entities, err := client.Entities(434)
	if err != nil {
		t.Error("Unable to retrieve entities:", err)
	}

	if elen := len(entities.Teams()); elen != 6 {
		t.Error("Expected 6 teams, got", elen)
	}

	if plen := len(entities.Players()); plen != 4 {
		t.Error("Expected 4 players, got", plen)
	}

	if clen := len(entities.Coaches()); clen != 2 {
		t.Error("Expected 2 coaches, got", clen)
	}

	if rlen := len(entities.Referees()); rlen != 2 {
		t.Error("Expected 2 referees, got", rlen)
	}

	if slen := len(entities.Stadiums()); slen != 2 {
		t.Error("Expected 2 stadiums, got", slen)
	}
}

func TestLiveData(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/AoVivo", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("idpartida")
		if id != "10999" {
			t.Error("Expected match ID 10999, got", id)
		}

		writeFileToResponse(w, "api-samples/live.xml")
	})

	live, err := client.LiveData(10999)
	if err != nil {
		t.Error("Unable to retrieve goals:", err)
	}

	if glen := len(live.Goals()); glen != 2 {
		t.Error("Expected 3 goals, got", glen)
	}
}

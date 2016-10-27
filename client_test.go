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
	server   *httptest.Server
	username string = "username"
	password string = "password"
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
		t.Error("missing username")
	}

	if !strings.Contains(url, "senha="+password) {
		t.Error("missing password")
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
		t.Fatal("unable to retrieve championships:", err)
	}

	if clen := len(championships); clen != 2 {
		t.Error("expected 2 championships, got", clen)
	}
}

func TestMatches(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/ListaPartidas", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("campeonato")
		if id != "434" {
			t.Error("expected championship ID 434, got", id)
		}

		writeFileToResponse(w, "api-samples/matches.xml")
	})

	matches, err := client.Matches(434)
	if err != nil {
		t.Fatal("unable to retrieve matches:", err)
	}

	if mlen := len(matches); mlen != 4 {
		t.Error("expected 4 matches, got", mlen)
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
		t.Fatal("unable to retrieve entities:", err)
	}

	if elen := len(entities.Teams); elen != 6 {
		t.Error("expected 6 teams, got", elen)
	}

	if plen := len(entities.Players); plen != 4 {
		t.Error("expected 4 players, got", plen)
	}

	if clen := len(entities.Coaches); clen != 2 {
		t.Error("expected 2 coaches, got", clen)
	}

	if rlen := len(entities.Referees); rlen != 2 {
		t.Error("expected 2 referees, got", rlen)
	}

	if slen := len(entities.Stadiums); slen != 2 {
		t.Error("expected 2 stadiums, got", slen)
	}
}

func TestMatchEvents(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/AoVivo", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("idpartida")
		if id != "10999" {
			t.Error("expected match ID 10999, got", id)
		}

		writeFileToResponse(w, "api-samples/live.xml")
	})

	events, err := client.MatchEvents(10999)
	if err != nil {
		t.Fatal("unable to retrieve match events:", err)
	}

	if glen := len(events.Goals); glen != 2 {
		t.Error("expected 2 goals, got", glen)
	}

	if clen := len(events.Cards); clen != 3 {
		t.Error("expected 3 cards, got", clen)
	}

}

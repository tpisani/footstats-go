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
		t.Error("Unable to retrieve championships")
	}

	if clen := len(championships); clen != 2 {
		t.Error("Expected 2 championships, got", clen)
	}
}

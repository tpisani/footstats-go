package footstats

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(server.URL, "username", "password")
}

func teardown() {
	server.Close()
}

func writeFileToResponse(w http.ResponseWriter, name string) {
	f, _ := os.Open("api-samples/championships.xml")
	b, _ := ioutil.ReadAll(f)
	w.Write(b)
}

func TestChampionships(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/ListaCampeonatos", func(w http.ResponseWriter, r *http.Request) {
		writeFileToResponse(w, "api-samples/championships.xml")
	})

	_, err := client.Championships()
	if err != nil {
		t.Fail()
	}
}

package discogs

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	mux           *http.ServeMux
	server        *httptest.Server
	testAuthToken string
	testUserAgent string
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	testAuthToken = "1234"
	testUserAgent = "GoDiscogs/0.0.1 +http://test.net"
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()

	if got := r.Method; got != want {
		t.Errorf("request method: %v, want %v", got, want)
	}
}

func testEqual(t *testing.T, want interface{}, got interface{}) {
	t.Helper()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("values not equal (-want / +got):\n%s", diff)
	}
}

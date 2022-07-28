package discogs

import (
	"context"
	"net/http"
	"testing"

	"github.com/peacefixation/go-discogs/discogs/auth"
)

func Test_GetIdentity(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/oauth/identity", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		_, _ = w.Write([]byte(`{"id": 0, "username": "Username", "resource_url": "https://api.discogs.com/users/Username", "consumer_name": "Username"}`))
	})

	client := NewClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	identity, err := client.GetIdentity(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := &Identity{
		ID:           0,
		Username:     "Username",
		ResourceURL:  "https://api.discogs.com/users/Username",
		ConsumerName: "Username",
	}

	testEqual(t, want, identity)
}

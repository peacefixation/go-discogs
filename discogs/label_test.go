package discogs

import (
	"context"
	"net/http"
	"testing"

	"github.com/peacefixation/go-discogs/discogs/auth"
)

func Test_GetLabel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/labels/1611434", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		_, _ = w.Write([]byte(`{"id": 1611434, "name": "Forgotten World", "resource_url": "https://api.discogs.com/labels/1611434", "uri": "https://www.discogs.com/label/1611434-Forgotten-World", "releases_url": "https://api.discogs.com/labels/1611434/releases", "parent_label": {"id": 1447907, "name": "All Possible Worlds", "resource_url": "https://api.discogs.com/labels/1447907"}, "data_quality": "Needs Vote"}`))
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	label, err := client.GetLabel(context.Background(), 1611434)
	if err != nil {
		t.Fatal(err)
	}

	want := &Label{
		ID:          1611434,
		Name:        "Forgotten World",
		ResourceURL: "https://api.discogs.com/labels/1611434",
		URI:         "https://www.discogs.com/label/1611434-Forgotten-World",
		ReleasesURL: "https://api.discogs.com/labels/1611434/releases",
		ParentLabel: ParentLabel{
			ID:          1447907,
			Name:        "All Possible Worlds",
			ResourceURL: "https://api.discogs.com/labels/1447907",
		},
		DataQuality: "Needs Vote",
	}

	testEqual(t, want, label)
}

func Test_GetLabelReleases(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/labels/1611434/releases", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		_, _ = w.Write([]byte(`{"pagination": {"page": 1, "pages": 1, "per_page": 50, "items": 2, "urls": {}}, "releases": [{"status": "Accepted", "format": "12\"", "catno": "fw1", "thumb": "https://i.discogs.com/8FsOyncx8agIs2Kq2xvIL0orEwWdyDeWEpc9Rh5hNC8/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEzMzU4/NzM3LTE1NTI3MzUw/NjctNzUwMy5qcGVn.jpeg", "resource_url": "https://api.discogs.com/releases/13358737", "title": "Forgotten World 1", "id": 13358737, "year": 2019, "artist": "Golden Baby", "stats": {"community": {"in_wantlist": 616, "in_collection": 1051}, "user": {"in_wantlist": 0, "in_collection": 0}}}, {"status": "Accepted", "format": "12\"", "catno": "fw2", "thumb": "https://i.discogs.com/aoS-Gcc7iYEWaIK3WXrGGWS_bfpOuKiD106B9ZCPh_U/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEzMzU4/ODIxLTE1NTI3MzUy/MTAtNzAyNS5qcGVn.jpeg", "resource_url": "https://api.discogs.com/releases/13358821", "title": "Forgotten World 2", "id": 13358821, "year": 2019, "artist": "Golden Baby", "stats": {"community": {"in_wantlist": 527, "in_collection": 1046}, "user": {"in_wantlist": 0, "in_collection": 0}}}]}`))
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	releases, err := client.GetLabelReleases(context.Background(), 1611434, nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &LabelReleases{
		Pagination: Pagination{
			Page:    1,
			Pages:   1,
			PerPage: 50,
			Items:   2,
			URLs:    PaginationURLs{},
		},
		Releases: []LabelRelease{
			{
				Status:      "Accepted",
				Format:      "12\"",
				CatNo:       "fw1",
				Thumb:       "https://i.discogs.com/8FsOyncx8agIs2Kq2xvIL0orEwWdyDeWEpc9Rh5hNC8/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEzMzU4/NzM3LTE1NTI3MzUw/NjctNzUwMy5qcGVn.jpeg",
				ResourceURL: "https://api.discogs.com/releases/13358737",
				Title:       "Forgotten World 1",
				ID:          13358737,
				Year:        2019,
				Artist:      "Golden Baby",
				Stats: Stats{
					Community: CommunityStats{
						InWantlist:   616,
						InCollection: 1051,
					},
					User: UserStats{
						InWantlist:   0,
						InCollection: 0,
					},
				},
			},
			{
				Status:      "Accepted",
				Format:      "12\"",
				CatNo:       "fw2",
				Thumb:       "https://i.discogs.com/aoS-Gcc7iYEWaIK3WXrGGWS_bfpOuKiD106B9ZCPh_U/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEzMzU4/ODIxLTE1NTI3MzUy/MTAtNzAyNS5qcGVn.jpeg",
				ResourceURL: "https://api.discogs.com/releases/13358821",
				Title:       "Forgotten World 2",
				ID:          13358821,
				Year:        2019,
				Artist:      "Golden Baby",
				Stats: Stats{
					Community: CommunityStats{
						InWantlist:   527,
						InCollection: 1046,
					},
					User: UserStats{
						InWantlist:   0,
						InCollection: 0,
					},
				},
			},
		},
	}

	testEqual(t, want, releases)
}

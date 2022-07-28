package discogs

import (
	"context"
	"net/http"
	"testing"

	"github.com/peacefixation/go-discogs/discogs/auth"
)

func Test_GetRelease(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/releases/10507194", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		_, _ = w.Write([]byte(`{"id": 10507194, "status": "Accepted", "year": 2017, "resource_url": "https://api.discogs.com/releases/10507194", "uri": "https://www.discogs.com/release/10507194-HOLOVR-Likemind-05", "artists": [{"name": "HOLOVR", "anv": "", "join": "", "role": "", "tracks": "", "id": 3246671, "resource_url": "https://api.discogs.com/artists/3246671", "thumbnail_url": "https://i.discogs.com/N6d55ItJopwZsUkp0KTbmdEGhJ3Z5oNq-TT8fbiafdM/rs:fit/g:sm/q:40/h:196/w:196/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTMyNDY2/NzEtMTUwNDk2OTQw/OC04MjI2LmpwZWc.jpeg"}], "artists_sort": "HOLOVR", "labels": [{"name": "Likemind", "catno": "LM-05", "entity_type": "1", "entity_type_name": "Label", "id": 1225, "resource_url": "https://api.discogs.com/labels/1225", "thumbnail_url": "https://i.discogs.com/6Vu6M5HZajj_jwV0z35b4yI9e0bYBTeoqR9LV68nAxI/rs:fit/g:sm/q:40/h:198/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9MLTEyMjUt/MTQ3Mzc2OTU1My05/NzA5LmpwZWc.jpeg"}], "series": [], "companies": [], "formats": [{"name": "Vinyl", "qty": "1", "descriptions": ["12\"", "45 RPM"]}], "data_quality": "Needs Vote", "community": {"have": 270, "want": 464, "rating": {"count": 77, "average": 4.75}, "submitter": {"username": "adelaidesoundworks", "resource_url": "https://api.discogs.com/users/adelaidesoundworks"}, "contributors": [{"username": "adelaidesoundworks", "resource_url": "https://api.discogs.com/users/adelaidesoundworks"}, {"username": "section9", "resource_url": "https://api.discogs.com/users/section9"}, {"username": "tom", "resource_url": "https://api.discogs.com/users/tom"}, {"username": "simione001", "resource_url": "https://api.discogs.com/users/simione001"}], "data_quality": "Needs Vote", "status": "Accepted"}, "format_quantity": 1, "date_added": "2017-06-30T16:30:35-07:00", "date_changed": "2017-07-14T03:31:40-07:00", "num_for_sale": 5, "lowest_price": 12.74, "title": "Likemind 05", "country": "UK", "released": "2017-07-00", "released_formatted": "Jul 2017", "identifiers": [], "videos": [{"uri": "https://www.youtube.com/watch?v=GlfaOqBH_NU", "title": "HOLOVR - Split Infinity", "description": "Likemind release from 2017 \n\r\n\r-Video Upload powered by https://www.TunesToTube.com", "duration": 481, "embed": true}, {"uri": "https://www.youtube.com/watch?v=RR04HQBW43w", "title": "HOLOVR - Hillraiser", "description": "\u30c4 HOLOVR - Likemind 05, Likemind (2017)\n\nFollow me: http://www.facebook.com/itsmerubi\n\nBuy: https://www.discogs.com/HOLOVR-Likemind-05/release/10507194", "duration": 458, "embed": true}], "genres": ["Electronic"], "styles": ["Techno", "Abstract"], "tracklist": [{"position": "A", "type_": "track", "title": "Split Infinity", "duration": ""}, {"position": "B", "type_": "track", "title": "Hillraiser", "duration": ""}], "extraartists": [], "images": [{"type": "primary", "uri": "https://i.discogs.com/DCEl4HshBJ2KdeRSlXzus9fuyN1rB7lXWIWjEeJZ0t8/rs:fit/g:sm/q:90/h:602/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEwNTA3/MTk0LTE1MDAwMjgy/MjMtOTY4Mi5qcGVn.jpeg", "resource_url": "https://i.discogs.com/DCEl4HshBJ2KdeRSlXzus9fuyN1rB7lXWIWjEeJZ0t8/rs:fit/g:sm/q:90/h:602/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEwNTA3/MTk0LTE1MDAwMjgy/MjMtOTY4Mi5qcGVn.jpeg", "uri150": "https://i.discogs.com/gbMKS5xmgodMYU-ljHobR_JnexPGjqvwb4qbvtjgMO4/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEwNTA3/MTk0LTE1MDAwMjgy/MjMtOTY4Mi5qcGVn.jpeg", "width": 600, "height": 602}], "thumb": "https://i.discogs.com/gbMKS5xmgodMYU-ljHobR_JnexPGjqvwb4qbvtjgMO4/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEwNTA3/MTk0LTE1MDAwMjgy/MjMtOTY4Mi5qcGVn.jpeg", "estimated_weight": 230, "blocked_from_sale": false}`))
	})

	client := NewClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	release, err := client.GetRelease(context.Background(), 10507194, nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &Release{
		ID:          10507194,
		Status:      "Accepted",
		Year:        2017,
		ResourceURL: "https://api.discogs.com/releases/10507194",
		URI:         "https://www.discogs.com/release/10507194-HOLOVR-Likemind-05",
		Artists: []ReleaseArtist{
			{
				Name:         "HOLOVR",
				ANV:          "",
				Join:         "",
				Role:         "",
				Tracks:       "",
				ID:           3246671,
				ResourceURL:  "https://api.discogs.com/artists/3246671",
				ThumbnailURL: "https://i.discogs.com/N6d55ItJopwZsUkp0KTbmdEGhJ3Z5oNq-TT8fbiafdM/rs:fit/g:sm/q:40/h:196/w:196/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTMyNDY2/NzEtMTUwNDk2OTQw/OC04MjI2LmpwZWc.jpeg",
			},
		},
		ArtistsSort: "HOLOVR",
		Labels: []ReleaseLabel{
			{
				Name:           "Likemind",
				CatNo:          "LM-05",
				EntityType:     "1",
				EntityTypeName: "Label",
				ID:             1225,
				ResourceURL:    "https://api.discogs.com/labels/1225",
				ThumbnailURL:   "https://i.discogs.com/6Vu6M5HZajj_jwV0z35b4yI9e0bYBTeoqR9LV68nAxI/rs:fit/g:sm/q:40/h:198/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9MLTEyMjUt/MTQ3Mzc2OTU1My05/NzA5LmpwZWc.jpeg",
			},
		},
		Series:    []string{},
		Companies: []Company{},
		Formats: []Format{
			{
				Name:     "Vinyl",
				Quantity: "1",
				Descriptions: []string{
					"12\"",
					"45 RPM",
				},
			},
		},
		DataQuality: "Needs Vote",
		Community: Community{
			Have: 270,
			Want: 464,
			Rating: Rating{
				Count:   77,
				Average: 4.75,
			},
			Submitter: Submitter{
				Username:    "adelaidesoundworks",
				ResourceURL: "https://api.discogs.com/users/adelaidesoundworks",
			},
			Contributors: []Contributor{
				{
					Username:    "adelaidesoundworks",
					ResourceURL: "https://api.discogs.com/users/adelaidesoundworks",
				},
				{
					Username:    "section9",
					ResourceURL: "https://api.discogs.com/users/section9",
				},
				{
					Username:    "tom",
					ResourceURL: "https://api.discogs.com/users/tom",
				},
				{
					Username:    "simione001",
					ResourceURL: "https://api.discogs.com/users/simione001",
				},
			},
			DataQuality: "Needs Vote",
			Status:      "Accepted",
		},
		FormatQuantity:    1,
		DateAdded:         "2017-06-30T16:30:35-07:00",
		DateChanged:       "2017-07-14T03:31:40-07:00",
		NumForSale:        5,
		LowestPrice:       12.74,
		Title:             "Likemind 05",
		Country:           "UK",
		Released:          "2017-07-00",
		ReleasedFormatted: "Jul 2017",
		Identifiers:       []Identifier{},
		Videos: []Video{
			{
				URI:         "https://www.youtube.com/watch?v=GlfaOqBH_NU",
				Title:       "HOLOVR - Split Infinity",
				Description: "Likemind release from 2017 \n\r\n\r-Video Upload powered by https://www.TunesToTube.com",
				Duration:    481,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=RR04HQBW43w",
				Title:       "HOLOVR - Hillraiser",
				Description: "\u30c4 HOLOVR - Likemind 05, Likemind (2017)\n\nFollow me: http://www.facebook.com/itsmerubi\n\nBuy: https://www.discogs.com/HOLOVR-Likemind-05/release/10507194",
				Duration:    458,
				Embed:       true,
			},
		},
		Genres: []string{
			"Electronic",
		},
		Styles: []string{
			"Techno",
			"Abstract",
		},
		Tracklist: []Track{
			{
				Position: "A",
				Type:     "track",
				Title:    "Split Infinity",
				Duration: "",
			},
			{
				Position: "B",
				Type:     "track",
				Title:    "Hillraiser",
				Duration: "",
			},
		},
		ExtraArtists: []ReleaseArtist{},
		Images: []Image{
			{
				Type:        "primary",
				URI:         "https://i.discogs.com/DCEl4HshBJ2KdeRSlXzus9fuyN1rB7lXWIWjEeJZ0t8/rs:fit/g:sm/q:90/h:602/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEwNTA3/MTk0LTE1MDAwMjgy/MjMtOTY4Mi5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/DCEl4HshBJ2KdeRSlXzus9fuyN1rB7lXWIWjEeJZ0t8/rs:fit/g:sm/q:90/h:602/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEwNTA3/MTk0LTE1MDAwMjgy/MjMtOTY4Mi5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/gbMKS5xmgodMYU-ljHobR_JnexPGjqvwb4qbvtjgMO4/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEwNTA3/MTk0LTE1MDAwMjgy/MjMtOTY4Mi5qcGVn.jpeg",
				Width:       600,
				Height:      602,
			},
		},
		Thumb:           "https://i.discogs.com/gbMKS5xmgodMYU-ljHobR_JnexPGjqvwb4qbvtjgMO4/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEwNTA3/MTk0LTE1MDAwMjgy/MjMtOTY4Mi5qcGVn.jpeg",
		EstimatedWeight: 230,
		BlockedFromSale: false,
	}

	testEqual(t, want, release)
}

func Test_GetReleaseRatingByUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/releases/10507194/rating/Username", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		_, _ = w.Write([]byte(`{"release_id": 10507194, "username": "Username", "rating": 0}`))
	})

	client := NewClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	rating, err := client.GetReleaseRatingByUser(context.Background(), 10507194, "Username")
	if err != nil {
		t.Fatal(err)
	}

	want := &UserReleaseRating{
		ReleaseID: 10507194,
		Rating:    0,
		Username:  "Username",
	}

	testEqual(t, want, rating)
}

func Test_UpdateReleaseRating(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/releases/10507194/rating/Username", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"release_id": 10507194, "rating": 5, "username": "Username"}`))
	})

	client := NewClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	rating, err := client.UpdateReleaseRating(context.Background(), 10507194, "Username", 5)
	if err != nil {
		t.Fatal(err)
	}

	want := &UserReleaseRating{
		ReleaseID: 10507194,
		Rating:    5,
		Username:  "Username",
	}

	testEqual(t, want, rating)
}

func Test_DeleteReleaseRating(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/releases/10507194/rating/Username", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)

		w.WriteHeader(http.StatusNoContent)
	})

	client := NewClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	err := client.DeleteReleaseRating(context.Background(), 10507194, "Username")
	if err != nil {
		t.Fatal(err)
	}
}

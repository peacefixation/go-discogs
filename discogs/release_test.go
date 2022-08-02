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
		Series:    []ReleaseSeries{},
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

func Test_GetReleaseWithSeries(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/releases/71989", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		_, _ = w.Write([]byte(`{"id": 71989, "status": "Accepted", "year": 2002, "resource_url": "https://api.discogs.com/releases/71989", "uri": "https://www.discogs.com/release/71989-James-Lavelle-Barcelona-023", "artists": [{"name": "James Lavelle", "anv": "", "join": "", "role": "", "tracks": "", "id": 21209, "resource_url": "https://api.discogs.com/artists/21209", "thumbnail_url": "https://i.discogs.com/0Xi9E8euEnT7Tc8MZ5bDxqykuOw8cnAjjYULk07eQGg/rs:fit/g:sm/q:40/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTIxMjA5/LTE1NDM4MzA2MTYt/NTg3OC5qcGVn.jpeg"}], "artists_sort": "James Lavelle", "labels": [{"name": "Global Underground (3)", "catno": "GU023VIN", "entity_type": "1", "entity_type_name": "Label", "id": 1120989, "resource_url": "https://api.discogs.com/labels/1120989", "thumbnail_url": "https://i.discogs.com/yNFVArtsLhBj_noaBklZJ-eqMYOuM3han95-Q4WWUT0/rs:fit/g:sm/q:40/h:230/w:384/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9MLTExMjA5/ODktMTQ4MTg5OTk5/MS05MDgwLmpwZWc.jpeg"}], "series": [{"name": "Global Underground Series", "catno": "023", "entity_type": "2", "entity_type_name": "Series", "id": 333447, "resource_url": "https://api.discogs.com/labels/333447", "thumbnail_url": "https://i.discogs.com/TQh6_-BquOjm90Piyxo1iBfvxgryOPSd-SdhmWWyRIg/rs:fit/g:sm/q:40/h:336/w:357/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9MLTMzMzQ0/Ny0xMzMyMTgzMTk3/LmpwZWc.jpeg"}], "companies": [], "formats": [{"name": "Vinyl", "qty": "3", "descriptions": ["12\"", "Compilation", "Limited Edition"]}], "data_quality": "Correct", "community": {"have": 437, "want": 243, "rating": {"count": 60, "average": 4.4}, "submitter": {"username": "ck1", "resource_url": "https://api.discogs.com/users/ck1"}, "contributors": [{"username": "ck1", "resource_url": "https://api.discogs.com/users/ck1"}, {"username": "mitsou", "resource_url": "https://api.discogs.com/users/mitsou"}, {"username": "delysid", "resource_url": "https://api.discogs.com/users/delysid"}, {"username": "DiscogsUpdateBot", "resource_url": "https://api.discogs.com/users/DiscogsUpdateBot"}, {"username": "digirami", "resource_url": "https://api.discogs.com/users/digirami"}, {"username": "maxxyme", "resource_url": "https://api.discogs.com/users/maxxyme"}, {"username": "bobdobolina", "resource_url": "https://api.discogs.com/users/bobdobolina"}, {"username": "zevulon", "resource_url": "https://api.discogs.com/users/zevulon"}, {"username": "dancemania", "resource_url": "https://api.discogs.com/users/dancemania"}, {"username": "bapemania", "resource_url": "https://api.discogs.com/users/bapemania"}, {"username": "FromLondon", "resource_url": "https://api.discogs.com/users/FromLondon"}, {"username": "Plastic-Demons", "resource_url": "https://api.discogs.com/users/Plastic-Demons"}], "data_quality": "Correct", "status": "Accepted"}, "format_quantity": 3, "date_added": "2002-11-17T14:55:03-08:00", "date_changed": "2021-04-29T12:19:04-07:00", "num_for_sale": 12, "lowest_price": 7.42, "master_id": 49104, "master_url": "https://api.discogs.com/masters/49104", "title": "Barcelona #023", "country": "UK", "released": "2002-09-30", "notes": "Printed inner sleeves.\r\n\r\nTracks are presented as \"Side #1\" to \"Side #6\" instead of A to F.\r\n", "released_formatted": "30 Sep 2002", "identifiers": [{"type": "Barcode", "value": "664612202353"}], "videos": [{"uri": "https://www.youtube.com/watch?v=vLgvFpsgEDQ", "title": "Global Underground 023: Barcelona (CD1)", "description": "Label: Global Underground\nhttps://www.discogs.com/release/56133\nMixed by James Lavelle.\n\n(0:00:00) UNKLESOUNDS - 1-3 AM Intro Part One\n(0:01:50) DJ Shadow - Mongrell...Meets His Maker\n(0:06:06) Leftfield feat. Roots Manuva - Dusted\n(0:10:25) Thomas Bangal", "duration": 4149, "embed": true}, {"uri": "https://www.youtube.com/watch?v=w-oLZu96DkU", "title": "Sotero - The Vibe", "description": "One of my hard to find vinyls.. This track is taken from the global underground sampler vinyl of James Lavelle in Barcelona.", "duration": 333, "embed": true}, {"uri": "https://www.youtube.com/watch?v=mlg4mm7liWo", "title": "Global Underground 023: Barcelona (CD2)", "description": "Label: Global Underground\nhttps://www.discogs.com/release/56133\nMixed by James Lavelle.\n\n(0:00:00) UNKLESOUNDS - 3-5 AM Intro Part Two\n(0:01:34) Pitch Black/Eddie Amador - Underground Sound / The Funk (King Unique's Speech-apella)\n(0:06:54) Kybosh - Revol", "duration": 4379, "embed": true}, {"uri": "https://www.youtube.com/watch?v=Ek9xsW4GDGU", "title": "Layo & Bushwacka! - Let the Good Times Roll", "description": "Great tune from their Nightworks album.", "duration": 520, "embed": true}, {"uri": "https://www.youtube.com/watch?v=zWWim3Qv9UM", "title": "Force Mass Motion VS Dylan Rhymes - Hold back", "description": " ", "duration": 443, "embed": true}, {"uri": "https://www.youtube.com/watch?v=_DrOMNynGZ0", "title": "Meat Katie - Next Life", "description": "", "duration": 412, "embed": true}, {"uri": "https://www.youtube.com/watch?v=ZWaoQUifWKo", "title": "Dj Shadow - Mongrel Meets His Maker", "description": "De l'album The Private Press sorti en 2002. Musique trip-hop. Sur l'album, les pistes 7 et 8 ont le m\u00eame titre Mongrel Meets His Maker, je les ai regroup\u00e9es.", "duration": 317, "embed": true}, {"uri": "https://www.youtube.com/watch?v=llXpjsYen7w", "title": "doves- there goes the fear (james lavelle remix closing GU)", "description": "Global underground closing BCN", "duration": 355, "embed": true}], "genres": ["Electronic"], "styles": ["Breakbeat", "Progressive House"], "tracklist": [{"position": "A", "type_": "track", "artists": [{"name": "Sotero", "anv": "", "join": "", "role": "", "tracks": "", "id": 22262, "resource_url": "https://api.discogs.com/artists/22262", "thumbnail_url": ""}], "title": "The Vibe", "duration": ""}, {"position": "B", "type_": "track", "artists": [{"name": "Layo & Bushwacka!", "anv": "", "join": "", "role": "", "tracks": "", "id": 90, "resource_url": "https://api.discogs.com/artists/90", "thumbnail_url": "https://i.discogs.com/WonJbrQq3ZpRkQNyo3NXJatbriLWy-vk3Ns03YKLEfU/rs:fit/g:sm/q:40/h:262/w:350/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTkwLTE1/OTYwNTg5NDctNzM4/NC5qcGVn.jpeg"}], "title": "Let The Good Times Roll", "duration": ""}, {"position": "C", "type_": "track", "artists": [{"name": "Force Mass Motion", "anv": "", "join": "vs", "role": "", "tracks": "", "id": 5550, "resource_url": "https://api.discogs.com/artists/5550", "thumbnail_url": "https://i.discogs.com/6RJd9LDCuK2pzud5YBdCnb0UNHyRasOhRB8sSszafds/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTU1NTAt/MTE0OTQ1ODE5NC5q/cGVn.jpeg"}, {"name": "Dylan Rhymes", "anv": "", "join": "", "role": "", "tracks": "", "id": 6521, "resource_url": "https://api.discogs.com/artists/6521", "thumbnail_url": "https://i.discogs.com/221VVg8sn41cVaCgVt5yuzBsZOWC5QsXsBY8M2JaApY/rs:fit/g:sm/q:40/h:399/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTY1MjEt/MTQ0NzUzMzQwOC03/MjM4LmpwZWc.jpeg"}], "title": "Hold Back", "duration": ""}, {"position": "D", "type_": "track", "artists": [{"name": "Meat Katie", "anv": "", "join": "", "role": "", "tracks": "", "id": 3172, "resource_url": "https://api.discogs.com/artists/3172", "thumbnail_url": "https://i.discogs.com/NvUosz1UpsrPK5f2hFtYrEi5hnTe9kzaFMuCbrASgy8/rs:fit/g:sm/q:40/h:632/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTMxNzIt/MTQ0NzUxNTM4Mi0y/MjA1LmpwZWc.jpeg"}], "title": "Next Life", "duration": ""}, {"position": "E", "type_": "track", "artists": [{"name": "DJ Shadow", "anv": "", "join": "", "role": "", "tracks": "", "id": 4478, "resource_url": "https://api.discogs.com/artists/4478", "thumbnail_url": "https://i.discogs.com/2ftedFe7WP_gmGoP-edeOkKcgxDljzXm_HkhkMJevC0/rs:fit/g:sm/q:40/h:455/w:367/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTQ0Nzgt/MTU4NDUzMTkxNy04/MjMwLmpwZWc.jpeg"}], "title": "Mongrel\u2026Meets His Maker", "duration": ""}, {"position": "F", "type_": "track", "artists": [{"name": "Doves", "anv": "", "join": "", "role": "", "tracks": "", "id": 3855, "resource_url": "https://api.discogs.com/artists/3855", "thumbnail_url": "https://i.discogs.com/lrr2e5dsootQfcG0_V79PQi5Q_QVKNx5BZ3QAg03cjc/rs:fit/g:sm/q:40/h:381/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTM4NTUt/MTU5NDMxNzQ5NC03/NDAxLmpwZWc.jpeg"}], "title": "There Goes The Fear (Unklesounds GU Fear-edit Part Two)", "extraartists": [{"name": "UNKLE Sounds", "anv": "Unklesounds", "join": "", "role": "Remix", "tracks": "", "id": 54120, "resource_url": "https://api.discogs.com/artists/54120", "thumbnail_url": "https://i.discogs.com/BJuDuuvUcd0kbuFWReRUuYoFwVmZ1d6pcBT48HXBgHE/rs:fit/g:sm/q:40/h:278/w:425/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTU0MTIw/LTEyNDMwODcwNDYu/anBlZw.jpeg"}], "duration": ""}], "extraartists": [{"name": "James Lavelle", "anv": "", "join": "", "role": "Compiled By", "tracks": "", "id": 21209, "resource_url": "https://api.discogs.com/artists/21209", "thumbnail_url": "https://i.discogs.com/0Xi9E8euEnT7Tc8MZ5bDxqykuOw8cnAjjYULk07eQGg/rs:fit/g:sm/q:40/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTIxMjA5/LTE1NDM4MzA2MTYt/NTg3OC5qcGVn.jpeg"}], "images": [{"type": "primary", "uri": "https://i.discogs.com/aPsUSgP63K5Pafzuwy-Gy-VLZ5No2SKG5Lrdnlo2i0o/rs:fit/g:sm/q:90/h:595/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM3ODQt/Mjk2MS5qcGVn.jpeg", "resource_url": "https://i.discogs.com/aPsUSgP63K5Pafzuwy-Gy-VLZ5No2SKG5Lrdnlo2i0o/rs:fit/g:sm/q:90/h:595/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM3ODQt/Mjk2MS5qcGVn.jpeg", "uri150": "https://i.discogs.com/UGFJFwk2IUyVgmh_BItSab76v2MhipF7YQyMPDstQzs/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM3ODQt/Mjk2MS5qcGVn.jpeg", "width": 600, "height": 595}, {"type": "secondary", "uri": "https://i.discogs.com/t5wm5lQH14ZgMNY1xN6xNU9rUu81PqSF_bOnJH9hSik/rs:fit/g:sm/q:90/h:594/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4Njgt/ODk1NS5qcGVn.jpeg", "resource_url": "https://i.discogs.com/t5wm5lQH14ZgMNY1xN6xNU9rUu81PqSF_bOnJH9hSik/rs:fit/g:sm/q:90/h:594/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4Njgt/ODk1NS5qcGVn.jpeg", "uri150": "https://i.discogs.com/Tjf2OY_Zo2Iff_eKPT6kEYgD8V2l621kDwUQQ2LcaoM/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4Njgt/ODk1NS5qcGVn.jpeg", "width": 600, "height": 594}, {"type": "secondary", "uri": "https://i.discogs.com/D23Yt-IO1x8QHK-eJidjE42mK4lTcSuG3cW4mXef0rk/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4NzMt/NDEwOC5qcGVn.jpeg", "resource_url": "https://i.discogs.com/D23Yt-IO1x8QHK-eJidjE42mK4lTcSuG3cW4mXef0rk/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4NzMt/NDEwOC5qcGVn.jpeg", "uri150": "https://i.discogs.com/MsRWbAj8xhYq5evhorjPb2MYZRZ1_H4SkcvScMp_Upc/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4NzMt/NDEwOC5qcGVn.jpeg", "width": 600, "height": 600}, {"type": "secondary", "uri": "https://i.discogs.com/schKpcxfO2NJ6wejY0lQJD2HZi-jBaDxsb4Uriis40s/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4Nzkt/MTg0Mi5qcGVn.jpeg", "resource_url": "https://i.discogs.com/schKpcxfO2NJ6wejY0lQJD2HZi-jBaDxsb4Uriis40s/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4Nzkt/MTg0Mi5qcGVn.jpeg", "uri150": "https://i.discogs.com/j-yNpXTjTePZjDEF8MrnKk6EzoECz-DphsqWMRWY90E/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4Nzkt/MTg0Mi5qcGVn.jpeg", "width": 600, "height": 600}, {"type": "secondary", "uri": "https://i.discogs.com/3rMLclHQ--tUXrWs8ND10ePmmRBz1fIhKF7HionQjoQ/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4ODMt/NDE2Ni5qcGVn.jpeg", "resource_url": "https://i.discogs.com/3rMLclHQ--tUXrWs8ND10ePmmRBz1fIhKF7HionQjoQ/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4ODMt/NDE2Ni5qcGVn.jpeg", "uri150": "https://i.discogs.com/AuOgjD4xysNYo1yTXLtpPFVxoUM23ogR9iQgnUNim3c/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4ODMt/NDE2Ni5qcGVn.jpeg", "width": 600, "height": 600}, {"type": "secondary", "uri": "https://i.discogs.com/fML1mt5GQVQ1EzPSDg8I0q-p8q48eAF3Qa7-O-T_Nuo/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4ODkt/ODg3Ny5qcGVn.jpeg", "resource_url": "https://i.discogs.com/fML1mt5GQVQ1EzPSDg8I0q-p8q48eAF3Qa7-O-T_Nuo/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4ODkt/ODg3Ny5qcGVn.jpeg", "uri150": "https://i.discogs.com/Gfl0enLp5FwW_jRsMH0JxhZ_JSUo6J-mj-1-z-B47To/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4ODkt/ODg3Ny5qcGVn.jpeg", "width": 600, "height": 600}, {"type": "secondary", "uri": "https://i.discogs.com/vqp25TTgBAzjHykLb8uGStZBSVC2ypqFY8beZOvWgow/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4OTQt/NjMyNy5qcGVn.jpeg", "resource_url": "https://i.discogs.com/vqp25TTgBAzjHykLb8uGStZBSVC2ypqFY8beZOvWgow/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4OTQt/NjMyNy5qcGVn.jpeg", "uri150": "https://i.discogs.com/5fqiTrb-yjA7bjKqbf6F6AEi8cFiBvRnL4FcqLo3VvE/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4OTQt/NjMyNy5qcGVn.jpeg", "width": 600, "height": 600}, {"type": "secondary", "uri": "https://i.discogs.com/ZX9r4EwRY4Kdto1UjeQsyFJAabSB7u182TW5bUYkauY/rs:fit/g:sm/q:90/h:604/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MDAt/ODI2MS5qcGVn.jpeg", "resource_url": "https://i.discogs.com/ZX9r4EwRY4Kdto1UjeQsyFJAabSB7u182TW5bUYkauY/rs:fit/g:sm/q:90/h:604/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MDAt/ODI2MS5qcGVn.jpeg", "uri150": "https://i.discogs.com/C_aYfVBYTTxUXhubL9aU2w7iVzbTAC8CxO0HG5wBRBo/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MDAt/ODI2MS5qcGVn.jpeg", "width": 600, "height": 604}, {"type": "secondary", "uri": "https://i.discogs.com/GY9nK7h0gMQGfm9-LMa1My1BJPfBOLsGyVlWlsDdqzk/rs:fit/g:sm/q:90/h:596/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MDYt/NDgzNi5qcGVn.jpeg", "resource_url": "https://i.discogs.com/GY9nK7h0gMQGfm9-LMa1My1BJPfBOLsGyVlWlsDdqzk/rs:fit/g:sm/q:90/h:596/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MDYt/NDgzNi5qcGVn.jpeg", "uri150": "https://i.discogs.com/YpEX7vkMAA5XYWezdW0dgY6rmxrEtWHEKOozHvLk5s4/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MDYt/NDgzNi5qcGVn.jpeg", "width": 600, "height": 596}, {"type": "secondary", "uri": "https://i.discogs.com/VijMPtyt2gGnFzoB-XQKhA0mUf6fJMMzxnufiCK8Iqg/rs:fit/g:sm/q:90/h:597/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MTMt/Njc1Ni5qcGVn.jpeg", "resource_url": "https://i.discogs.com/VijMPtyt2gGnFzoB-XQKhA0mUf6fJMMzxnufiCK8Iqg/rs:fit/g:sm/q:90/h:597/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MTMt/Njc1Ni5qcGVn.jpeg", "uri150": "https://i.discogs.com/Zqb_kA6YVQdnr-D4RzDPKW1ISkC0C4EJX2V8ai_YmGc/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MTMt/Njc1Ni5qcGVn.jpeg", "width": 600, "height": 597}, {"type": "secondary", "uri": "https://i.discogs.com/6npIbSdLgzMbsIu5JVc-TQ9hzrj83GWryvRsrQForYM/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MjEt/MzI0My5qcGVn.jpeg", "resource_url": "https://i.discogs.com/6npIbSdLgzMbsIu5JVc-TQ9hzrj83GWryvRsrQForYM/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MjEt/MzI0My5qcGVn.jpeg", "uri150": "https://i.discogs.com/KTghamB3k7UOMUG6uK6ega4JQZAT8s6L-a0UsyoDZJo/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MjEt/MzI0My5qcGVn.jpeg", "width": 600, "height": 600}, {"type": "secondary", "uri": "https://i.discogs.com/-4Nl0VieXfCQVZhv3sYVA0OHMwrLjsQ22Rlk4XpVyMU/rs:fit/g:sm/q:90/h:597/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MjUt/OTQ2Ny5qcGVn.jpeg", "resource_url": "https://i.discogs.com/-4Nl0VieXfCQVZhv3sYVA0OHMwrLjsQ22Rlk4XpVyMU/rs:fit/g:sm/q:90/h:597/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MjUt/OTQ2Ny5qcGVn.jpeg", "uri150": "https://i.discogs.com/eiXksLzk6Z6tH4jrUJj80GXi8GB2cM0sDp4_r72xLGg/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MjUt/OTQ2Ny5qcGVn.jpeg", "width": 600, "height": 597}, {"type": "secondary", "uri": "https://i.discogs.com/g0yjX2zA8xg75Wrwgvt0zaOZhReCMOpZ_kwY36yjDAc/rs:fit/g:sm/q:90/h:598/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MzMt/OTY4MS5qcGVn.jpeg", "resource_url": "https://i.discogs.com/g0yjX2zA8xg75Wrwgvt0zaOZhReCMOpZ_kwY36yjDAc/rs:fit/g:sm/q:90/h:598/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MzMt/OTY4MS5qcGVn.jpeg", "uri150": "https://i.discogs.com/sL2nJRop2mwsnR6Pi-PP6i64woU6hqA7dhiGO1kprxo/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MzMt/OTY4MS5qcGVn.jpeg", "width": 600, "height": 598}, {"type": "secondary", "uri": "https://i.discogs.com/o1j8_MjdliyAcMtlwGD8tipW0G515qlZZwT-8N2HCj8/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5Mzgt/Mjk4MC5qcGVn.jpeg", "resource_url": "https://i.discogs.com/o1j8_MjdliyAcMtlwGD8tipW0G515qlZZwT-8N2HCj8/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5Mzgt/Mjk4MC5qcGVn.jpeg", "uri150": "https://i.discogs.com/qFbEMMuUy8DhdbnSkJGw_hrUGsfhL8yeuqgxYAtXr3Y/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5Mzgt/Mjk4MC5qcGVn.jpeg", "width": 600, "height": 600}], "thumb": "https://i.discogs.com/UGFJFwk2IUyVgmh_BItSab76v2MhipF7YQyMPDstQzs/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM3ODQt/Mjk2MS5qcGVn.jpeg", "estimated_weight": 690, "blocked_from_sale": false}`))
	})

	client := NewClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	release, err := client.GetRelease(context.Background(), 71989, nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &Release{
		ID:          71989,
		Status:      "Accepted",
		Year:        2002,
		ResourceURL: "https://api.discogs.com/releases/71989",
		URI:         "https://www.discogs.com/release/71989-James-Lavelle-Barcelona-023",
		Artists: []ReleaseArtist{
			{
				Name:         "James Lavelle",
				ANV:          "",
				Join:         "",
				Role:         "",
				Tracks:       "",
				ID:           21209,
				ResourceURL:  "https://api.discogs.com/artists/21209",
				ThumbnailURL: "https://i.discogs.com/0Xi9E8euEnT7Tc8MZ5bDxqykuOw8cnAjjYULk07eQGg/rs:fit/g:sm/q:40/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTIxMjA5/LTE1NDM4MzA2MTYt/NTg3OC5qcGVn.jpeg",
			},
		},
		ArtistsSort: "James Lavelle",
		Labels: []ReleaseLabel{
			{
				Name:           "Global Underground (3)",
				CatNo:          "GU023VIN",
				EntityType:     "1",
				EntityTypeName: "Label",
				ID:             1120989,
				ResourceURL:    "https://api.discogs.com/labels/1120989",
				ThumbnailURL:   "https://i.discogs.com/yNFVArtsLhBj_noaBklZJ-eqMYOuM3han95-Q4WWUT0/rs:fit/g:sm/q:40/h:230/w:384/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9MLTExMjA5/ODktMTQ4MTg5OTk5/MS05MDgwLmpwZWc.jpeg",
			},
		},
		Series: []ReleaseSeries{
			{
				Name:           "Global Underground Series",
				CatNo:          "023",
				EntityType:     "2",
				EntityTypeName: "Series",
				ID:             333447,
				ResourceURL:    "https://api.discogs.com/labels/333447",
				ThumbnailURL:   "https://i.discogs.com/TQh6_-BquOjm90Piyxo1iBfvxgryOPSd-SdhmWWyRIg/rs:fit/g:sm/q:40/h:336/w:357/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9MLTMzMzQ0/Ny0xMzMyMTgzMTk3/LmpwZWc.jpeg",
			},
		},
		Companies: []Company{},
		Formats: []Format{
			{
				Name:     "Vinyl",
				Quantity: "3",
				Descriptions: []string{
					"12\"",
					"Compilation",
					"Limited Edition",
				},
			},
		},
		DataQuality: "Correct",
		Community: Community{
			Have: 437,
			Want: 243,
			Rating: Rating{
				Count:   60,
				Average: 4.4,
			},
			Submitter: Submitter{
				Username:    "ck1",
				ResourceURL: "https://api.discogs.com/users/ck1",
			},
			Contributors: []Contributor{
				{
					Username:    "ck1",
					ResourceURL: "https://api.discogs.com/users/ck1",
				},
				{
					Username:    "mitsou",
					ResourceURL: "https://api.discogs.com/users/mitsou",
				},
				{
					Username:    "delysid",
					ResourceURL: "https://api.discogs.com/users/delysid",
				},
				{
					Username:    "DiscogsUpdateBot",
					ResourceURL: "https://api.discogs.com/users/DiscogsUpdateBot",
				},
				{
					Username:    "digirami",
					ResourceURL: "https://api.discogs.com/users/digirami",
				},
				{
					Username:    "maxxyme",
					ResourceURL: "https://api.discogs.com/users/maxxyme",
				},
				{
					Username:    "bobdobolina",
					ResourceURL: "https://api.discogs.com/users/bobdobolina",
				},
				{
					Username:    "zevulon",
					ResourceURL: "https://api.discogs.com/users/zevulon",
				},
				{
					Username:    "dancemania",
					ResourceURL: "https://api.discogs.com/users/dancemania",
				},
				{
					Username:    "bapemania",
					ResourceURL: "https://api.discogs.com/users/bapemania",
				},
				{
					Username:    "FromLondon",
					ResourceURL: "https://api.discogs.com/users/FromLondon",
				},
				{
					Username:    "Plastic-Demons",
					ResourceURL: "https://api.discogs.com/users/Plastic-Demons",
				},
			},
			DataQuality: "Correct",
			Status:      "Accepted",
		},
		FormatQuantity:    3,
		DateAdded:         "2002-11-17T14:55:03-08:00",
		DateChanged:       "2021-04-29T12:19:04-07:00",
		NumForSale:        12,
		LowestPrice:       7.42,
		MasterID:          49104,
		MasterURL:         "https://api.discogs.com/masters/49104",
		Title:             "Barcelona #023",
		Country:           "UK",
		Released:          "2002-09-30",
		Notes:             "Printed inner sleeves.\r\n\r\nTracks are presented as \"Side #1\" to \"Side #6\" instead of A to F.\r\n",
		ReleasedFormatted: "30 Sep 2002",
		Identifiers: []Identifier{
			{
				Type:  "Barcode",
				Value: "664612202353",
			},
		},
		Videos: []Video{
			{
				URI:         "https://www.youtube.com/watch?v=vLgvFpsgEDQ",
				Title:       "Global Underground 023: Barcelona (CD1)",
				Description: "Label: Global Underground\nhttps://www.discogs.com/release/56133\nMixed by James Lavelle.\n\n(0:00:00) UNKLESOUNDS - 1-3 AM Intro Part One\n(0:01:50) DJ Shadow - Mongrell...Meets His Maker\n(0:06:06) Leftfield feat. Roots Manuva - Dusted\n(0:10:25) Thomas Bangal",
				Duration:    4149,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=w-oLZu96DkU",
				Title:       "Sotero - The Vibe",
				Description: "One of my hard to find vinyls.. This track is taken from the global underground sampler vinyl of James Lavelle in Barcelona.",
				Duration:    333,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=mlg4mm7liWo",
				Title:       "Global Underground 023: Barcelona (CD2)",
				Description: "Label: Global Underground\nhttps://www.discogs.com/release/56133\nMixed by James Lavelle.\n\n(0:00:00) UNKLESOUNDS - 3-5 AM Intro Part Two\n(0:01:34) Pitch Black/Eddie Amador - Underground Sound / The Funk (King Unique's Speech-apella)\n(0:06:54) Kybosh - Revol",
				Duration:    4379,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=Ek9xsW4GDGU",
				Title:       "Layo & Bushwacka! - Let the Good Times Roll",
				Description: "Great tune from their Nightworks album.",
				Duration:    520,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=zWWim3Qv9UM",
				Title:       "Force Mass Motion VS Dylan Rhymes - Hold back",
				Description: " ",
				Duration:    443,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=_DrOMNynGZ0",
				Title:       "Meat Katie - Next Life",
				Description: "",
				Duration:    412,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=ZWaoQUifWKo",
				Title:       "Dj Shadow - Mongrel Meets His Maker",
				Description: "De l'album The Private Press sorti en 2002. Musique trip-hop. Sur l'album, les pistes 7 et 8 ont le m\u00eame titre Mongrel Meets His Maker, je les ai regroup\u00e9es.",
				Duration:    317,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=llXpjsYen7w",
				Title:       "doves- there goes the fear (james lavelle remix closing GU)",
				Description: "Global underground closing BCN",
				Duration:    355,
				Embed:       true,
			},
		},
		Genres: []string{
			"Electronic",
		},
		Styles: []string{
			"Breakbeat",
			"Progressive House",
		},
		Tracklist: []Track{
			{
				Position: "A",
				Type:     "track",
				Artists: []ReleaseArtist{
					{
						Name:         "Sotero",
						ANV:          "",
						Join:         "",
						Role:         "",
						Tracks:       "",
						ID:           22262,
						ResourceURL:  "https://api.discogs.com/artists/22262",
						ThumbnailURL: "",
					},
				},
				Title:    "The Vibe",
				Duration: "",
			},
			{
				Position: "B",
				Type:     "track",
				Artists: []ReleaseArtist{
					{
						Name:         "Layo & Bushwacka!",
						ANV:          "",
						Join:         "",
						Role:         "",
						Tracks:       "",
						ID:           90,
						ResourceURL:  "https://api.discogs.com/artists/90",
						ThumbnailURL: "https://i.discogs.com/WonJbrQq3ZpRkQNyo3NXJatbriLWy-vk3Ns03YKLEfU/rs:fit/g:sm/q:40/h:262/w:350/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTkwLTE1/OTYwNTg5NDctNzM4/NC5qcGVn.jpeg",
					},
				},
				Title:    "Let The Good Times Roll",
				Duration: "",
			},
			{
				Position: "C",
				Type:     "track",
				Artists: []ReleaseArtist{
					{
						Name:         "Force Mass Motion",
						ANV:          "",
						Join:         "vs",
						Role:         "",
						Tracks:       "",
						ID:           5550,
						ResourceURL:  "https://api.discogs.com/artists/5550",
						ThumbnailURL: "https://i.discogs.com/6RJd9LDCuK2pzud5YBdCnb0UNHyRasOhRB8sSszafds/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTU1NTAt/MTE0OTQ1ODE5NC5q/cGVn.jpeg",
					},
					{
						Name:         "Dylan Rhymes",
						ANV:          "",
						Join:         "",
						Role:         "",
						Tracks:       "",
						ID:           6521,
						ResourceURL:  "https://api.discogs.com/artists/6521",
						ThumbnailURL: "https://i.discogs.com/221VVg8sn41cVaCgVt5yuzBsZOWC5QsXsBY8M2JaApY/rs:fit/g:sm/q:40/h:399/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTY1MjEt/MTQ0NzUzMzQwOC03/MjM4LmpwZWc.jpeg",
					},
				},
				Title:    "Hold Back",
				Duration: "",
			},
			{
				Position: "D",
				Type:     "track",
				Artists: []ReleaseArtist{
					{
						Name:         "Meat Katie",
						ANV:          "",
						Join:         "",
						Role:         "",
						Tracks:       "",
						ID:           3172,
						ResourceURL:  "https://api.discogs.com/artists/3172",
						ThumbnailURL: "https://i.discogs.com/NvUosz1UpsrPK5f2hFtYrEi5hnTe9kzaFMuCbrASgy8/rs:fit/g:sm/q:40/h:632/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTMxNzIt/MTQ0NzUxNTM4Mi0y/MjA1LmpwZWc.jpeg",
					},
				},
				Title:    "Next Life",
				Duration: "",
			},
			{
				Position: "E",
				Type:     "track",
				Artists: []ReleaseArtist{
					{
						Name:         "DJ Shadow",
						ANV:          "",
						Join:         "",
						Role:         "",
						Tracks:       "",
						ID:           4478,
						ResourceURL:  "https://api.discogs.com/artists/4478",
						ThumbnailURL: "https://i.discogs.com/2ftedFe7WP_gmGoP-edeOkKcgxDljzXm_HkhkMJevC0/rs:fit/g:sm/q:40/h:455/w:367/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTQ0Nzgt/MTU4NDUzMTkxNy04/MjMwLmpwZWc.jpeg",
					},
				},
				Title:    "Mongrel\u2026Meets His Maker",
				Duration: "",
			},
			{
				Position: "F",
				Type:     "track",
				Artists: []ReleaseArtist{
					{
						Name:         "Doves",
						ANV:          "",
						Join:         "",
						Role:         "",
						Tracks:       "",
						ID:           3855,
						ResourceURL:  "https://api.discogs.com/artists/3855",
						ThumbnailURL: "https://i.discogs.com/lrr2e5dsootQfcG0_V79PQi5Q_QVKNx5BZ3QAg03cjc/rs:fit/g:sm/q:40/h:381/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTM4NTUt/MTU5NDMxNzQ5NC03/NDAxLmpwZWc.jpeg",
					},
				},
				Title: "There Goes The Fear (Unklesounds GU Fear-edit Part Two)",
				ExtraArtists: []ReleaseArtist{
					{
						Name:         "UNKLE Sounds",
						ANV:          "Unklesounds",
						Join:         "",
						Role:         "Remix",
						Tracks:       "",
						ID:           54120,
						ResourceURL:  "https://api.discogs.com/artists/54120",
						ThumbnailURL: "https://i.discogs.com/BJuDuuvUcd0kbuFWReRUuYoFwVmZ1d6pcBT48HXBgHE/rs:fit/g:sm/q:40/h:278/w:425/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTU0MTIw/LTEyNDMwODcwNDYu/anBlZw.jpeg",
					},
				},
				Duration: "",
			},
		},
		ExtraArtists: []ReleaseArtist{
			{
				Name:         "James Lavelle",
				ANV:          "",
				Join:         "",
				Role:         "Compiled By",
				Tracks:       "",
				ID:           21209,
				ResourceURL:  "https://api.discogs.com/artists/21209",
				ThumbnailURL: "https://i.discogs.com/0Xi9E8euEnT7Tc8MZ5bDxqykuOw8cnAjjYULk07eQGg/rs:fit/g:sm/q:40/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTIxMjA5/LTE1NDM4MzA2MTYt/NTg3OC5qcGVn.jpeg",
			},
		},
		Images: []Image{
			{
				Type:        "primary",
				URI:         "https://i.discogs.com/aPsUSgP63K5Pafzuwy-Gy-VLZ5No2SKG5Lrdnlo2i0o/rs:fit/g:sm/q:90/h:595/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM3ODQt/Mjk2MS5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/aPsUSgP63K5Pafzuwy-Gy-VLZ5No2SKG5Lrdnlo2i0o/rs:fit/g:sm/q:90/h:595/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM3ODQt/Mjk2MS5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/UGFJFwk2IUyVgmh_BItSab76v2MhipF7YQyMPDstQzs/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM3ODQt/Mjk2MS5qcGVn.jpeg",
				Width:       600,
				Height:      595,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/t5wm5lQH14ZgMNY1xN6xNU9rUu81PqSF_bOnJH9hSik/rs:fit/g:sm/q:90/h:594/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4Njgt/ODk1NS5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/t5wm5lQH14ZgMNY1xN6xNU9rUu81PqSF_bOnJH9hSik/rs:fit/g:sm/q:90/h:594/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4Njgt/ODk1NS5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/Tjf2OY_Zo2Iff_eKPT6kEYgD8V2l621kDwUQQ2LcaoM/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4Njgt/ODk1NS5qcGVn.jpeg",
				Width:       600,
				Height:      594,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/D23Yt-IO1x8QHK-eJidjE42mK4lTcSuG3cW4mXef0rk/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4NzMt/NDEwOC5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/D23Yt-IO1x8QHK-eJidjE42mK4lTcSuG3cW4mXef0rk/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4NzMt/NDEwOC5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/MsRWbAj8xhYq5evhorjPb2MYZRZ1_H4SkcvScMp_Upc/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4NzMt/NDEwOC5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/schKpcxfO2NJ6wejY0lQJD2HZi-jBaDxsb4Uriis40s/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4Nzkt/MTg0Mi5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/schKpcxfO2NJ6wejY0lQJD2HZi-jBaDxsb4Uriis40s/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4Nzkt/MTg0Mi5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/j-yNpXTjTePZjDEF8MrnKk6EzoECz-DphsqWMRWY90E/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4Nzkt/MTg0Mi5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/3rMLclHQ--tUXrWs8ND10ePmmRBz1fIhKF7HionQjoQ/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4ODMt/NDE2Ni5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/3rMLclHQ--tUXrWs8ND10ePmmRBz1fIhKF7HionQjoQ/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4ODMt/NDE2Ni5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/AuOgjD4xysNYo1yTXLtpPFVxoUM23ogR9iQgnUNim3c/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4ODMt/NDE2Ni5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/fML1mt5GQVQ1EzPSDg8I0q-p8q48eAF3Qa7-O-T_Nuo/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4ODkt/ODg3Ny5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/fML1mt5GQVQ1EzPSDg8I0q-p8q48eAF3Qa7-O-T_Nuo/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4ODkt/ODg3Ny5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/Gfl0enLp5FwW_jRsMH0JxhZ_JSUo6J-mj-1-z-B47To/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4ODkt/ODg3Ny5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/vqp25TTgBAzjHykLb8uGStZBSVC2ypqFY8beZOvWgow/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4OTQt/NjMyNy5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/vqp25TTgBAzjHykLb8uGStZBSVC2ypqFY8beZOvWgow/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4OTQt/NjMyNy5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/5fqiTrb-yjA7bjKqbf6F6AEi8cFiBvRnL4FcqLo3VvE/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM4OTQt/NjMyNy5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/ZX9r4EwRY4Kdto1UjeQsyFJAabSB7u182TW5bUYkauY/rs:fit/g:sm/q:90/h:604/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MDAt/ODI2MS5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/ZX9r4EwRY4Kdto1UjeQsyFJAabSB7u182TW5bUYkauY/rs:fit/g:sm/q:90/h:604/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MDAt/ODI2MS5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/C_aYfVBYTTxUXhubL9aU2w7iVzbTAC8CxO0HG5wBRBo/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MDAt/ODI2MS5qcGVn.jpeg",
				Width:       600,
				Height:      604,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/GY9nK7h0gMQGfm9-LMa1My1BJPfBOLsGyVlWlsDdqzk/rs:fit/g:sm/q:90/h:596/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MDYt/NDgzNi5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/GY9nK7h0gMQGfm9-LMa1My1BJPfBOLsGyVlWlsDdqzk/rs:fit/g:sm/q:90/h:596/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MDYt/NDgzNi5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/YpEX7vkMAA5XYWezdW0dgY6rmxrEtWHEKOozHvLk5s4/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MDYt/NDgzNi5qcGVn.jpeg",
				Width:       600,
				Height:      596,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/VijMPtyt2gGnFzoB-XQKhA0mUf6fJMMzxnufiCK8Iqg/rs:fit/g:sm/q:90/h:597/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MTMt/Njc1Ni5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/VijMPtyt2gGnFzoB-XQKhA0mUf6fJMMzxnufiCK8Iqg/rs:fit/g:sm/q:90/h:597/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MTMt/Njc1Ni5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/Zqb_kA6YVQdnr-D4RzDPKW1ISkC0C4EJX2V8ai_YmGc/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MTMt/Njc1Ni5qcGVn.jpeg",
				Width:       600,
				Height:      597,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/6npIbSdLgzMbsIu5JVc-TQ9hzrj83GWryvRsrQForYM/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MjEt/MzI0My5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/6npIbSdLgzMbsIu5JVc-TQ9hzrj83GWryvRsrQForYM/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MjEt/MzI0My5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/KTghamB3k7UOMUG6uK6ega4JQZAT8s6L-a0UsyoDZJo/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MjEt/MzI0My5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/-4Nl0VieXfCQVZhv3sYVA0OHMwrLjsQ22Rlk4XpVyMU/rs:fit/g:sm/q:90/h:597/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MjUt/OTQ2Ny5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/-4Nl0VieXfCQVZhv3sYVA0OHMwrLjsQ22Rlk4XpVyMU/rs:fit/g:sm/q:90/h:597/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MjUt/OTQ2Ny5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/eiXksLzk6Z6tH4jrUJj80GXi8GB2cM0sDp4_r72xLGg/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MjUt/OTQ2Ny5qcGVn.jpeg",
				Width:       600,
				Height:      597,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/g0yjX2zA8xg75Wrwgvt0zaOZhReCMOpZ_kwY36yjDAc/rs:fit/g:sm/q:90/h:598/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MzMt/OTY4MS5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/g0yjX2zA8xg75Wrwgvt0zaOZhReCMOpZ_kwY36yjDAc/rs:fit/g:sm/q:90/h:598/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MzMt/OTY4MS5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/sL2nJRop2mwsnR6Pi-PP6i64woU6hqA7dhiGO1kprxo/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5MzMt/OTY4MS5qcGVn.jpeg",
				Width:       600,
				Height:      598,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/o1j8_MjdliyAcMtlwGD8tipW0G515qlZZwT-8N2HCj8/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5Mzgt/Mjk4MC5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/o1j8_MjdliyAcMtlwGD8tipW0G515qlZZwT-8N2HCj8/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5Mzgt/Mjk4MC5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/qFbEMMuUy8DhdbnSkJGw_hrUGsfhL8yeuqgxYAtXr3Y/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM5Mzgt/Mjk4MC5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
		},
		Thumb:           "https://i.discogs.com/UGFJFwk2IUyVgmh_BItSab76v2MhipF7YQyMPDstQzs/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTcxOTg5/LTE2MTk3MjM3ODQt/Mjk2MS5qcGVn.jpeg",
		EstimatedWeight: 690,
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

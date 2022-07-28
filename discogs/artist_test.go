package discogs

import (
	"context"
	"net/http"
	"testing"

	"github.com/peacefixation/go-discogs/discogs/auth"
)

func Test_GetArtist(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/artists/6418741", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		_, _ = w.Write([]byte(`{"name": "Dj Healer", "id": 6418741, "resource_url": "https://api.discogs.com/artists/6418741", "uri": "https://www.discogs.com/artist/6418741-Dj-Healer", "releases_url": "https://api.discogs.com/artists/6418741/releases", "images": [{"type": "primary", "uri": "https://i.discogs.com/eSKlGusbppKIDeKHCJaOVMqRFwM7YfGhuh9yJc9Cxlk/rs:fit/g:sm/q:90/h:381/w:509/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTY0MTg3/NDEtMTUyMzg2OTEw/Mi0yMTUyLnBuZw.jpeg", "resource_url": "https://i.discogs.com/eSKlGusbppKIDeKHCJaOVMqRFwM7YfGhuh9yJc9Cxlk/rs:fit/g:sm/q:90/h:381/w:509/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTY0MTg3/NDEtMTUyMzg2OTEw/Mi0yMTUyLnBuZw.jpeg", "uri150": "https://i.discogs.com/MjNWuvxe11tlcXkxKhy0cFjKKPRoGDHdtsTA_86OTHI/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTY0MTg3/NDEtMTUyMzg2OTEw/Mi0yMTUyLnBuZw.jpeg", "width": 509, "height": 381}], "profile": "", "aliases": [{"id": 1993158, "name": "Prince Of Denmark", "resource_url": "https://api.discogs.com/artists/1993158", "thumbnail_url": "https://i.discogs.com/ZKpS8hPpmqPNk9AfvPc9joPZdeqGMhH3PWff3XJ29IE/rs:fit/g:sm/q:40/h:312/w:472/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTE5OTMx/NTgtMTUxMzgwMjkw/Mi0xMzA3LmpwZWc.jpeg"}, {"id": 2300500, "name": "Traumprinz", "resource_url": "https://api.discogs.com/artists/2300500", "thumbnail_url": "https://i.discogs.com/sTtfp4CAKwR4ngQJ8y6nVP2oNWmwYhvB2MmmFpfxM9M/rs:fit/g:sm/q:40/h:300/w:300/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTIzMDA1/MDAtMTUxMzgwMTky/NC0zNjYzLmpwZWc.jpeg"}, {"id": 4123107, "name": "DJ Metatron", "resource_url": "https://api.discogs.com/artists/4123107", "thumbnail_url": ""}, {"id": 5615333, "name": "Dr Sun", "resource_url": "https://api.discogs.com/artists/5615333", "thumbnail_url": ""}, {"id": 6418725, "name": "Prime Minister of Doom", "resource_url": "https://api.discogs.com/artists/6418725", "thumbnail_url": "https://i.discogs.com/_GlkVeipiohdCXzwO5ZpBNIjg7X49qlTjBgEllQndIk/rs:fit/g:sm/q:40/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTY0MTg3/MjUtMTUyMzQ2NTE3/NC0zNjA0LnBuZw.jpeg"}, {"id": 7044647, "name": "Golden Baby", "resource_url": "https://api.discogs.com/artists/7044647", "thumbnail_url": "https://i.discogs.com/bhnzNjRsx8lLrQcfL8Af446fvYPCPNsS4UE_ehZlFQA/rs:fit/g:sm/q:40/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTcwNDQ2/NDctMTU1MjgyMDk5/OS0xNDg5LmpwZWc.jpeg"}, {"id": 8136077, "name": "The Phantasy", "resource_url": "https://api.discogs.com/artists/8136077", "thumbnail_url": ""}], "data_quality": "Needs Vote"}`))
	})

	client := NewClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	artist, err := client.GetArtist(context.Background(), 6418741)
	if err != nil {
		t.Fatal(err)
	}

	want := &Artist{
		Name:        "Dj Healer",
		ID:          6418741,
		ResourceURL: "https://api.discogs.com/artists/6418741",
		URI:         "https://www.discogs.com/artist/6418741-Dj-Healer",
		ReleasesURL: "https://api.discogs.com/artists/6418741/releases",
		Images: []Image{
			{
				Type:        "primary",
				URI:         "https://i.discogs.com/eSKlGusbppKIDeKHCJaOVMqRFwM7YfGhuh9yJc9Cxlk/rs:fit/g:sm/q:90/h:381/w:509/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTY0MTg3/NDEtMTUyMzg2OTEw/Mi0yMTUyLnBuZw.jpeg",
				ResourceURL: "https://i.discogs.com/eSKlGusbppKIDeKHCJaOVMqRFwM7YfGhuh9yJc9Cxlk/rs:fit/g:sm/q:90/h:381/w:509/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTY0MTg3/NDEtMTUyMzg2OTEw/Mi0yMTUyLnBuZw.jpeg",
				URI150:      "https://i.discogs.com/MjNWuvxe11tlcXkxKhy0cFjKKPRoGDHdtsTA_86OTHI/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTY0MTg3/NDEtMTUyMzg2OTEw/Mi0yMTUyLnBuZw.jpeg",
				Width:       509,
				Height:      381,
			},
		},
		Profile: "",
		Aliases: []Alias{
			{
				ID:           1993158,
				Name:         "Prince Of Denmark",
				ResourceURL:  "https://api.discogs.com/artists/1993158",
				ThumbnailURL: "https://i.discogs.com/ZKpS8hPpmqPNk9AfvPc9joPZdeqGMhH3PWff3XJ29IE/rs:fit/g:sm/q:40/h:312/w:472/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTE5OTMx/NTgtMTUxMzgwMjkw/Mi0xMzA3LmpwZWc.jpeg",
			},
			{
				ID:           2300500,
				Name:         "Traumprinz",
				ResourceURL:  "https://api.discogs.com/artists/2300500",
				ThumbnailURL: "https://i.discogs.com/sTtfp4CAKwR4ngQJ8y6nVP2oNWmwYhvB2MmmFpfxM9M/rs:fit/g:sm/q:40/h:300/w:300/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTIzMDA1/MDAtMTUxMzgwMTky/NC0zNjYzLmpwZWc.jpeg",
			},
			{
				ID:           4123107,
				Name:         "DJ Metatron",
				ResourceURL:  "https://api.discogs.com/artists/4123107",
				ThumbnailURL: "",
			},
			{
				ID:           5615333,
				Name:         "Dr Sun",
				ResourceURL:  "https://api.discogs.com/artists/5615333",
				ThumbnailURL: "",
			},
			{
				ID:           6418725,
				Name:         "Prime Minister of Doom",
				ResourceURL:  "https://api.discogs.com/artists/6418725",
				ThumbnailURL: "https://i.discogs.com/_GlkVeipiohdCXzwO5ZpBNIjg7X49qlTjBgEllQndIk/rs:fit/g:sm/q:40/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTY0MTg3/MjUtMTUyMzQ2NTE3/NC0zNjA0LnBuZw.jpeg",
			},
			{
				ID:           7044647,
				Name:         "Golden Baby",
				ResourceURL:  "https://api.discogs.com/artists/7044647",
				ThumbnailURL: "https://i.discogs.com/bhnzNjRsx8lLrQcfL8Af446fvYPCPNsS4UE_ehZlFQA/rs:fit/g:sm/q:40/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTcwNDQ2/NDctMTU1MjgyMDk5/OS0xNDg5LmpwZWc.jpeg",
			},
			{
				ID:           8136077,
				Name:         "The Phantasy",
				ResourceURL:  "https://api.discogs.com/artists/8136077",
				ThumbnailURL: "",
			},
		},
		DataQuality: "Needs Vote",
	}

	testEqual(t, want, artist)
}

func Test_GetArtistReleases(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/artists/6418741/releases", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		_, _ = w.Write([]byte(`{"pagination": {"page": 1, "pages": 1, "per_page": 50, "items": 3, "urls": {}}, "releases": [{"id": 1521183, "title": "Nothing 2 Loose", "type": "master", "main_release": 11849140, "artist": "DJ Healer", "role": "Main", "resource_url": "https://api.discogs.com/masters/1521183", "year": 2018, "thumb": "https://i.discogs.com/l71C5a4ersYHoN48qLm8KIpXSL9X8NngQgqDzKc0u1I/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTE2/MjItODA4NS5qcGVn.jpeg", "stats": {"community": {"in_wantlist": 1839, "in_collection": 1368}, "user": {"in_wantlist": 0, "in_collection": 0}}}, {"id": 12859092, "status": "Accepted", "type": "release", "format": "File, MP3, Mixed, 320", "label": "A Strangely Isolated Place", "title": "Isolatedmix 83", "resource_url": "https://api.discogs.com/releases/12859092", "role": "TrackAppearance", "artist": "Olaf Stuut / Inner River", "trackinfo": "Geminiden Regen", "year": 2018, "thumb": "https://i.discogs.com/4WgiXWK30OpyOeUtSkhZQ6YevVqtUjFceczNRsNR8pg/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEyODU5/MDkyLTE1NDMzMjYx/NjAtNDg1My5wbmc.jpeg", "stats": {"community": {"in_wantlist": 1, "in_collection": 9}, "user": {"in_wantlist": 0, "in_collection": 0}}}, {"id": 2514943, "title": "Hopes And Fears (Moritz Remix)", "type": "master", "main_release": 22200913, "artist": "DJ Healer", "role": "UnofficialRelease", "resource_url": "https://api.discogs.com/masters/2514943", "year": 2021, "thumb": "https://i.discogs.com/LuBHBdHbMC89VTDaNGa-vZQuXN4hPYh7N8hGgUXYYBk/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTIyMjAw/OTEzLTE2NDUxNDA2/NzktNDIzNC5qcGVn.jpeg", "stats": {"community": {"in_wantlist": 2, "in_collection": 0}, "user": {"in_wantlist": 0, "in_collection": 0}}}]}`))
	})

	client := NewClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	releases, err := client.GetArtistReleases(context.Background(), 6418741, nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &ArtistReleases{
		Pagination: Pagination{
			Page:    1,
			Pages:   1,
			PerPage: 50,
			Items:   3,
			URLs:    PaginationURLs{},
		},
		Releases: []ArtistRelease{
			{
				ID:          1521183,
				Title:       "Nothing 2 Loose",
				Type:        "master",
				MainRelease: 11849140,
				Artist:      "DJ Healer",
				Role:        "Main",
				ResourceURL: "https://api.discogs.com/masters/1521183",
				Year:        2018,
				Thumb:       "https://i.discogs.com/l71C5a4ersYHoN48qLm8KIpXSL9X8NngQgqDzKc0u1I/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTE2/MjItODA4NS5qcGVn.jpeg",
				Stats: Stats{
					Community: CommunityStats{
						InWantlist:   1839,
						InCollection: 1368,
					},
					User: UserStats{
						InWantlist:   0,
						InCollection: 0,
					},
				},
			},
			{
				ID:          12859092,
				Status:      "Accepted",
				Type:        "release",
				Format:      "File, MP3, Mixed, 320",
				Label:       "A Strangely Isolated Place",
				Title:       "Isolatedmix 83",
				ResourceURL: "https://api.discogs.com/releases/12859092",
				Role:        "TrackAppearance",
				Artist:      "Olaf Stuut / Inner River",
				TrackInfo:   "Geminiden Regen",
				Year:        2018,
				Thumb:       "https://i.discogs.com/4WgiXWK30OpyOeUtSkhZQ6YevVqtUjFceczNRsNR8pg/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEyODU5/MDkyLTE1NDMzMjYx/NjAtNDg1My5wbmc.jpeg",
				Stats: Stats{
					Community: CommunityStats{
						InWantlist:   1,
						InCollection: 9,
					},
					User: UserStats{
						InWantlist:   0,
						InCollection: 0,
					},
				},
			},
			{
				ID:          2514943,
				Title:       "Hopes And Fears (Moritz Remix)",
				Type:        "master",
				MainRelease: 22200913,
				Artist:      "DJ Healer",
				Role:        "UnofficialRelease",
				ResourceURL: "https://api.discogs.com/masters/2514943",
				Year:        2021,
				Thumb:       "https://i.discogs.com/LuBHBdHbMC89VTDaNGa-vZQuXN4hPYh7N8hGgUXYYBk/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTIyMjAw/OTEzLTE2NDUxNDA2/NzktNDIzNC5qcGVn.jpeg",
				Stats: Stats{
					Community: CommunityStats{
						InWantlist:   2,
						InCollection: 0,
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

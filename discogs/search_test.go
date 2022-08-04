package discogs

import (
	"context"
	"net/http"
	"testing"

	"github.com/peacefixation/go-discogs/discogs/auth"
	"github.com/peacefixation/go-discogs/discogs/param"
)

func Test_Search(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/database/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		_, _ = w.Write([]byte(`{"pagination": {"page": 1, "pages": 1, "per_page": 50, "items": 3, "urls": {}}, "results": [{"country": "Germany", "year": "2018", "format": ["Vinyl", "12\"", "33 \u2153 RPM", "Album"], "label": ["All Possible Worlds"], "type": "master", "genre": ["Electronic"], "style": ["Ambient", "Breakbeat", "House"], "id": 1521183, "barcode": ["APW-2 A Dmm H.P. mastered by rico at the basement all possible worlds", "APW-2 B Dmm H.P. mastered by rico at the basement all possible worlds", "APW-2 C Dmm H.P. mastered by rico at the basement all possible worlds", "APW-2 D Dmm H.P. mastered by rico at the basement all possible worlds", "APW-2 E Dmm H.P. mastered by rico at the basement all possible worlds", "APW-2 F Dmm H.P. mastered by rico at the basement all possible worlds"], "user_data": {"in_wantlist": false, "in_collection": false}, "master_id": 1521183, "master_url": "https://api.discogs.com/masters/1521183", "uri": "/DJ-Healer-Nothing-2-Loose/master/1521183", "catno": "APW-2", "title": "DJ Healer - Nothing 2 Loose", "thumb": "https://i.discogs.com/l71C5a4ersYHoN48qLm8KIpXSL9X8NngQgqDzKc0u1I/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTE2/MjItODA4NS5qcGVn.jpeg", "cover_image": "https://i.discogs.com/qAqSEZsVcMvvyMJfKXbv7KdmxY0HMvZLYx4eMt7kvPc/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTE2/MjItODA4NS5qcGVn.jpeg", "resource_url": "https://api.discogs.com/masters/1521183", "community": {"want": 2496, "have": 1850}}, {"country": "Germany", "year": "2018", "format": ["Vinyl", "12\"", "33 \u2153 RPM", "Album"], "label": ["All Possible Worlds"], "type": "release", "genre": ["Electronic"], "style": ["Ambient", "Breakbeat", "House"], "id": 11849140, "barcode": ["APW-2 A Dmm H.P. mastered by rico at the basement all possible worlds", "APW-2 B Dmm H.P. mastered by rico at the basement all possible worlds", "APW-2 C Dmm H.P. mastered by rico at the basement all possible worlds", "APW-2 D Dmm H.P. mastered by rico at the basement all possible worlds", "APW-2 E Dmm H.P. mastered by rico at the basement all possible worlds", "APW-2 F Dmm H.P. mastered by rico at the basement all possible worlds"], "user_data": {"in_wantlist": false, "in_collection": false}, "master_id": 1521183, "master_url": "https://api.discogs.com/masters/1521183", "uri": "/DJ-Healer-Nothing-2-Loose/release/11849140", "catno": "APW-2", "title": "DJ Healer - Nothing 2 Loose", "thumb": "https://i.discogs.com/l71C5a4ersYHoN48qLm8KIpXSL9X8NngQgqDzKc0u1I/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTE2/MjItODA4NS5qcGVn.jpeg", "cover_image": "https://i.discogs.com/qAqSEZsVcMvvyMJfKXbv7KdmxY0HMvZLYx4eMt7kvPc/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTE2/MjItODA4NS5qcGVn.jpeg", "resource_url": "https://api.discogs.com/releases/11849140", "community": {"want": 1839, "have": 1368}, "format_quantity": 3, "formats": [{"name": "Vinyl", "qty": "3", "descriptions": ["12\"", "33 \u2153 RPM", "Album"]}]}, {"country": "Germany", "year": "2019", "format": ["Vinyl", "12\"", "33 \u2153 RPM", "Album", "Repress"], "label": ["All Possible Worlds"], "type": "release", "genre": ["Electronic"], "style": ["Ambient", "Breakbeat", "House"], "id": 13416116, "barcode": ["APW-2 A", "APW-2 B", "APW-2 C", "APW-2 D", "APW-2 E", "APW-2 F"], "user_data": {"in_wantlist": false, "in_collection": false}, "master_id": 1521183, "master_url": "https://api.discogs.com/masters/1521183", "uri": "/DJ-Healer-Nothing-2-Loose/release/13416116", "catno": "APW-2", "title": "DJ Healer - Nothing 2 Loose", "thumb": "https://i.discogs.com/o73HPuPE1awkg3sa0VTQQmYbFo7DOLMq-SIzEBGPoBQ/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEzNDE2/MTE2LTE1NTM3ODk4/ODItMzEwNi5qcGVn.jpeg", "cover_image": "https://i.discogs.com/gxvY56hirLOf7yv6uCVl6vtiIJOFOqY51_6Y57AeOFM/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEzNDE2/MTE2LTE1NTM3ODk4/ODItMzEwNi5qcGVn.jpeg", "resource_url": "https://api.discogs.com/releases/13416116", "community": {"want": 657, "have": 482}, "format_quantity": 3, "formats": [{"name": "Vinyl", "qty": "3", "descriptions": ["12\"", "33 \u2153 RPM", "Album", "Repress"]}]}]}`))
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	searchTerm := "DJ Healer"
	params := []param.Param{
		param.NewSearchArtist("DJ Healer"),
		param.NewSearchTitle("Nothing 2 Loose"),
	}

	search, err := client.Search(context.Background(), searchTerm, params)
	if err != nil {
		t.Fatal(err)
	}

	want := &SearchResponse{
		Pagination: Pagination{
			Page:    1,
			Pages:   1,
			Items:   3,
			PerPage: 50,
			URLs: PaginationURLs{
				First: "",
				Prev:  "",
				Next:  "",
				Last:  "",
			},
		},
		Results: []SearchResult{
			{
				Style: []string{
					"Ambient",
					"Breakbeat",
					"House",
				},
				Thumb:   "https://i.discogs.com/l71C5a4ersYHoN48qLm8KIpXSL9X8NngQgqDzKc0u1I/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTE2/MjItODA4NS5qcGVn.jpeg",
				Title:   "DJ Healer - Nothing 2 Loose",
				Country: "Germany",
				Format: []string{
					"Vinyl",
					"12\"",
					"33 ⅓ RPM",
					"Album",
				},
				URI: "/DJ-Healer-Nothing-2-Loose/master/1521183",
				Community: SearchResultCommunity{
					Want: 2496,
					Have: 1850,
				},
				Label: []string{
					"All Possible Worlds",
				},
				CatNo: "APW-2",
				Year:  "2018",
				Genre: []string{
					"Electronic",
				},
				ResourceURL: "https://api.discogs.com/masters/1521183",
				Type:        "master",
				ID:          1521183,
			},
			{
				Style: []string{
					"Ambient",
					"Breakbeat",
					"House",
				},
				Thumb:   "https://i.discogs.com/l71C5a4ersYHoN48qLm8KIpXSL9X8NngQgqDzKc0u1I/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTE2/MjItODA4NS5qcGVn.jpeg",
				Title:   "DJ Healer - Nothing 2 Loose",
				Country: "Germany",
				Format: []string{
					"Vinyl",
					"12\"",
					"33 ⅓ RPM",
					"Album",
				},
				URI: "/DJ-Healer-Nothing-2-Loose/release/11849140",
				Community: SearchResultCommunity{
					Want: 1839,
					Have: 1368,
				},
				Label: []string{
					"All Possible Worlds",
				},
				CatNo: "APW-2",
				Year:  "2018",
				Genre: []string{
					"Electronic",
				},
				ResourceURL: "https://api.discogs.com/releases/11849140",
				Type:        "release",
				ID:          11849140,
			},
			{
				Style: []string{
					"Ambient",
					"Breakbeat",
					"House",
				},
				Thumb:   "https://i.discogs.com/o73HPuPE1awkg3sa0VTQQmYbFo7DOLMq-SIzEBGPoBQ/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEzNDE2/MTE2LTE1NTM3ODk4/ODItMzEwNi5qcGVn.jpeg",
				Title:   "DJ Healer - Nothing 2 Loose",
				Country: "Germany",
				Format: []string{
					"Vinyl",
					"12\"",
					"33 ⅓ RPM",
					"Album",
					"Repress",
				},
				URI: "/DJ-Healer-Nothing-2-Loose/release/13416116",
				Community: SearchResultCommunity{
					Want: 657,
					Have: 482,
				},
				Label: []string{
					"All Possible Worlds",
				},
				CatNo: "APW-2",
				Year:  "2019",
				Genre: []string{
					"Electronic",
				},
				ResourceURL: "https://api.discogs.com/releases/13416116",
				Type:        "release",
				ID:          13416116,
			},
		},
	}

	testEqual(t, want, search)
}

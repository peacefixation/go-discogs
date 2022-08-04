package discogs

import (
	"context"
	"net/http"
	"testing"

	"github.com/peacefixation/go-discogs/discogs/auth"
)

func Test_GetCollectionFolders(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/Username/collection/folders", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		_, _ = w.Write([]byte(`{"folders": [{"id": 0, "name": "All", "count": 245, "resource_url": "https://api.discogs.com/users/Username/collection/folders/0"}, {"id": 1, "name": "Uncategorized", "count": 245, "resource_url": "https://api.discogs.com/users/Username/collection/folders/1"}]}`))
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	resp, err := client.GetCollectionFolders(context.Background(), "Username")
	if err != nil {
		t.Fatal(err)
	}

	want := &CollectionFolders{
		Folders: []CollectionFolder{
			{
				ID:          0,
				Name:        "All",
				Count:       245,
				ResourceURL: "https://api.discogs.com/users/Username/collection/folders/0",
			},
			{
				ID:          1,
				Name:        "Uncategorized",
				Count:       245,
				ResourceURL: "https://api.discogs.com/users/Username/collection/folders/1",
			},
		},
	}

	testEqual(t, want, resp)
}

func Test_CreateCollectionFolder(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/Username/collection/folders", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id": 4739258, "name": "test-folder", "count": 0, "resource_url": "https://api.discogs.com/users/Username/collection/folders/4739258"}`))
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	resp, err := client.CreateCollectionFolder(context.Background(), "Username", "test-folder")
	if err != nil {
		t.Fatal(err)
	}

	want := &CollectionFolder{
		ID:          4739258,
		Name:        "test-folder",
		Count:       0,
		ResourceURL: "https://api.discogs.com/users/Username/collection/folders/4739258",
	}

	testEqual(t, want, resp)
}

func Test_GetCollectionFolder(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/Username/collection/folders/4739258", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		_, _ = w.Write([]byte(`{"id": 4739258, "name": "test-folder", "count": 0, "resource_url": "https://api.discogs.com/users/Username/collection/folders/4739258"}`))
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	resp, err := client.GetCollectionFolder(context.Background(), "Username", 4739258)
	if err != nil {
		t.Fatal(err)
	}

	want := &CollectionFolder{
		ID:          4739258,
		Name:        "test-folder",
		Count:       0,
		ResourceURL: "https://api.discogs.com/users/Username/collection/folders/4739258",
	}

	testEqual(t, want, resp)
}

func Test_UpdateCollectionFolder(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/Username/collection/folders/4739258", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		_, _ = w.Write([]byte(`{"id": 4739258, "name": "test-folder-2", "count": 0, "resource_url": "https://api.discogs.com/users/Username/collection/folders/4739258"}`))
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	resp, err := client.UpdateCollectionFolder(context.Background(), "Username", 4739258, "test-folder-2")
	if err != nil {
		t.Fatal(err)
	}

	want := &CollectionFolder{
		ID:          4739258,
		Name:        "test-folder-2",
		Count:       0,
		ResourceURL: "https://api.discogs.com/users/Username/collection/folders/4739258",
	}

	testEqual(t, want, resp)
}

func Test_DeleteCollectionFolder(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/Username/collection/folders/4739258", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)

		w.WriteHeader(http.StatusNoContent)
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	err := client.DeleteCollectionFolder(context.Background(), "Username", 4739258)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_GetCollectionItemsByRelease(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/Username/collection/releases/20078449", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		_, _ = w.Write([]byte(`{"pagination": {"page": 1, "pages": 1, "per_page": 50, "items": 1, "urls": {}}, "releases": [{"id": 20078449, "instance_id": 1081693136, "date_added": "2022-07-17T22:30:56-07:00", "rating": 0, "basic_information": {"id": 20078449, "master_id": 2299618, "master_url": "https://api.discogs.com/masters/2299618", "resource_url": "https://api.discogs.com/releases/20078449", "thumb": "https://i.discogs.com/6W0-_QHC8jWSMj7LGWm9l5QJVtVCGth-kFqaoakq9lA/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTIwMDc4/NDQ5LTE2MzA1MjE3/NTAtNzIzMy5qcGVn.jpeg", "cover_image": "https://i.discogs.com/ofHIh1Ty3p9kiPagzEs598mxJOuaWcnOr3kQYMcb-dA/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTIwMDc4/NDQ5LTE2MzA1MjE3/NTAtNzIzMy5qcGVn.jpeg", "title": "Realism EP", "year": 2021, "formats": [{"name": "Vinyl", "qty": "1", "text": "White Marble", "descriptions": ["12\"", "33 \u2153 RPM", "EP", "Repress"]}], "labels": [{"name": "Reality Used To Be A Friend Of Mine", "catno": "REALITY 20182", "entity_type": "1", "entity_type_name": "Label", "id": 1365217, "resource_url": "https://api.discogs.com/labels/1365217"}], "artists": [{"name": "Rising Sun (7)", "anv": "", "join": "", "role": "", "tracks": "", "id": 1167655, "resource_url": "https://api.discogs.com/artists/1167655"}], "genres": ["Electronic"], "styles": ["Abstract", "Ambient", "Breakbeat", "Chillwave", "Deep House"]}, "folder_id": 1}]}`))
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	resp, err := client.GetCollectionItemsByRelease(context.Background(), "Username", 20078449)
	if err != nil {
		t.Fatal(err)
	}

	want := &CollectionItems{
		Pagination: Pagination{
			Page:    1,
			Pages:   1,
			PerPage: 50,
			Items:   1,
			URLs:    PaginationURLs{},
		},
		Releases: []CollectionRelease{
			{
				ID:         20078449,
				InstanceID: 1081693136,
				DateAdded:  "2022-07-17T22:30:56-07:00",
				Rating:     0,
				BasicInformation: BasicInformation{
					ID:          20078449,
					MasterID:    2299618,
					MasterURL:   "https://api.discogs.com/masters/2299618",
					ResourceURL: "https://api.discogs.com/releases/20078449",
					Thumb:       "https://i.discogs.com/6W0-_QHC8jWSMj7LGWm9l5QJVtVCGth-kFqaoakq9lA/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTIwMDc4/NDQ5LTE2MzA1MjE3/NTAtNzIzMy5qcGVn.jpeg",
					CoverImage:  "https://i.discogs.com/ofHIh1Ty3p9kiPagzEs598mxJOuaWcnOr3kQYMcb-dA/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTIwMDc4/NDQ5LTE2MzA1MjE3/NTAtNzIzMy5qcGVn.jpeg",
					Title:       "Realism EP",
					Year:        2021,
					Formats: []Format{
						{
							Name:     "Vinyl",
							Quantity: "1",
							Text:     "White Marble",
							Descriptions: []string{
								"12\"",
								"33 \u2153 RPM",
								"EP",
								"Repress",
							},
						},
					},
					Labels: []ReleaseLabel{
						{
							Name:           "Reality Used To Be A Friend Of Mine",
							CatNo:          "REALITY 20182",
							EntityType:     "1",
							EntityTypeName: "Label",
							ID:             1365217,
							ResourceURL:    "https://api.discogs.com/labels/1365217",
						},
					},
					Artists: []ReleaseArtist{
						{
							Name:        "Rising Sun (7)",
							ANV:         "",
							Join:        "",
							Role:        "",
							Tracks:      "",
							ID:          1167655,
							ResourceURL: "https://api.discogs.com/artists/1167655",
						},
					},
					Genres: []string{
						"Electronic",
					},
					Styles: []string{
						"Abstract",
						"Ambient",
						"Breakbeat",
						"Chillwave",
						"Deep House",
					},
				},
				FolderID: 1,
			},
		},
	}

	testEqual(t, want, resp)
}

func Test_GetCollectionItemsByFolder(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/Username/collection/folders/0/releases", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		_, _ = w.Write([]byte(`{"pagination":{"page":1,"pages":1,"per_page":50,"items":1,"urls":{}},"releases":[{"id":54524,"instance_id":519662456,"date_added":"2020-10-28T09:37:50-07:00","rating":0,"basic_information":{"id":54524,"master_id":36105,"master_url":"https://api.discogs.com/masters/36105","resource_url":"https://api.discogs.com/releases/54524","thumb":"https://i.discogs.com/M9b2jLZlZrWkv2SMzIvN_z-FcZD_jlX9m8HuWK-oonw/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NTI0/LTE1MzY4OTQxMjUt/Njg0Mi5qcGVn.jpeg","cover_image":"https://i.discogs.com/7AFzapz17kuF_XZcOnI7CxqWgDiBMwHJC-rkU2DeJHE/rs:fit/g:sm/q:90/h:599/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NTI0/LTE1MzY4OTQxMjUt/Njg0Mi5qcGVn.jpeg","title":"Purple","year":1997,"formats":[{"name":"Vinyl","qty":"1","descriptions":["12\"","45 RPM","Promo","White Label"]}],"labels":[{"name":"4AD","catno":"GUS 7","entity_type":"1","entity_type_name":"Label","id":634,"resource_url":"https://api.discogs.com/labels/634"}],"artists":[{"name":"GusGus","anv":"Gus Gus","join":"","role":"","tracks":"","id":231513,"resource_url":"https://api.discogs.com/artists/231513"}],"genres": ["Electronic"],"styles": ["Breaks","Downtempo","Ambient"]},"folder_id": 1,"notes":[{"field_id": 1,"value":"Very Good Plus (VG+)"},{"field_id":2,"value":"Generic"}]}]}`))
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	resp, err := client.GetCollectionItemsByFolder(context.Background(), "Username", 0, nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &CollectionItems{
		Pagination: Pagination{
			Page:    1,
			Pages:   1,
			PerPage: 50,
			Items:   1,
			URLs:    PaginationURLs{},
		},
		Releases: []CollectionRelease{
			{
				ID:         54524,
				InstanceID: 519662456,
				DateAdded:  "2020-10-28T09:37:50-07:00",
				Rating:     0,
				BasicInformation: BasicInformation{
					ID:          54524,
					MasterID:    36105,
					MasterURL:   "https://api.discogs.com/masters/36105",
					ResourceURL: "https://api.discogs.com/releases/54524",
					Thumb:       "https://i.discogs.com/M9b2jLZlZrWkv2SMzIvN_z-FcZD_jlX9m8HuWK-oonw/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NTI0/LTE1MzY4OTQxMjUt/Njg0Mi5qcGVn.jpeg",
					CoverImage:  "https://i.discogs.com/7AFzapz17kuF_XZcOnI7CxqWgDiBMwHJC-rkU2DeJHE/rs:fit/g:sm/q:90/h:599/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NTI0/LTE1MzY4OTQxMjUt/Njg0Mi5qcGVn.jpeg",
					Title:       "Purple",
					Year:        1997,
					Formats: []Format{
						{
							Name:     "Vinyl",
							Quantity: "1",
							Descriptions: []string{
								"12\"",
								"45 RPM",
								"Promo",
								"White Label",
							},
						},
					},
					Labels: []ReleaseLabel{
						{
							Name:           "4AD",
							CatNo:          "GUS 7",
							EntityType:     "1",
							EntityTypeName: "Label",
							ID:             634,
							ResourceURL:    "https://api.discogs.com/labels/634",
						},
					},
					Artists: []ReleaseArtist{
						{
							Name:        "GusGus",
							ANV:         "Gus Gus",
							Join:        "",
							Role:        "",
							Tracks:      "",
							ID:          231513,
							ResourceURL: "https://api.discogs.com/artists/231513",
						},
					},
					Genres: []string{
						"Electronic",
					},
					Styles: []string{
						"Breaks",
						"Downtempo",
						"Ambient",
					},
				},
				FolderID: 1,
				Notes: []Note{
					{
						FieldID: 0,
						Value:   "Very Good Plus (VG+)",
					},
					{
						FieldID: 0,
						Value:   "Generic",
					},
				},
			},
		},
	}

	testEqual(t, want, resp)
}

func Test_AddReleaseToCollectionFolder(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/Username/collection/folders/1/releases/213658", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id": 213658, "instance_id": 1089887351, "rating": 0, "basic_information": {"id": 213658, "master_id": 36506, "master_url": "https://api.discogs.com/masters/36506", "resource_url": "https://api.discogs.com/releases/213658", "cover_image": "https://i.discogs.com/rWJlcOi9Haq_2CcMPws9cvHyq209o9Lj-RwN-vlU6wE/rs:fit/g:sm/q:90/h:600/w:593/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTIxMzY1/OC0xMjAyMjg3NDk4/LmpwZWc.jpeg", "thumb": "https://i.discogs.com/fdcbq5t-Q1O5NlZu2m2cwRKl7v3-kjR2wCOnx9b43xY/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTIxMzY1/OC0xMjAyMjg3NDk4/LmpwZWc.jpeg", "title": "Robot.O.Chan", "year": 2004, "formats": [{"name": "CD", "qty": "1", "descriptions": ["Album"]}], "labels": [{"name": "Twisted Records", "catno": "TWSCD25", "entity_type": "1", "entity_type_name": "Label", "id": 3336, "resource_url": "https://api.discogs.com/labels/3336"}], "artists": [{"name": "Prometheus", "anv": "", "join": "", "role": "", "tracks": "", "id": 19870, "resource_url": "https://api.discogs.com/artists/19870"}]}, "date_added": "2022-07-28T04:28:27-07:00", "folder_id": 1, "resource_url": "https://api.discogs.com/users/Username/collection/folders/1/releases/213658?instance_id=1089887351"}`))
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	resp, err := client.AddReleaseToCollectionFolder(context.Background(), "Username", 1, 213658)
	if err != nil {
		t.Fatal(err)
	}

	want := &CollectionRelease{
		ID:         213658,
		InstanceID: 1089887351,
		Rating:     0,
		BasicInformation: BasicInformation{
			ID:          213658,
			MasterID:    36506,
			MasterURL:   "https://api.discogs.com/masters/36506",
			ResourceURL: "https://api.discogs.com/releases/213658",
			CoverImage:  "https://i.discogs.com/rWJlcOi9Haq_2CcMPws9cvHyq209o9Lj-RwN-vlU6wE/rs:fit/g:sm/q:90/h:600/w:593/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTIxMzY1/OC0xMjAyMjg3NDk4/LmpwZWc.jpeg",
			Thumb:       "https://i.discogs.com/fdcbq5t-Q1O5NlZu2m2cwRKl7v3-kjR2wCOnx9b43xY/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTIxMzY1/OC0xMjAyMjg3NDk4/LmpwZWc.jpeg",
			Title:       "Robot.O.Chan",
			Year:        2004,
			Formats: []Format{
				{
					Name:     "CD",
					Quantity: "1",
					Descriptions: []string{
						"Album",
					},
				},
			},
			Labels: []ReleaseLabel{
				{
					Name:           "Twisted Records",
					CatNo:          "TWSCD25",
					EntityType:     "1",
					EntityTypeName: "Label",
					ID:             3336,
					ResourceURL:    "https://api.discogs.com/labels/3336",
				},
			},
			Artists: []ReleaseArtist{
				{
					Name:        "Prometheus",
					ANV:         "",
					Join:        "",
					Role:        "",
					Tracks:      "",
					ID:          19870,
					ResourceURL: "https://api.discogs.com/artists/19870",
				},
			},
		},
		DateAdded:   "2022-07-28T04:28:27-07:00",
		FolderID:    1,
		ResourceURL: "https://api.discogs.com/users/Username/collection/folders/1/releases/213658?instance_id=1089887351",
	}

	testEqual(t, want, resp)
}

func Test_DeleteReleaseFromCollectionFolder(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/Username/collection/folders/1/releases/213658/instances/1089857009", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)

		w.WriteHeader(http.StatusNoContent)
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	err := client.DeleteReleaseFromCollectionFolder(context.Background(), "Username", 1, 213658, 1089857009)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ListCustomFields(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/Username/collection/fields", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		_, _ = w.Write([]byte(`{"fields": [{"id": 1, "name": "Media Condition", "type": "dropdown", "position": 1, "public": false, "options": ["Mint (M)", "Near Mint (NM or M-)", "Very Good Plus (VG+)", "Very Good (VG)", "Good Plus (G+)", "Good (G)", "Fair (F)", "Poor (P)"]}, {"id": 2, "name": "Sleeve Condition", "type": "dropdown", "position": 2, "public": false, "options": ["Generic", "No Cover", "Mint (M)", "Near Mint (NM or M-)", "Very Good Plus (VG+)", "Very Good (VG)", "Good Plus (G+)", "Good (G)", "Fair (F)", "Poor (P)"]}, {"id": 3, "name": "Notes", "type": "textarea", "position": 3, "public": false, "lines": 3}]}`))
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	resp, err := client.ListCustomFields(context.Background(), "Username")
	if err != nil {
		t.Fatal(err)
	}

	want := &ListFields{
		Fields: []ListField{
			{
				ID:       1,
				Name:     "Media Condition",
				Type:     "dropdown",
				Position: 1,
				Public:   false,
				Options: []string{
					"Mint (M)",
					"Near Mint (NM or M-)",
					"Very Good Plus (VG+)",
					"Very Good (VG)",
					"Good Plus (G+)",
					"Good (G)",
					"Fair (F)",
					"Poor (P)",
				},
			},
			{
				ID:       2,
				Name:     "Sleeve Condition",
				Type:     "dropdown",
				Position: 2,
				Public:   false,
				Options: []string{
					"Generic",
					"No Cover",
					"Mint (M)",
					"Near Mint (NM or M-)",
					"Very Good Plus (VG+)",
					"Very Good (VG)",
					"Good Plus (G+)",
					"Good (G)",
					"Fair (F)",
					"Poor (P)",
				},
			},
			{
				ID:       3,
				Name:     "Notes",
				Type:     "textarea",
				Position: 3,
				Public:   false,
				Lines:    3,
			},
		},
	}

	testEqual(t, want, resp)
}

func Test_EditField(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/Username/collection/folders/1/releases/213658/instances/1089887351/fields/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		w.WriteHeader(http.StatusNoContent)
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	err := client.EditFieldsInstance(context.Background(), "Username", 1, 213658, 1089887351, 1, "Very Good (VG)")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_GetCollectionValue(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/Username/collection/value", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		_, _ = w.Write([]byte(`{"minimum": "A$1,597.46", "median": "A$3,506.77", "maximum": "A$7,505.54"}`))
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	resp, err := client.GetCollectionValue(context.Background(), "Username")
	if err != nil {
		t.Fatal(err)
	}

	want := &CollectionValue{
		Maximum: "A$7,505.54",
		Median:  "A$3,506.77",
		Minimum: "A$1,597.46",
	}

	testEqual(t, want, resp)
}

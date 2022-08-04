package discogs

import (
	"context"
	"net/http"
	"testing"

	"github.com/peacefixation/go-discogs/discogs/auth"
)

func Test_GetMasterRelease(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/masters/1521183", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		_, _ = w.Write([]byte(`{"id": 1521183, "main_release": 11849140, "most_recent_release": 13416116, "resource_url": "https://api.discogs.com/masters/1521183", "uri": "https://www.discogs.com/master/1521183-DJ-Healer-Nothing-2-Loose", "versions_url": "https://api.discogs.com/masters/1521183/versions", "main_release_url": "https://api.discogs.com/releases/11849140", "most_recent_release_url": "https://api.discogs.com/releases/13416116", "num_for_sale": 19, "lowest_price": 151.52, "images": [{"type": "primary", "uri": "https://i.discogs.com/qAqSEZsVcMvvyMJfKXbv7KdmxY0HMvZLYx4eMt7kvPc/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTE2/MjItODA4NS5qcGVn.jpeg", "resource_url": "https://i.discogs.com/qAqSEZsVcMvvyMJfKXbv7KdmxY0HMvZLYx4eMt7kvPc/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTE2/MjItODA4NS5qcGVn.jpeg", "uri150": "https://i.discogs.com/l71C5a4ersYHoN48qLm8KIpXSL9X8NngQgqDzKc0u1I/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTE2/MjItODA4NS5qcGVn.jpeg", "width": 600, "height": 600}, {"type": "secondary", "uri": "https://i.discogs.com/809DLP4xiHA0uT2TolWEwRonZpCAUDZCoLcnmCr4LnY/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTg4/NzMtOTA0Mi5qcGVn.jpeg", "resource_url": "https://i.discogs.com/809DLP4xiHA0uT2TolWEwRonZpCAUDZCoLcnmCr4LnY/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTg4/NzMtOTA0Mi5qcGVn.jpeg", "uri150": "https://i.discogs.com/yMBLLOFr5r6VnDPaJtew1oLy8BB50UdgONP3VjeY7NI/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTg4/NzMtOTA0Mi5qcGVn.jpeg", "width": 600, "height": 600}, {"type": "secondary", "uri": "https://i.discogs.com/0ed3On5prgm7ZLmKx2rSqpXMlM1svDCkde1m1dZzYs0/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NDIz/MzItMzA5MC5qcGVn.jpeg", "resource_url": "https://i.discogs.com/0ed3On5prgm7ZLmKx2rSqpXMlM1svDCkde1m1dZzYs0/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NDIz/MzItMzA5MC5qcGVn.jpeg", "uri150": "https://i.discogs.com/wnhuiqzldI-4hn7Rao3-pu-R5kIozka3F7ouFoSEaZw/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NDIz/MzItMzA5MC5qcGVn.jpeg", "width": 600, "height": 600}, {"type": "secondary", "uri": "https://i.discogs.com/k1z8DRjz5sqbTxaNHyvx9mV32NyANOENWS-rc3BPk2U/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjQ5MTE2/ODEtOTU0OS5qcGVn.jpeg", "resource_url": "https://i.discogs.com/k1z8DRjz5sqbTxaNHyvx9mV32NyANOENWS-rc3BPk2U/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjQ5MTE2/ODEtOTU0OS5qcGVn.jpeg", "uri150": "https://i.discogs.com/FkjqlfBtGzN9kzx3ThUT0MoYoR6Afa7_AOIJUDww1PI/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjQ5MTE2/ODEtOTU0OS5qcGVn.jpeg", "width": 600, "height": 600}, {"type": "secondary", "uri": "https://i.discogs.com/a94M2BpoDjcExLXgDh6Bi32VozF5QkRPHbg5kNkd5ZQ/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NDIz/MjMtOTY4MS5qcGVn.jpeg", "resource_url": "https://i.discogs.com/a94M2BpoDjcExLXgDh6Bi32VozF5QkRPHbg5kNkd5ZQ/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NDIz/MjMtOTY4MS5qcGVn.jpeg", "uri150": "https://i.discogs.com/qhTAIPPvjT_-MUCxdMLKLGCpO3vSakq3AjhmpFsUHaE/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NDIz/MjMtOTY4MS5qcGVn.jpeg", "width": 600, "height": 600}, {"type": "secondary", "uri": "https://i.discogs.com/uOPk_nMhbjWMpRfHG4D0ZtAFot8AZtCjmKd49X6f79c/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE1/OTUtMTcwOC5qcGVn.jpeg", "resource_url": "https://i.discogs.com/uOPk_nMhbjWMpRfHG4D0ZtAFot8AZtCjmKd49X6f79c/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE1/OTUtMTcwOC5qcGVn.jpeg", "uri150": "https://i.discogs.com/cK3AiimTiQmFL4dJWvMeJVnMpQde8DRAW8YgFlsmlfU/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE1/OTUtMTcwOC5qcGVn.jpeg", "width": 600, "height": 600}, {"type": "secondary", "uri": "https://i.discogs.com/B0ajqnhT-hXS3SjzJmEOCyN8jpfNG7hdWApA9_4MsxU/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTEtNTQ1MS5qcGVn.jpeg", "resource_url": "https://i.discogs.com/B0ajqnhT-hXS3SjzJmEOCyN8jpfNG7hdWApA9_4MsxU/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTEtNTQ1MS5qcGVn.jpeg", "uri150": "https://i.discogs.com/VVvhrf_deK_bjraHaSVuMaN5DR7T7R6UgY502TvL3Ok/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTEtNTQ1MS5qcGVn.jpeg", "width": 600, "height": 600}, {"type": "secondary", "uri": "https://i.discogs.com/meRonYJSZNlJtrBfB6MnvR-2L7jDgw0cQEcCUaNR2kU/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MDktOTA4NS5qcGVn.jpeg", "resource_url": "https://i.discogs.com/meRonYJSZNlJtrBfB6MnvR-2L7jDgw0cQEcCUaNR2kU/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MDktOTA4NS5qcGVn.jpeg", "uri150": "https://i.discogs.com/pixDcsIIYZ5ajhbm6Q7zMt9cSbb-OosiWdZ6M5Qjx1k/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MDktOTA4NS5qcGVn.jpeg", "width": 600, "height": 600}, {"type": "secondary", "uri": "https://i.discogs.com/46C3fFL6V-WErSwPlS_edVnZMmW6FysZO6Y1BIY-3Ak/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTYtNDk3Ni5qcGVn.jpeg", "resource_url": "https://i.discogs.com/46C3fFL6V-WErSwPlS_edVnZMmW6FysZO6Y1BIY-3Ak/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTYtNDk3Ni5qcGVn.jpeg", "uri150": "https://i.discogs.com/slARIJKnFPs3FLQ4YSX6ugmZmv_f16D8DKUHgfHQ2O0/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTYtNDk3Ni5qcGVn.jpeg", "width": 600, "height": 600}, {"type": "secondary", "uri": "https://i.discogs.com/a5a20FS60UjYfkgDqzrlqFxwvQVzgGAyR2MW0PZujbg/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTAtNDY5MS5qcGVn.jpeg", "resource_url": "https://i.discogs.com/a5a20FS60UjYfkgDqzrlqFxwvQVzgGAyR2MW0PZujbg/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTAtNDY5MS5qcGVn.jpeg", "uri150": "https://i.discogs.com/Vv8BuIlMHXeAmdJJVGb8idLxLxEPYE_I9shCHeCGzZk/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTAtNDY5MS5qcGVn.jpeg", "width": 600, "height": 600}, {"type": "secondary", "uri": "https://i.discogs.com/LgGqid4ErvSC8Wn6yA0_zQXgB7CIrDrOLGvUq-DCBKc/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTYtNDQ2OS5qcGVn.jpeg", "resource_url": "https://i.discogs.com/LgGqid4ErvSC8Wn6yA0_zQXgB7CIrDrOLGvUq-DCBKc/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTYtNDQ2OS5qcGVn.jpeg", "uri150": "https://i.discogs.com/mmbqWgflsRzZcUdM26OU1-uaCqoc7BxAQqOn2FJONxs/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTYtNDQ2OS5qcGVn.jpeg", "width": 600, "height": 600}], "genres": ["Electronic"], "styles": ["Ambient", "Breakbeat", "House"], "year": 2018, "tracklist": [{"position": "A1", "type_": "track", "title": "At Last (Becalming The Storm)", "duration": "3:27"}, {"position": "A2", "type_": "track", "title": "Great Escape", "duration": "3:08"}, {"position": "B", "type_": "track", "title": "2 The Dark", "duration": "8:24"}, {"position": "C1", "type_": "track", "title": "Gods Creation", "duration": "8:01"}, {"position": "C2", "type_": "track", "title": "Untitled", "duration": "1:52"}, {"position": "D1", "type_": "track", "title": "Planet Lonely", "duration": "6:42"}, {"position": "D2", "type_": "track", "title": "Hopes And Fears", "duration": "6:11"}, {"position": "E", "type_": "track", "title": "We Are Going Nowhere", "duration": "12:08"}, {"position": "F1", "type_": "track", "title": "The Interview", "duration": "6:10"}, {"position": "F2", "type_": "track", "title": "Gone", "duration": "3:45"}, {"position": "F3", "type_": "track", "title": "Protectionspell", "duration": "8:07"}], "artists": [{"name": "DJ Healer", "anv": "", "join": "", "role": "", "tracks": "", "id": 6418741, "resource_url": "https://api.discogs.com/artists/6418741", "thumbnail_url": "https://i.discogs.com/yHVbfC3naoozshAMSvhInxFb7EsS03YGWBLuZvGLjbQ/rs:fit/g:sm/q:40/h:381/w:509/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTY0MTg3/NDEtMTUyMzg2OTEw/Mi0yMTUyLnBuZw.jpeg"}], "title": "Nothing 2 Loose", "data_quality": "Correct", "videos": [{"uri": "https://www.youtube.com/watch?v=8hB_WMlyDEM", "title": "Dj Healer - At Last (Becalming The Storm) [APW-2]", "description": "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ", "duration": 208, "embed": true}, {"uri": "https://www.youtube.com/watch?v=n2fMoxSZZEU", "title": "Dj Healer - The Interview [APW-2]", "description": "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ", "duration": 370, "embed": true}, {"uri": "https://www.youtube.com/watch?v=H3rqs1GKJ1c", "title": "Dj Healer - Great Escape [APW-2]", "description": "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ", "duration": 190, "embed": true}, {"uri": "https://www.youtube.com/watch?v=MEBS_IMPnrs", "title": "Dj Healer - 2 The Dark [APW-2]", "description": "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ", "duration": 507, "embed": true}, {"uri": "https://www.youtube.com/watch?v=Dt1LSsWWnL0", "title": "DJ Healer Gods Creation", "description": "\ua74e", "duration": 483, "embed": true}, {"uri": "https://www.youtube.com/watch?v=YGWRxr-8AEU", "title": "Dj Healer - Untitled [APW-2]", "description": "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ", "duration": 113, "embed": true}, {"uri": "https://www.youtube.com/watch?v=ceP9MQwwxsA", "title": "Dj Healer - Planet Lonely [APW-2]", "description": "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ", "duration": 402, "embed": true}, {"uri": "https://www.youtube.com/watch?v=XsnOVtR9BJA", "title": "Dj Healer - Hopes And Fears [APW-2]", "description": "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ", "duration": 372, "embed": true}, {"uri": "https://www.youtube.com/watch?v=23dPsHnBE5Y", "title": "Dj Healer - We Are Going Nowhere [APW-2]", "description": "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ", "duration": 728, "embed": true}, {"uri": "https://www.youtube.com/watch?v=_xD4ZVM_dRo", "title": "Dj Healer - Gone [APW-2]", "description": "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ", "duration": 225, "embed": true}, {"uri": "https://www.youtube.com/watch?v=UTMVF_H8sTs", "title": "Dj Healer - Protectionspell [APW-2]", "description": "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ", "duration": 489, "embed": true}]}`))
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	master, err := client.GetMasterRelease(context.Background(), 1521183)
	if err != nil {
		t.Fatal(err)
	}

	want := &MasterRelease{
		ID:                   1521183,
		MainRelease:          11849140,
		MostRecentRelease:    13416116,
		ResourceURL:          "https://api.discogs.com/masters/1521183",
		URI:                  "https://www.discogs.com/master/1521183-DJ-Healer-Nothing-2-Loose",
		VersionsURL:          "https://api.discogs.com/masters/1521183/versions",
		MainReleaseURL:       "https://api.discogs.com/releases/11849140",
		MostRecentReleaseURL: "https://api.discogs.com/releases/13416116",
		NumForSale:           19,
		LowestPrice:          151.52,
		Images: []Image{
			{
				Type:        "primary",
				URI:         "https://i.discogs.com/qAqSEZsVcMvvyMJfKXbv7KdmxY0HMvZLYx4eMt7kvPc/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTE2/MjItODA4NS5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/qAqSEZsVcMvvyMJfKXbv7KdmxY0HMvZLYx4eMt7kvPc/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTE2/MjItODA4NS5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/l71C5a4ersYHoN48qLm8KIpXSL9X8NngQgqDzKc0u1I/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTE2/MjItODA4NS5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/809DLP4xiHA0uT2TolWEwRonZpCAUDZCoLcnmCr4LnY/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTg4/NzMtOTA0Mi5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/809DLP4xiHA0uT2TolWEwRonZpCAUDZCoLcnmCr4LnY/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTg4/NzMtOTA0Mi5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/yMBLLOFr5r6VnDPaJtew1oLy8BB50UdgONP3VjeY7NI/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTg4/NzMtOTA0Mi5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/0ed3On5prgm7ZLmKx2rSqpXMlM1svDCkde1m1dZzYs0/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NDIz/MzItMzA5MC5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/0ed3On5prgm7ZLmKx2rSqpXMlM1svDCkde1m1dZzYs0/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NDIz/MzItMzA5MC5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/wnhuiqzldI-4hn7Rao3-pu-R5kIozka3F7ouFoSEaZw/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NDIz/MzItMzA5MC5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/k1z8DRjz5sqbTxaNHyvx9mV32NyANOENWS-rc3BPk2U/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjQ5MTE2/ODEtOTU0OS5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/k1z8DRjz5sqbTxaNHyvx9mV32NyANOENWS-rc3BPk2U/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjQ5MTE2/ODEtOTU0OS5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/FkjqlfBtGzN9kzx3ThUT0MoYoR6Afa7_AOIJUDww1PI/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjQ5MTE2/ODEtOTU0OS5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/a94M2BpoDjcExLXgDh6Bi32VozF5QkRPHbg5kNkd5ZQ/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NDIz/MjMtOTY4MS5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/a94M2BpoDjcExLXgDh6Bi32VozF5QkRPHbg5kNkd5ZQ/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NDIz/MjMtOTY4MS5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/qhTAIPPvjT_-MUCxdMLKLGCpO3vSakq3AjhmpFsUHaE/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NDIz/MjMtOTY4MS5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/uOPk_nMhbjWMpRfHG4D0ZtAFot8AZtCjmKd49X6f79c/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE1/OTUtMTcwOC5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/uOPk_nMhbjWMpRfHG4D0ZtAFot8AZtCjmKd49X6f79c/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE1/OTUtMTcwOC5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/cK3AiimTiQmFL4dJWvMeJVnMpQde8DRAW8YgFlsmlfU/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE1/OTUtMTcwOC5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/B0ajqnhT-hXS3SjzJmEOCyN8jpfNG7hdWApA9_4MsxU/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTEtNTQ1MS5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/B0ajqnhT-hXS3SjzJmEOCyN8jpfNG7hdWApA9_4MsxU/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTEtNTQ1MS5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/VVvhrf_deK_bjraHaSVuMaN5DR7T7R6UgY502TvL3Ok/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTEtNTQ1MS5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/meRonYJSZNlJtrBfB6MnvR-2L7jDgw0cQEcCUaNR2kU/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MDktOTA4NS5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/meRonYJSZNlJtrBfB6MnvR-2L7jDgw0cQEcCUaNR2kU/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MDktOTA4NS5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/pixDcsIIYZ5ajhbm6Q7zMt9cSbb-OosiWdZ6M5Qjx1k/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MDktOTA4NS5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/46C3fFL6V-WErSwPlS_edVnZMmW6FysZO6Y1BIY-3Ak/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTYtNDk3Ni5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/46C3fFL6V-WErSwPlS_edVnZMmW6FysZO6Y1BIY-3Ak/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTYtNDk3Ni5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/slARIJKnFPs3FLQ4YSX6ugmZmv_f16D8DKUHgfHQ2O0/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTYtNDk3Ni5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/a5a20FS60UjYfkgDqzrlqFxwvQVzgGAyR2MW0PZujbg/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTAtNDY5MS5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/a5a20FS60UjYfkgDqzrlqFxwvQVzgGAyR2MW0PZujbg/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTAtNDY5MS5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/Vv8BuIlMHXeAmdJJVGb8idLxLxEPYE_I9shCHeCGzZk/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTAtNDY5MS5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
			{
				Type:        "secondary",
				URI:         "https://i.discogs.com/LgGqid4ErvSC8Wn6yA0_zQXgB7CIrDrOLGvUq-DCBKc/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTYtNDQ2OS5qcGVn.jpeg",
				ResourceURL: "https://i.discogs.com/LgGqid4ErvSC8Wn6yA0_zQXgB7CIrDrOLGvUq-DCBKc/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTYtNDQ2OS5qcGVn.jpeg",
				URI150:      "https://i.discogs.com/mmbqWgflsRzZcUdM26OU1-uaCqoc7BxAQqOn2FJONxs/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NjE2/MTYtNDQ2OS5qcGVn.jpeg",
				Width:       600,
				Height:      600,
			},
		},
		Genres: []string{
			"Electronic",
		},
		Styles: []string{
			"Ambient",
			"Breakbeat",
			"House",
		},
		Year: 2018,
		Tracklist: []Track{
			{
				Position: "A1",
				Type:     "track",
				Title:    "At Last (Becalming The Storm)",
				Duration: "3:27",
			},
			{
				Position: "A2",
				Type:     "track",
				Title:    "Great Escape",
				Duration: "3:08",
			},
			{
				Position: "B",
				Type:     "track",
				Title:    "2 The Dark",
				Duration: "8:24",
			},
			{
				Position: "C1",
				Type:     "track",
				Title:    "Gods Creation",
				Duration: "8:01",
			},
			{
				Position: "C2",
				Type:     "track",
				Title:    "Untitled",
				Duration: "1:52",
			},
			{
				Position: "D1",
				Type:     "track",
				Title:    "Planet Lonely",
				Duration: "6:42",
			},
			{
				Position: "D2",
				Type:     "track",
				Title:    "Hopes And Fears",
				Duration: "6:11",
			},
			{
				Position: "E",
				Type:     "track",
				Title:    "We Are Going Nowhere",
				Duration: "12:08",
			},
			{
				Position: "F1",
				Type:     "track",
				Title:    "The Interview",
				Duration: "6:10",
			},
			{
				Position: "F2",
				Type:     "track",
				Title:    "Gone",
				Duration: "3:45",
			},
			{
				Position: "F3",
				Type:     "track",
				Title:    "Protectionspell",
				Duration: "8:07",
			},
		},
		Artists: []ReleaseArtist{
			{
				Name:         "DJ Healer",
				ANV:          "",
				Join:         "",
				Role:         "",
				Tracks:       "",
				ID:           6418741,
				ResourceURL:  "https://api.discogs.com/artists/6418741",
				ThumbnailURL: "https://i.discogs.com/yHVbfC3naoozshAMSvhInxFb7EsS03YGWBLuZvGLjbQ/rs:fit/g:sm/q:40/h:381/w:509/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTY0MTg3/NDEtMTUyMzg2OTEw/Mi0yMTUyLnBuZw.jpeg",
			},
		},
		Title:       "Nothing 2 Loose",
		DataQuality: "Correct",
		Videos: []Video{
			{
				URI:         "https://www.youtube.com/watch?v=8hB_WMlyDEM",
				Title:       "Dj Healer - At Last (Becalming The Storm) [APW-2]",
				Description: "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ",
				Duration:    208,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=n2fMoxSZZEU",
				Title:       "Dj Healer - The Interview [APW-2]",
				Description: "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ",
				Duration:    370,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=H3rqs1GKJ1c",
				Title:       "Dj Healer - Great Escape [APW-2]",
				Description: "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ",
				Duration:    190,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=MEBS_IMPnrs",
				Title:       "Dj Healer - 2 The Dark [APW-2]",
				Description: "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ",
				Duration:    507,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=Dt1LSsWWnL0",
				Title:       "DJ Healer Gods Creation",
				Description: "\ua74e",
				Duration:    483,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=YGWRxr-8AEU",
				Title:       "Dj Healer - Untitled [APW-2]",
				Description: "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ",
				Duration:    113,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=ceP9MQwwxsA",
				Title:       "Dj Healer - Planet Lonely [APW-2]",
				Description: "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ",
				Duration:    402,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=XsnOVtR9BJA",
				Title:       "Dj Healer - Hopes And Fears [APW-2]",
				Description: "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ",
				Duration:    372,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=23dPsHnBE5Y",
				Title:       "Dj Healer - We Are Going Nowhere [APW-2]",
				Description: "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ",
				Duration:    728,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=_xD4ZVM_dRo",
				Title:       "Dj Healer - Gone [APW-2]",
				Description: "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ",
				Duration:    225,
				Embed:       true,
			},
			{
				URI:         "https://www.youtube.com/watch?v=UTMVF_H8sTs",
				Title:       "Dj Healer - Protectionspell [APW-2]",
				Description: "BUY: https://allpossibleworlds.de/shop/\n\nDJ Healer \u200e\u2013 Nothing 2 Loose\nLabel: All Possible Worlds \u200e\u2013 APW-2\nFormat: 3 \u00d7 Vinyl, 12, Album \nCountry: Germany\nReleased: 01 April 2018\nGenre: Electronic\nStyle: Ambient, Breakbeat, House\n\nTracklist\nA1 At ",
				Duration:    489,
				Embed:       true,
			},
		},
	}

	testEqual(t, want, master)
}

func Test_GetMasterReleaseVersions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/masters/1521183/versions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		// NOTE: this response includes a "filters" property which is not represented in the MasterReleaseVersions struct at this time
		// there's no documented way to send filter parameters with a request
		_, _ = w.Write([]byte(`{"pagination": {"page": 1, "pages": 1, "per_page": 50, "items": 2, "urls": {}}, "filters": {"applied": {}, "available": {"format": {"12\"": 2, "33 \u2153 RPM": 2, "Album": 2, "Vinyl": 2, "Repress": 1}, "label": {"All Possible Worlds": 2}, "country": {"Germany": 2}, "released": {"2018": 1, "2019": 1}}}, "filter_facets": [{"title": "Format", "id": "format", "values": [{"title": "12\"", "value": "12%22", "count": 2}, {"title": "33 \u2153 RPM", "value": "33+%E2%85%93+RPM", "count": 2}, {"title": "Album", "value": "Album", "count": 2}, {"title": "Vinyl", "value": "Vinyl", "count": 2}, {"title": "Repress", "value": "Repress", "count": 1}], "allows_multiple_values": true}, {"title": "Label", "id": "label", "values": [{"title": "All Possible Worlds", "value": "All+Possible+Worlds", "count": 2}], "allows_multiple_values": true}, {"title": "Country", "id": "country", "values": [{"title": "Germany", "value": "Germany", "count": 2}], "allows_multiple_values": true}, {"title": "Released", "id": "released", "values": [{"title": "2018", "value": "2018", "count": 1}, {"title": "2019", "value": "2019", "count": 1}], "allows_multiple_values": true}], "versions": [{"id": 11849140, "label": "All Possible Worlds", "country": "Germany", "title": "Nothing 2 Loose", "major_formats": ["Vinyl"], "format": "12\", 33 \u2153 RPM, Album", "catno": "APW-2", "released": "2018", "status": "Accepted", "resource_url": "https://api.discogs.com/releases/11849140", "thumb": "https://i.discogs.com/l71C5a4ersYHoN48qLm8KIpXSL9X8NngQgqDzKc0u1I/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExODQ5/MTQwLTE1MjM0NTE2/MjItODA4NS5qcGVn.jpeg", "stats": {"community": {"in_wantlist": 1839, "in_collection": 1368}, "user": {"in_wantlist": 0, "in_collection": 0}}}, {"id": 13416116, "label": "All Possible Worlds", "country": "Germany", "title": "Nothing 2 Loose", "major_formats": ["Vinyl"], "format": "12\", 33 \u2153 RPM, Album, Repress", "catno": "APW-2", "released": "2019", "status": "Accepted", "resource_url": "https://api.discogs.com/releases/13416116", "thumb": "https://i.discogs.com/o73HPuPE1awkg3sa0VTQQmYbFo7DOLMq-SIzEBGPoBQ/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEzNDE2/MTE2LTE1NTM3ODk4/ODItMzEwNi5qcGVn.jpeg", "stats": {"community": {"in_wantlist": 657, "in_collection": 482}, "user": {"in_wantlist": 0, "in_collection": 0}}}]}`))
	})

	client := newTestClient(server.URL, testUserAgent, auth.NewToken(testAuthToken))

	versions, err := client.GetMasterReleaseVersions(context.Background(), 1521183, nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &MasterReleaseVersions{
		Pagination: Pagination{
			Page:    1,
			Pages:   1,
			PerPage: 50,
			Items:   2,
			URLs:    PaginationURLs{},
		},
		Versions: []MasterReleaseVersion{
			{
				ID:      11849140,
				Label:   "All Possible Worlds",
				Country: "Germany",
				Title:   "Nothing 2 Loose",
				MajorFormats: []string{
					"Vinyl",
				},
				Format:      "12\", 33 \u2153 RPM, Album",
				CatNo:       "APW-2",
				Released:    "2018",
				Status:      "Accepted",
				ResourceURL: "https://api.discogs.com/releases/11849140",
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
				ID:      13416116,
				Label:   "All Possible Worlds",
				Country: "Germany",
				Title:   "Nothing 2 Loose",
				MajorFormats: []string{
					"Vinyl",
				},
				Format:      "12\", 33 \u2153 RPM, Album, Repress",
				CatNo:       "APW-2",
				Released:    "2019",
				Status:      "Accepted",
				ResourceURL: "https://api.discogs.com/releases/13416116",
				Thumb:       "https://i.discogs.com/o73HPuPE1awkg3sa0VTQQmYbFo7DOLMq-SIzEBGPoBQ/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEzNDE2/MTE2LTE1NTM3ODk4/ODItMzEwNi5qcGVn.jpeg",
				Stats: Stats{
					Community: CommunityStats{
						InWantlist:   657,
						InCollection: 482,
					},
					User: UserStats{
						InWantlist:   0,
						InCollection: 0,
					},
				},
			},
		},
	}

	testEqual(t, want, versions)
}

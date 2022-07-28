package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/peacefixation/go-discogs/discogs"
	"github.com/peacefixation/go-discogs/discogs/auth"
)

func main() {
	apiToken := "your api token"
	userAgent := "GoDiscogs/0.0.1 +http://yournamehere.net"

	client := discogs.NewClient(discogs.BaseURL, userAgent, auth.NewToken(apiToken))

	resp, err := client.Search(context.Background(), "Nirvana", nil)
	if err != nil {
		log.Fatal(err)
	}

	print(resp)
}

func print(data any) {
	marshalled, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(marshalled))
	}
}

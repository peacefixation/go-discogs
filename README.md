# go-discogs

A Go client for the [Discogs API](https://www.discogs.com/developers/).

This is a work in progress. I will implement parts of the API in order of usefulness to me, which at present looks like this:

- Database (&check;)  
- User Collection (&check;)  
- User Wantlist  
- List  
- Inventory Export  
- Inventory Upload  
- User Identity 
- Marketplace  
- OAuth  

There is a GraphQL API (Discogs API v3) on the horizon that should make some things a lot easiser but I have no idea when it will be available.

## Authentication

The Discogs API requires authentication for some requests. Using authentication will also increase the number of requests you can make per minute. See the Discogs [documentation](https://www.discogs.com/developers/#page:authentication) for more information.

## Acknowledgements

I took some inspiration from Scott McAllister's [go-discogs](https://github.com/stmcallister/go-discogs) project, particularly for testing the API client with a mock server.

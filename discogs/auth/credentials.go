package auth

import "fmt"

type Credentials struct {
	key    string
	secret string
}

func NewCredentials(key, secret string) *Credentials {
	return &Credentials{
		key:    key,
		secret: secret,
	}
}

func (c *Credentials) String() string {
	return fmt.Sprintf("Discogs key=%s, secret=%s", c.key, c.secret)
}

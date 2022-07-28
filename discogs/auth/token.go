package auth

import "fmt"

type Token struct {
	Token string
}

func NewToken(token string) *Token {
	return &Token{
		Token: token,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("Discogs token=%s", t.Token)
}

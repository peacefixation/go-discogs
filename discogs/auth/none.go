package auth

type None struct{}

func NewNone() *None {
	return &None{}
}

func (c *None) String() string {
	return ""
}

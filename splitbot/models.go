package splitbot

type BorderDirection int

const (
	Up BorderDirection = iota
	Down
)

func (d BorderDirection) String() string {
	return []string{
		"⬆⬆⬆⬆⬆",
		"⬇⬇⬇⬇⬇",
	}[d]
}

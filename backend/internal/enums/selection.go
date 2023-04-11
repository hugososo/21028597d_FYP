package enums

type Selection uint8

const (
	DISAGREE = iota
	AGREE
)

func (s Selection) String() string {
	return [...]string{"DISAGREE", "AGREE"}[s]
}

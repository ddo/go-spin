package spin

const (
	DEFAULT_SPRITE = `-\|/`
)

type Spinner struct {
	index  int
	sprite string
}

func New(sprite string) *Spinner {
	if sprite == "" {
		sprite = DEFAULT_SPRITE
	}

	return &Spinner{
		index:  -1,
		sprite: sprite,
	}
}

func (s *Spinner) Spin() string {
	s.index++
	return string(s.sprite[s.index%len(s.sprite)])
}

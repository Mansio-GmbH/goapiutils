package ct

type Distance int64

const (
	Meter     Distance = 1
	Kilometer Distance = 1000
)

func (d Distance) Meters() int64 {
	return int64(d)
}

func (d Distance) Kilometers() int64 {
	return int64(d) / 1000
}

func (d Distance) Truncate(m Distance) Distance {
	if m <= 0 {
		return d
	}
	return d - d%m
}

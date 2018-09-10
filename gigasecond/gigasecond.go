package gigasecond

import "time"

const (
	gigasecond = time.Duration(1e9) * time.Second
)

// AddGigasecond will add 1,000,000,000 seconds to any time it is given
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}

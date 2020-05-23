package gigasecond

import "time"

// AddGigasecond compute the gigasecond later of a given time.
func AddGigasecond(t time.Time) time.Time {
	// Add time.Second * 1e9 to the given time then return.
	return t.Add(time.Second * 1e9)
}

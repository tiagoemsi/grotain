package grotain

import "math"

// Cards is the basic structure for reviewing an item
type Cards struct {
	Interval int
	Efactor  float64
}

// CalcInterval calculates the next revision in days
func (c Cards) CalcInterval() int {
	// for n>2 I(n):=I(n-1)*EF
	switch c.Interval {
	case 0:
		return 1
	case 1:
		return 6 // 3?
	default:
		return int(math.Round(float64((c.Interval - 1)) * c.Efactor))
	}
}

// CalcEfactor calculates the new easiness factor reflecting the easiness of memorizing and retaining a given item in memory
func (c Cards) CalcEfactor(quality int) float64 {
	// EF':=EF+(0.1-(5-q)*(0.08+(5-q)*0.02))
	c.Efactor = c.Efactor + (0.1 - float64(5-quality)*(0.08+float64(5-quality)*float64(0.02)))
	if c.Efactor < 1.3 {
		return 1.3
	}

	return c.Efactor
}

// Review ranks with a score between 0 and 5 and returns the new range and the new Efactor to use in the next revision
func (c *Cards) Review(quality int) *Cards {
	// quality = [0 ~ 5]

	c.Efactor = c.CalcEfactor(quality)

	if quality < 3 {
		c.Interval = 1
		// c.interval = int(math.Round(float64((c.interval / 2))))
	} else {
		c.Interval = c.CalcInterval()
	}

	// Store this data for use in the next review
	return &Cards{c.Interval, c.Efactor}
}

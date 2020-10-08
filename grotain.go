package grotain

import "math"

const (
	minimumEaseFactor           = 1.3
	firstInterval               = 0
	newFirstInterval            = 1
	secondInterval              = newFirstInterval
	newSecondInterval           = 6
	wrongResponse               = 3
	newIntervalAfterWrongAnswer = newFirstInterval
	oneDayAgo                   = 1
)

// Cards is the basic structure for reviewing an item
type Cards struct {
	Interval int
	Efactor  float64
}

// CalcInterval calculates the next revision in days
func (c Cards) CalcInterval(quality int) int {
	if quality < wrongResponse {
		return newIntervalAfterWrongAnswer
	}

	switch c.Interval {
	case firstInterval:
		return newFirstInterval
	case secondInterval:
		return newSecondInterval
	default:
		return int(math.Round(float64((c.Interval - oneDayAgo)) * c.Efactor))
	}
}

// CalcEfactor calculates the new easiness factor reflecting the easiness of memorizing and retaining a given item in memory
func (c Cards) CalcEfactor(quality int) float64 {
	c.Efactor = c.Efactor + (0.1 - float64(5-quality)*(0.08+float64(5-quality)*float64(0.02)))
	if c.Efactor < minimumEaseFactor {
		return minimumEaseFactor
	}

	return c.Efactor
}

// Review ranks with a score between 0 and 5 and returns the new range and the new Efactor to use in the next revision
func (c *Cards) Review(quality int) *Cards {
	c.Efactor = c.CalcEfactor(quality)
	c.Interval = c.CalcInterval(quality)

	// Store this data for use in the next review
	return &Cards{c.Interval, c.Efactor}
}

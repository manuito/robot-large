package pilote

import (
	"robotoby/application"
	"strconv"
)

/*
 * Robot modules can send back crappy values.
 * The sensors are mostly low cost components, and the links are not perfect.
 *
 * For example, HS distance sensors can return values like
 * 1, 2, 2, 1, 2, 1200, 2, 1, 2, 999, 2, 1 ...
 * So the "large values" should be ignored, "debounced" when too large regarding the AVG
 */

type debouncedValue struct {
	values []int
}

// avgDeb : Get the average value in debouncedValue, cleaned
func (deb *debouncedValue) avgDeb() int {

	var proc, sum int

	for _, cur := range deb.values {

		if cur > -1 {
			sum += cur
			proc++
		}
	}

	// If nothing yet, use -1
	if proc == 0 {
		return -1
	}

	// Else get clean AVG
	return sum / proc
}

// storeIfClean : Regarding debouncing
func (deb *debouncedValue) storeIfClean(value int) {
	max := application.State.Config.DebounceBufferSize
	maxVar := application.State.Config.DebounceMaxVariation
	// Save immediately + init debounc fixed size, -1 values

	if len(deb.values) == 0 {
		deb.values = append(deb.values, value)
		for index := 1; index < max; index++ {
			deb.values = append(deb.values, -1)
		}
	} else { // Check with debounce
		avg := deb.avgDeb()
		if value <= maxVar*avg {
			// Check ok, add to top of debounced values
			for index := max - 1; index > 0; index-- {
				deb.values[index] = deb.values[index-1]
			}
			deb.values[0] = value
		} else {
			application.Debug("Found var to debounce : current AVG = " + strconv.Itoa(avg) + " VAR = " + strconv.Itoa(maxVar) + " for VALUE = " + strconv.Itoa(value))
		}
	}
}

// getCurrent : Current debounced value
func (deb *debouncedValue) getCurrent() int {
	if len(deb.values) > 0 {
		return deb.values[0]
	}
	return -1
}

package util

// ValidateDate returns false if the given calendar date
// is out of the normalized range.
func ValidateDate(year, month, day int32) bool {
	if (year < 0) || (month > 12 || month < 1) || (day > 31 || day < 1) {
		return false
	}
	return true
}

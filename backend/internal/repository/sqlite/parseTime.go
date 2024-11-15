package sqlite

import "time"

const format = time.RFC3339

// format time to the string format
func timeToString(t time.Time) string {
	return t.Format(format)
}

// will parse the time in the format saved to the database
func parseTime(tString string) (time.Time, error) {
	return time.Parse(format, tString)
}

package utils

import "time"

//StrIsEmpty is an helper function used check if a string is nil or empty
func StrIsEmpty(s *string) bool {
	return s == nil || *s == ``
}

//TimeIsEmpty is an helper function used check if a time is nil or empty
func TimeIsEmpty(t *time.Time) bool {
	return t == nil || (*t).IsZero()
}

package utils

import (
	"time"
)

//StrToPtr is an helper function used get the ptr of a string
func StrToPtr(s string) *string {
	return &s
}

//ArrStrToPtr is an helper function used get the ptr of an array of string
func ArrStrToPtr(aS []string) []*string {
	ret := []*string{}
	for _, a := range aS {
		ret = append(ret, StrToPtr(a))
	}
	return ret
}

//BoolToPtr is an helper function used get the ptr of a bool
func BoolToPtr(b bool) *bool {
	return &b
}

//IntToPtr is an helper function used get the ptr of an int
func IntToPtr(i int) *int {
	return &i
}

//Int64ToPtr is an helper function used get the ptr of an int64
func Int64ToPtr(i int64) *int64 {
	return &i
}

//Uint64ToPtr is an helper function used get the ptr of an uint64
func Uint64ToPtr(u uint64) *uint64 {
	return &u
}

//TimeToPtr is an helper function used get the ptr of a time
func TimeToPtr(t time.Time) *time.Time {
	return &t
}

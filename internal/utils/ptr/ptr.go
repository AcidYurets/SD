package ptr

import "time"

func String(s string) *string {
	return &s
}

func Bool(b bool) *bool {
	return &b
}

func Time(t time.Time) *time.Time {
	return &t
}

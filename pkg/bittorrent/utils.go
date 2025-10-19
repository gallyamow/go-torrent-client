package bittorrent

import "fmt"

func Ptr[T any](val T) *T {
	return &val
}

func StringifyPtr[T any](val *T) string {
	if val == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", *val)
}

package random

import (
	"math/rand"
	"time"
)

func Bool() bool {
	return rand.Intn(2) == 1
}

func IntRange(min, max int) int {
	return min + rand.Intn(max-min)
}

func TimestampRange(min, max time.Time) time.Time {
	return min.Add(time.Duration(rand.Int63n(max.Unix()-min.Unix())) * time.Second)
}

func FromSlice[T any](slice []T) T {
	return slice[rand.Intn(len(slice))]
}

func FromSliceWithRemove[T any](slice []T) (T, []T) {
	index := rand.Intn(len(slice))
	res := slice[index]
	slice = append(slice[:index], slice[index+1:]...)
	return res, slice
}

func NFromSlice[T any](slice []T, n int) []T {
	if n > len(slice) {
		n = len(slice)
	}
	res := make([]T, n)
	for i := 0; i < n; i++ {
		res[i] = slice[rand.Intn(len(slice))]
	}
	return res
}

package utils

import "math/rand"

func GetRandomIntInRange(min int, max int) int {
	return rand.Intn(max-min) + min
}

func GetRandomElementOfSlice[T any](slice []T) T {
	return slice[rand.Intn(len(slice))]
}

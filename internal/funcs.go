package internal

import "log"

func Must[T any](v T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return v
}

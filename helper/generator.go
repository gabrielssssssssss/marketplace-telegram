package helper

import "math/rand/v2"

func RandomString(size int) string {
	var char = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var randomString string

	for i := 0; i < size; i++ {
		seeds := rand.NewPCG(0, 24)
		randomString += char[rand.New(seeds).Uint64()]
	}

	return randomString
}

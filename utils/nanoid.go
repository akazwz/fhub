package utils

import gonanoid "github.com/matoous/go-nanoid/v2"

var alphabet = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

func GenerateID(size int) string {
	id, err := gonanoid.Generate(alphabet, size)
	if err != nil {
		GenerateID(size)
	}
	return id
}

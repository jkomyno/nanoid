package main

import (
	"crypto/rand"
	"math"
	"fmt"
)

type RandomType func(int) ([]byte, error)

type defaultsType struct {
	Alphabet string
	Size int
	MaskSize int
}

func getDefaults() *defaultsType {
	return &defaultsType{
		Alphabet: "_~0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", // len=64
		Size: 22,
		MaskSize: 5,
	}
}

var defaults = getDefaults()

func initMasks(params ...int) []uint {
	var size int
	if len(params) == 0 {
		size = defaults.MaskSize
	} else {
		size = params[0]
	}
	/*
	https://github.com/ai/nanoid/blob/d6ad3412147fa4c2b0d404841ade245a00c2009f/format.js#L1
	As per 'var masks = [15, 31, 63, 127, 255]'

	The next block initializes an array of size elements, from 2^4-1 to 2^(3 + size)-1
	*/
	masks := make([]uint, size)
	for i := 0; i < size; i++ {
		shift := 3 + i
		masks[i] = (2 << uint(shift)) - 1
	}
	return masks
}

/*
https://github.com/ai/nanoid/blob/d6ad3412147fa4c2b0d404841ade245a00c2009f/format.js#L29-L31
var mask = masks.find(function (i) {
	return i >= alphabet.length - 1
})
*/
func getMask(alphabet string, masks []uint) int {
	for i := 0; i < len(masks); i++ {
		curr := int(masks[i])
		if curr >= len(alphabet) - 1 {
			return curr
		}
	}
	return 0
}

// Random generates cryptographically strong pseudo-random data.
// The size argument is a number indicating the number of bytes to generate.
func Random(size int) ([]byte, error) {
	var randomBytes = make([]byte, size)
	_, err := rand.Read(randomBytes)
	return randomBytes, err
}

// Format returns a secure random string with custom random generator and alphabet.
// `size` is the number of symbols in new random string
func Format(random RandomType, alphabet string, size int) (string, error) {
	masks := initMasks(size)
	mask := getMask(alphabet, masks)
	ceilArg := 1.6 * float64(mask * size) / float64(len(alphabet))
	step := int(math.Ceil(ceilArg))

	id := make([]byte, size)
	for j := 0;; {
		bytes, err := random(step)
		if (err != nil) {
			return "", err
		}

		for i := 0; i < step; i++ {
			currByte := bytes[i] & byte(mask)
			if currByte < byte(len(alphabet)) {
				id[j] = alphabet[currByte]
				j++
				if j == size {
					fmt.Println("id", id)
					return string(id[:size]), nil
				}
			}
		}
	}
}

// Generate is a low-level function to change alphabet and ID size.
func Generate(alphabet string, size int) (string, error) {
	return Format(Random, alphabet, size);
}

// Nanoid generates secure URL-friendly unique ID.
func Nanoid(size int) (string, error) {
	if size == 0 {
		size = defaults.Size
	}
	bytes, err := Random(size)
	if err != nil {
		return "", err
	}

	id := make([]byte, size)
	for i := 0; i < size; i++ {
		id[i] = defaults.Alphabet[bytes[i] & 63]
	}
	return string(id[:size]), nil
}

func main() {
	str, _ := Nanoid(5)
	fmt.Println("str %v", str);
	/*
	str, err := Generate(defaults.Alphabet, defaults.Size)
	if err != nil {
		panic(err)
	}
	fmt.Println("str %v", str);
	*/
}
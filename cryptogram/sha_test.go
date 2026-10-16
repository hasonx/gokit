package cryptogram

import (
	"log"
	"testing"
)

func TestSha256(t *testing.T) {

	res := Sha256Encrypt("mqcfhrcloud1541497912")
	if res == "c697d0f31d808d0b7aa092f100365080670f9f3b292f96b0beb9e13109073e9e" {
		log.Println("equal")
	}

	log.Print(res)
}

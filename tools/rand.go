package tools

import (
	"crypto/rand"
	"fmt"
)

// RandToken RandToken
func RandToken() string {
	b := make([]byte, 15)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

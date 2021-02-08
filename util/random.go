package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomAuthor generates a random author name
func RandomAuthor() string {
	return fmt.Sprintf("{} {}", RandomString(6), RandomString(10))
}

// RandomTitle generates a random book title
func RandomTitle() string {
	return fmt.Sprintf("{} {} {}", RandomString(6), RandomString(10), RandomString(4))
}

// RandomEmail generates a random email address
func RandomEmail() string {
	return fmt.Sprintf("{}@{}.{}", RandomString(8), RandomString(6), RandomString(3))
}

// RandomPhone generates a random phone number
func RandomPhone() string {
	return fmt.Sprintf("{}", RandomInt(10000000000, 99999999999))
}

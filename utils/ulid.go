package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

func GenerateULID() string {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(time.Now()), entropy)
	return strings.ToLower(id.String())
}

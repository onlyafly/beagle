package garden

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

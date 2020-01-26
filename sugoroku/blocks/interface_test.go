package blocks

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	count := 10
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	var m = make(map[int]Block)
	for i := 0; i < count; i++ {
		m[i] = New(BlockType(r.Intn(2)))
	}

	for _, v := range m {
		fmt.Println(v.String())
	}
}

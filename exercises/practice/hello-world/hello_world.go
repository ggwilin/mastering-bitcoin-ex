package hash

import (
	"fmt"
	"hash/fnv"
)

func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	fmt.Println(h.Sum32())
	return h.Sum32()
}

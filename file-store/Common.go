package fs

import "github.com/spaolacci/murmur3"

type Result struct {
	Success bool   `json:"success"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

const (
	mask            = 0x7FF
	murmur3HashSeed = 104729
)

func hashFileKey(fileKey string) int64 {
	h, _ := murmur3.Sum128WithSeed([]byte(fileKey), murmur3HashSeed)
	return int64(h) & mask
}

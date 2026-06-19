package helpers

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func ExtractPrefix(path string) string {
	path = strings.TrimPrefix(path, "/api")

	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	if len(parts) > 0 && parts[0] != "" {
		return "/" + parts[0]
	}
	return ""
}


func GenerateRequestID() string {
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Int63())
}
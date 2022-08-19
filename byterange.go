package groupcache

import (
	"fmt"
	"strings"
)

func RangeKey(key string, start, end int64) string {
	key = strings.ReplaceAll(key, "$", "\\$")
	return fmt.Sprintf("%s$$%d-%d", key, start, end)
}

func KeyToRange(rk string) (string, int64, int64, bool) {
	var key string
	var start, end int64
	if n, err := fmt.Sscanf(rk, "%s$$%d-%d", &key, &start, &end); err != nil && n == 3 {
		key = strings.ReplaceAll(key, "\\$", "$")
		return key, start, end, true
	}
	return rk, 0, 0, false
}

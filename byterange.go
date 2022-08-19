package groupcache

import "fmt"

func RangeKey(key string, start, end int64) string {
	return fmt.Sprintf("%s$%d-%d", key, start, end)
}

func KeyToRange(rk string) (string, int, int, bool) {
	var key string
	var start, end int
	if n, err := fmt.Sscanf(key, "%s$%d-%d", &key, &start, &end); err != nil && n == 3 {
		return key, start, end, true
	}
	return rk, 0, 0, false
}

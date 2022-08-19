package groupcache

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`^(.*)##(\d+):(\d+)$`)

func RangeKey(key string, start, end int64) string {
	key = strings.ReplaceAll(key, `#`, `\#`)
	return fmt.Sprintf("%s##%d:%d", key, start, end)
}

func KeyToRange(rk string) (string, int64, int64, bool) {
	match := re.FindStringSubmatch(rk)
	if len(match) == 4 {
		key := strings.ReplaceAll(match[1], `\#`, `#`)
		start, err1 := strconv.ParseInt(match[2], 10, 64)
		end, err2 := strconv.ParseInt(match[3], 10, 64)
		if err1 == nil && err2 == nil {
			return key, start, end, true
		}
	}
	return rk, 0, 0, false
}

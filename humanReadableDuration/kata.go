package kata

import (
	"fmt"
	"strings"
)

func FormatDuration(seconds int64) string {
	if seconds <= 0 {
		return "now"
	}
	result := ""
	sortedUnits := []string{"year", "day", "hour", "minute", "second"}
	unitsInSeconds := map[string]int64{
		"year":   31_536_000,
		"day":    86_400,
		"hour":   3_600,
		"minute": 60,
		"second": 1,
	}

	for _, k := range sortedUnits {
		v := unitsInSeconds[k]
		if add := seconds / v; add > 0 {
			seconds -= (v * add)
			result += fmt.Sprintf("%v %v", add, k)
			if add > 1 {
				result += "s"
			}
			result += ", "
		}
	}

	result = strings.TrimRight(strings.TrimRight(result, " "), ",")
	idx := strings.LastIndex(result, ",")
	if idx > 0 {
		result = fmt.Sprintf("%v and%v", result[:idx], result[idx+1:])
	}
	return result
}

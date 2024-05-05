package replacement

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func sprintf(str string, args ...interface{}) string {
	return fmt.Sprintf(str, args...)
}

func getIntPercentage(value, max int) int {
	return (100*value + max/2) / max
}

func getIntValueOfPercent(total, percent int) int {
	return (total*percent + 50) / 100
}

func formatUsageAndDurationMaps(usages map[string]int, dmap map[string]time.Duration) string {
	keys := make([]string, 0)
	for k, _ := range dmap {
		keys = append(keys, k)
	}
	// sort.Strings(keys)
	sort.Slice(keys, func(i, j int) bool {
		return usages[keys[i]] > usages[keys[j]]
	})
	res := fmt.Sprintf(" %-13s: TOTAL USES  WORST TIME\n", "RULE NAME")
	for _, k := range keys {
		res += fmt.Sprintf("  %-12s: %-10d  %-8v\n", k, usages[k], dmap[k])
	}
	return res
}

func progressBarCLI(title string, value, endvalue, bar_length int) { // because I can
	endvalue -= 1
	percent := float64(value) / float64(endvalue)
	arrow := ">"
	for i := 0; i < int(percent*float64(bar_length)); i++ {
		arrow = "-" + arrow
	}
	spaces := strings.Repeat(" ", bar_length-len(arrow)+1)
	percent_with_dec := fmt.Sprintf("%.1f", percent*100.0)
	fmt.Printf("\r%s [%s%s] %s%% (%d out of %d)", title, arrow, spaces, percent_with_dec, value, endvalue)
	if value == endvalue {
		fmt.Printf("\n")
	}
}

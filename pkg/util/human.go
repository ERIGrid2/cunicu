package util

import (
	"fmt"
	"strings"
	"time"

	t "github.com/stv0g/cunicu/pkg/util/terminal"
)

const (
	day  = 24 * time.Hour
	year = 365 * day
)

// PrettyDuration pretty prints a time just like `wg show`
// See: https://github.com/WireGuard/wireguard-tools/blob/71799a8f6d1450b63071a21cad6ed434b348d3d5/src/show.c#L129
func PrettyDuration(left time.Duration) string {
	out := []string{}
	comps := []struct {
		name    string
		divisor time.Duration
	}{
		{"year", year},
		{"day", day},
		{"hour", time.Hour},
		{"minute", time.Minute},
		{"second", time.Second},
	}

	for _, comp := range comps {
		num := left / comp.divisor

		if num > 0 {
			left -= num * comp.divisor

			unit := comp.name
			if num > 1 {
				unit += "s" // plural s
			}

			out = append(out, fmt.Sprintf("%d "+t.Color("%s", t.FgCyan), num, unit))

		}
	}

	return strings.Join(out, ", ")
}

// Ago pretty prints a duration with an `ago` suffix.
// See: https://github.com/WireGuard/wireguard-tools/blob/71799a8f6d1450b63071a21cad6ed434b348d3d5/src/show.c#L157
func Ago(ts time.Time) string {
	d := time.Since(ts)

	if d < time.Second {
		return "Now"
	}

	return PrettyDuration(d) + " ago"
}

// Every pretty prints a duration with an `every` prefix
// See: https://github.com/WireGuard/wireguard-tools/blob/71799a8f6d1450b63071a21cad6ed434b348d3d5/src/show.c#L176
func Every(d time.Duration) string {
	return "every " + PrettyDuration(d)
}

// PrettyBytes pretty prints a byte count
// See: https://github.com/WireGuard/wireguard-tools/blob/71799a8f6d1450b63071a21cad6ed434b348d3d5/src/show.c#L184
func PrettyBytes(b int64) string {
	if b < 1024 {
		return fmt.Sprintf("%d "+t.Color("B", t.FgCyan), b)
	}

	var suffices = []rune{'K', 'M', 'G', 'T', 'P', 'E', 'Z', 'Y'}
	var f float32
	var i int

	for i, f = 0, float32(b); i < len(suffices) && f >= 1024; i, f = i+1, f/1024 {
	}

	if i == 0 {
		return fmt.Sprintf("%.0f "+t.Color("B", t.FgCyan), f)
	} else {
		return fmt.Sprintf("%.2f "+t.Color("%ciB", t.FgCyan), f, suffices[i-1])
	}
}

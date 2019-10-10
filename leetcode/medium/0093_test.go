package medium

import (
	"sort"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		input string
		want  []string
	}{
		{"25525511135", []string{"255.255.111.35", "255.255.11.135"}},
		{"1111", []string{"1.1.1.1"}},
		{"0111", []string{"0.1.1.1"}},
		{"11101", []string{"11.1.0.1", "1.11.0.1", "1.1.10.1"}},
		{"255255111135", []string{"255.255.111.135"}},
		{"255255111335", []string{}},
		{"25525511335", []string{"255.255.113.35"}},
		{"010010", []string{"0.100.1.0", "0.10.0.10"}},
		{"010010", []string{"0.100.1.0", "0.10.0.10"}},
		{"", []string{}},
	}
	for _, c := range tests {
		var got = restoreIpAddresses(c.input)
		sort.Strings(got)
		sort.Strings(c.want)
		if len(got) != len(c.want) {
			t.Errorf("restoreIpAddresses(%q) == %q, want %q", c.input, got, c.want)
		}
		for idx, val := range got {
			if c.want[idx] != val {
				t.Errorf("restoreIpAddresses(%q) == %q, want %q", c.input, got, c.want)
			}
		}
	}

}

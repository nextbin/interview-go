package medium

import "testing"

func Test_validIPAddress(t *testing.T) {
	type args struct {
		IP string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1",args{"172.16.254.1"},"IPv4"},
		{"case2",args{"2001:0db8:85a3:0:0:8A2E:0370:7334"},"IPv6"},
		{"case3",args{"256.256.256.256"},"Neither"},

		{"case4",args{"01.01.01.01"},"Neither"},
		{"case5",args{"02001:0db8:85a3:0:0:8A2E:0370:7334"},"IPv6"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validIPAddress(tt.args.IP); got != tt.want {
				t.Errorf("validIPAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

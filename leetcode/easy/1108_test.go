package easy

import "testing"

func Test_defangIPaddr(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1",args{"1.1.1.1"},"1[.]1[.]1[.]1"},
		{"case1",args{"255.100.50.0"},"255[.]100[.]50[.]0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defangIPaddr(tt.args.address); got != tt.want {
				t.Errorf("defangIPaddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

package kata_test

import (
	"testing"

	kata "github.com/tobibot/katasGo/stringsMix"
)

func TestMix(t *testing.T) {
	tests := []struct {
		name string
		s1 string
		s2 string
		want string
	}{
		{"here?", "Are they here", "yes, they are here", "2:eeeee/2:yy/=:hh/=:rr"},
		{"some u's", "uuuuuu", "uuuuuu", "=:uuuuuu"},
		{"looping", "looping is fun but dangerous", "less dangerous than coding", "1:ooo/1:uuu/2:sss/=:nnn/1:ii/2:aa/2:dd/2:ee/=:gg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kata.Mix(tt.s1, tt.s2); got != tt.want {
				t.Errorf("Mix() = %v, want %v", got, tt.want)
			}
		})
	}
}

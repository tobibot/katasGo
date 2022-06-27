// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tobibot/katasGo/humanReadableDuration"
)

var _ = Describe("Test Example", func() {
	It("Test small static inputs", func() {
		Expect(FormatDuration(1)).To(Equal("1 second"))
		Expect(FormatDuration(62)).To(Equal("1 minute and 2 seconds"))
		Expect(FormatDuration(120)).To(Equal("2 minutes"))
		Expect(FormatDuration(3600)).To(Equal("1 hour"))
		Expect(FormatDuration(3662)).To(Equal("1 hour, 1 minute and 2 seconds"))
	})
})

func TestFormatDuration(t *testing.T) {	
	tests := []struct {
		name string
		seconds int64
		want string
	}{
		// {"1 second", 1, "1 second" },
		//{"1 minute and 2 seconds", 62, "1 minute and 2 seconds" },
		// {"2 minutes", 120, "2 minutes" },
		//{"1 hour", 3600, "1 hour" },
		{"1 hour, 1 minute and 2 seconds", 3662, "1 hour, 1 minute and 2 seconds" },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDuration(tt.seconds); got != tt.want {
				t.Errorf("FormatDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

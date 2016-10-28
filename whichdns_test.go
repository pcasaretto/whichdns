package whichdns_test

import (
	"fmt"
	"testing"

	"github.com/pcasaretto/whichdns"
)

func TestDefaultLookuper(t *testing.T) {
	var lookuper whichdns.Lookuper

	t.Run("DefaultLookuper conforms to Lookuper", func(t *testing.T) {
		lookuper = &whichdns.DefaultLookupper{}
	})

	var tests = []struct {
		input  string
		output string
	}{
		{"kurum.in", "DomainControl"},
		{"terra.com.br", "Terra"},
		{"kinghost.com.br", "Kinghost"},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("given %s, it outputs %s", test.input, test.output), func(t *testing.T) {
			if actual, expected := lookuper.Lookup(test.input), test.output; actual != expected {
				t.Errorf("Expected %s, got %s", expected, actual)
			}
		})
	}
}

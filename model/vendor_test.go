package model

import (
	"testing"
)

func TestVendor_slug(t *testing.T) {
	tests := []struct {
		name string
		v    *Vendor
		want string
	}{
		// TODO: Add test cases.
		{"Simple", &Vendor{Code: "ABC", Name: "Se aMe"}, "abc-se-ame"},
		{"TheGerman", &Vendor{Code: "ÄBC", Name: "Se äme"}, "abc-se-ame"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.slug(); got != tt.want {
				t.Errorf("Vendor.slug() = %v, want %v", got, tt.want)
			}
		})
	}
}

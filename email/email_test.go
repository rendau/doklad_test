package email

import "testing"

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email string
		valid bool
	}{
		{"test@example.com", true},
		{"invalid-email", false},
		{"test@.com", false},
		{"test@domain", false},
		{"test@domain.com", true},
		//{"test@domain.@.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			if got := IsValidEmail(tt.email); got != tt.valid {
				t.Errorf("IsValidEmail(%v) = %v; want %v", tt.email, got, tt.valid)
			}
		})
	}
}

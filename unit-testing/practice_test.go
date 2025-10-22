package main

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	tests := []struct {
		name   string
		amount int
		want   int
	}{
		{"positive", 100, 100},
		{"negative", -100, -100},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			d := &Account{Balance: 0}
			d.Deposit(tt.amount)
			if d.Balance != tt.want {
				t.Errorf("got: %d, expected: %d", tt.amount, tt.want)
			}

		})
	}
}

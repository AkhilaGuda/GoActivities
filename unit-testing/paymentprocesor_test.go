package main

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestPaymentProcessor(t *testing.T) {
	tests := []struct {
		name  string
		total float64
	}{
		{"small order", 10.0},
		{"large order", 99.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProcessor := NewMockPaymentProcessor(ctrl)
			mockProcessor.EXPECT().Charge(tt.total).Return(nil).Times(1)
			order := &Order{Processor: mockProcessor, Total: tt.total}
			if err := order.Checkout(); err != nil {
				t.Errorf("Checkout() returned error: %v", err)
			}
		})
	}

}

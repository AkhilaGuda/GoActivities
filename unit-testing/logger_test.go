package main

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestLogger(t *testing.T) {
	testcases := []struct {
		name    string
		message string
	}{
		{"single work", "Work done"},
		{"Another work", "Work done"},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockLogger := NewMockLogger(ctrl)
			mockLogger.EXPECT().Info(tt.message).Times(1)
			service := &Service{Logger: mockLogger}
			service.DoWork()
		})
	}

}

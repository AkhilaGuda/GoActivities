package main

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNotify(t *testing.T) {
	// create gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// create mock repository
	mockRepo := NewMockNotifier(ctrl)
	// set up expectations
	mockRepo.EXPECT().Send("hello").Return(nil)
	service := UserService{Notifier: mockRepo}
	err := service.Notify("hello")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

}

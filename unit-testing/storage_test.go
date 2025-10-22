package main

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestStorage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	storageMock := NewMockStorage(ctrl)
	storageMock.EXPECT().Save("Data").Return(nil)
	storageMock.EXPECT().Save("").Return(errors.New("Cannot be empty"))
	if err := SaveData(storageMock, "Data"); err != nil {
		t.Errorf("Error :%v", err)
	}
	err := SaveData(storageMock, "")
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
	expect := "failed to save: Cannot be empty"
	if err.Error() != expect {
		t.Errorf("expected %q, got %q", expect, err.Error())
	}

}

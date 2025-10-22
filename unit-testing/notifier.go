package main

type Notifier interface {
	Send(msg string) error
}

type UserService struct {
	Notifier Notifier
}

func (u *UserService) Notify(msg string) error {
	return u.Notifier.Send(msg)
}

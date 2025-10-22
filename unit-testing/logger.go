package main

type Logger interface {
	Info(msg string)
}

type Service struct {
	Logger
}

func (s *Service) DoWork() {
	s.Info("Work done")
}

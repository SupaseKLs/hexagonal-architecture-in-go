package main

import "context"

type SayHelloPort interface {
	SayHelloPort(ctx context.Context, userID int) (string, error)
}

type GettingUserNameRepository interface {
	GetUserName(ctx context.Context, userID int) (string, error)
}

type HelloService struct {
	gettingUserNameRepo GettingUserNameRepository
}

func NewHelloService(gettingUserNameRepo GettingUserNameRepository) *HelloService {
	return &HelloService{
		gettingUserNameRepo: gettingUserNameRepo,
	}
}

func (s *HelloService) SayHello(ctx context.Context, userID int) (string, error) {
	userName, err := s.gettingUserNameRepo.GetUserName(ctx, userID)
	if err != nil {
		return "", err
	}
	return "Hello, " + userName + "!", nil
}

package main

import (
	"context"
	"log"
)

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

// Kog Adapter for driven port
type LogAdapterForGettingUserNameRepository struct {
	repo GettingUserNameRepository
}

func NewLogAdapterForGettingUserNameRepository(repo GettingUserNameRepository) *LogAdapterForGettingUserNameRepository {
	return &LogAdapterForGettingUserNameRepository{
		repo: repo,
	}
}

func (l *LogAdapterForGettingUserNameRepository) GetUserName(ctx context.Context, userID int) (string, error) {
	log.Println("Getting user name for userID:", userID)
	userName, err := l.repo.GetUserName(ctx, userID)
	if err != nil {
		log.Printf("Error getting user name for userID %d: %v", userID, err)
		return "", err
	}
	log.Printf("Successfully got user anme for userID %d: %s", userID, userName)
	return userName, nil
}

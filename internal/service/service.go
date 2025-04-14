package service

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const codeLength = 6

type Store interface {
	Save(short, original string) error
	Get(short string) (string, error)
}

type Service struct {
	store Store
	rand  *rand.Rand
}

func NewService(s Store) *Service {
	src := rand.NewSource(time.Now().UnixNano())
	return &Service{store: s, rand: rand.New(src)}
}

func (s *Service) generateRandomCode(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[s.rand.Intn(len(charset))]
	}
	return string(b)
}

func (s *Service) Shorten(url string) (string, error) {
	for {
		short := s.generateRandomCode(codeLength)
		_, err := s.store.Get(short)
		if err != nil {
			if saveErr := s.store.Save(short, url); saveErr != nil {
				return "", saveErr
			}
			return short, nil
		}
	}
}

func (s *Service) Resolve(short string) (string, error) {
	return s.store.Get(short)
}

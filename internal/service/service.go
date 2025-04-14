package service

import (
	"crypto/sha1"
	"encoding/hex"
)

type Store interface {
	Save(short, original string) error
	Get(short string) (string, error)
}

type Service struct {
	store Store
}

func NewService(s Store) *Service {
	return &Service{store: s}
}

func (s *Service) Shorten(url string) (string, error) {
	h := sha1.New()
	h.Write([]byte(url))
	short := hex.EncodeToString(h.Sum(nil))[:6]
	if err := s.store.Save(short, url); err != nil {
		return "", err
	}
	return short, nil
}

func (s *Service) Resolve(short string) (string, error) {
	return s.store.Get(short)
}

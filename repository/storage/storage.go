package storage

import (
	"errors"
	"sync"
)

type Storage struct {
	shortVsOriginalKeys map[string]string
	rwMutex             sync.RWMutex
}

func (s *Storage) Close() {}
func Init() *Storage {
	return &Storage{
		shortVsOriginalKeys: make(map[string]string),
		rwMutex:             sync.RWMutex{},
	}
}
func (s *Storage) PushOriginalAndShort(original, short string) error {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	_, has := s.shortVsOriginalKeys[short]

	if !has {
		s.shortVsOriginalKeys[short] = original
		return nil
	}
	return errors.New("link already exists")
}

func (s *Storage) GetByShortLink(short string) (string, error) {
	s.rwMutex.Lock()
	value, has := s.shortVsOriginalKeys[short]
	s.rwMutex.Unlock()
	if has {
		return value, nil
	}
	return "", errors.New("cant find original link")
}

package main

import (
	"errors"
)

type Storage map[string]string

func (s Storage) Get(key string) (string, error) {
	value, ok := s[key]
	if !ok {
		return "", errors.New("key is not found")
	}
	return value, nil
}

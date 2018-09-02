package main

import (
	"fmt"
	"sync"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

var store Store

//  temp value
var inc int

type Store interface {
	TodoGetter() ([]Todo, error)
	TodoCreater(*Todo) error
	TodoDeleter(id int) error
	TodoUpdater(int, *Todo) error
}

type dbStore struct {
	db   map[int]Todo
	lock sync.RWMutex
}

func InitStore(s Store) {
	store = s
}

type Todo struct {
	Completed   *bool   `json:"completed"`
	Description *string `json:"description"`
	CreatedDate string  `json:"created_date"`
	Id          int     `json:"id"`
	Name        *string `json:"name"`
}

func (s *dbStore) TodoGetter() ([]Todo, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	var tL []Todo
	for _, v := range s.db {
		tL = append(tL, v)
	}
	return tL, nil
}

func (s *dbStore) TodoCreater(t *Todo) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	inc++
	t.CreatedDate = time.Now().Format(TimeFormat)
	t.Id = inc
	s.db[t.Id] = *t
	return nil
}

func (s *dbStore) TodoDeleter(id int) error {
	if _, ok := s.db[id]; ok {
		s.lock.Lock()
		defer s.lock.Unlock()
		delete(s.db, id)
		return nil
	}
	err := fmt.Errorf("id: %v not found", id)
	return err
}

func (s *dbStore) TodoUpdater(id int, t *Todo) error {
	if v, ok := s.db[id]; ok {
		if t.Completed != nil {
			v.Completed = t.Completed
		}
		if t.Description != nil {
			v.Description = t.Description
		}
		if t.Name != nil {
			v.Name = t.Name
		}
		s.lock.Lock()
		defer s.lock.Unlock()
		s.db[id] = v
		return nil
	}
	err := fmt.Errorf("id: %v not found", id)
	return err
}

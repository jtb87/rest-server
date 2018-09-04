package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

// const values
const TimeFormat = "2006-01-02 15:04:05"

var store Store

//  temp value
var inc int

type Store interface {
	TodoGetter() ([]Todo, error)
	TodoCreater(Todo) error
	TodoDeleter(id int) error
	TodoUpdater(int, *Todo) error
	loadFromJSON() error
	storeAsJSON() error
}

type dbStore struct {
	db   map[int]Todo
	lock sync.RWMutex
}

func InitStore(s Store) {
	store = s
}

type Todo struct {
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
	CreatedDate string `json:"created_date"`
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Tag         string `json:"complete_before"`
	Hide        bool   `json:"hide"`
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

func (s *dbStore) TodoCreater(t Todo) error {
	s.lock.Lock()
	s.db[t.Id] = t
	s.lock.Unlock()
	err := s.storeAsJSON()
	if err != nil {
		return err
	}
	return nil
}

func (s *dbStore) TodoDeleter(id int) error {
	if _, ok := s.db[id]; ok {
		s.lock.Lock()
		delete(s.db, id)
		s.lock.Unlock()
		err := s.storeAsJSON()
		if err != nil {
			return err
		}
		return nil
	}
	err := fmt.Errorf("id: %v not found", id)
	return err
}

func (s *dbStore) TodoUpdater(id int, t *Todo) error {
	if v, ok := s.db[id]; ok {
		v.Completed = t.Completed
		v.Description = t.Description
		v.Name = t.Name
		s.lock.Lock()
		s.db[id] = v
		s.lock.Unlock()
		err := s.storeAsJSON()
		if err != nil {
			return err
		}
		return nil
	}
	err := fmt.Errorf("id: %v not found", id)
	return err
}

func (s *dbStore) storeAsJSON() error {
	d, err := s.TodoGetter()
	if err != nil {
		return err
	}
	var b []byte
	b, err = json.Marshal(d)
	if err != nil {
		return err
	}
	writeJsonFile(b)
	return nil
}

func (s *dbStore) loadFromJSON() error {
	f, err := ioutil.ReadFile("database.json")
	if err != nil {
		b := []byte("[]")
		writeJsonFile(b)
		return fmt.Errorf("%s", "database.json did not exist created now")
	}
	var todos []Todo
	err = json.Unmarshal(f, &todos)
	if err != nil {
		return err
	}
	for _, v := range todos {
		err := s.TodoCreater(v)
		if err != nil {
			return err
		}
		setBiggestAsInc(v.Id)
	}
	return nil
}

func setBiggestAsInc(id int) {
	if id > inc {
		inc = id
	}
}

func setTimeNow() string {
	t := time.Now()
	return t.Format(TimeFormat)
}

func writeJsonFile(b []byte) {
	ioutil.WriteFile("database.json", b, 0666)
}

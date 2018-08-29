package main

import ()

var store Store

//  temp value
var inc int

type Store interface {
	TodoGetter() []Todo
	TodoCreater(*Todo)
	// TodoDeleter()
	// TodoPutter()
}

type dbStore struct {
	db map[int]Todo
}

type Todo struct {
	Status      bool   `json:"status"`
	Description string `json:"description"`
	CreatedDate string `json:"created_date"`
	Id          int    `json:"id"`
}

func (s *dbStore) TodoGetter() []Todo {
	var tL []Todo
	for _, v := range s.db {
		tL = append(tL, v)
	}
	return tL
}

func (s *dbStore) TodoCreater(t *Todo) {
	s.db[t.Id] = *t
}

func (s *dbStore) TodoDeleter() {
}

func (s *dbStore) TodoUpdater() {
}

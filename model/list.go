package model

import (
	"sync"
)

type List struct {
	mutex sync.Mutex
	List  []string
}

type CommandsList struct {
	mutex sync.RWMutex
	Store map[string]*List
}

func CreateNewCommandsList() *CommandsList {
	return &CommandsList{
		Store: make(map[string]*List),
	}
}

func (ls *CommandsList) ListExists(name string) bool {
	ls.mutex.Lock()
	defer ls.mutex.Unlock()

	_, isPresent := ls.Store[name]
	return isPresent
}

func (ls *CommandsList) CreateList(name string) {
	ls.mutex.Lock()
	defer ls.mutex.Unlock()

	ls.Store[name] = &List{
		List: make([]string, 0),
	}
}

func (ls *CommandsList) LeftPushInList(name string, vals []string) {
	ls.mutex.Lock()
	defer ls.mutex.Unlock()

	addToList(ls.Store[name], vals, "Left")
}

func (ls *CommandsList) RightPushInList(name string, vals []string) {
	ls.mutex.Lock()
	defer ls.mutex.Unlock()

	addToList(ls.Store[name], vals, "Right")
}

func addToList(list *List, vals []string, lor string) {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	if lor == "Left" {
		list.List = append(vals, list.List...)
	} else if lor == "Right" {
		list.List = append(list.List, vals...)
	}
}

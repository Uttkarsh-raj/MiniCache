package database

import (
	"fmt"
	"sync"
)

type Database struct {
	mutex sync.RWMutex
	Db    map[string]string
}

func CreateDatabase() *Database {
	return &Database{
		Db: make(map[string]string),
	}
}

func (db *Database) SetValue(key, value string) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Db[key] = value
}

func (db *Database) GetValue(key string) (string, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	value, ok := db.Db[key]
	if !ok {
		return "", fmt.Errorf("error: the value of the key %s is not already present. Please set the value before fetching", key)
	}
	return value, nil
}

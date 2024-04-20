package diskmap

import (
	"os"
	"path/filepath"
	"sync"
)

type DiskMap struct {
	path         string
	processedNow map[string]*sync.Mutex
}

func NewDiskMap(path string) *DiskMap {
	if path == "" {
		panic("please set path for DiskMap")
	}

	err := os.Mkdir(path, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		panic("cannot create db directory at " + path + ": " + err.Error())
	}

	return &DiskMap{
		path:         path,
		processedNow: make(map[string]*sync.Mutex),
	}
}

func (db *DiskMap) Set(key string, value []byte) error {

	if _, ok := db.processedNow[key]; !ok {
		db.processedNow[key] = &sync.Mutex{}
	}

	mu := db.processedNow[key]
	mu.Lock()
	defer mu.Unlock()

	filePath := filepath.Join(db.path, key)
	err := os.WriteFile(filePath, value, 0644)

	if err != nil {
		return err
	}
	return nil
}

func (db *DiskMap) Get(key string) ([]byte, error) {
	filePath := filepath.Join(db.path, key)
	return os.ReadFile(filePath)
}

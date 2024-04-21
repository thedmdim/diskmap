package diskmap

import (
	"os"
	"path/filepath"
	"sync"
)

// DiskMap represents a simple key-value storage system that persists data on disk.
type DiskMap struct {
	path         string                 // The directory path where data files are stored
	processedNow map[string]*sync.Mutex // Map of file locks for concurrent access control
}

// NewDiskMap creates a new instance of DiskMap with the specified storage path.
// It creates the directory if it doesn't exist.
func NewDiskMap(path string) *DiskMap {
	if path == "" {
		panic("please set path for DiskMap")
	}

	// Ensure the directory exists or create it
	err := os.Mkdir(path, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		panic("cannot create db directory at " + path + ": " + err.Error())
	}

	return &DiskMap{
		path:         path,
		processedNow: make(map[string]*sync.Mutex),
	}
}

// Set stores the given key-value pair in the DiskMap
// It ensures concurrent safety by locking access to the file associated with the key
func (db *DiskMap) Set(key string, value []byte) error {

	// Ensure only one goroutine can access the file for this key at a time
	if _, ok := db.processedNow[key]; !ok {
		db.processedNow[key] = &sync.Mutex{}
	}

	mu := db.processedNow[key]
	mu.Lock()
	defer mu.Unlock()

	// Write the value to a file
	filePath := filepath.Join(db.path, key)
	err := os.WriteFile(filePath, value, 0644)

	if err != nil {
		return err
	}
	return nil
}

// Get retrieves the value associated with the given key from the DiskMap
func (db *DiskMap) Get(key string) ([]byte, error) {
	filePath := filepath.Join(db.path, key)
	return os.ReadFile(filePath)
}

// Del deletes a file assitiated with a key
func (db *DiskMap) Del(key string) error {
    filePath := filepath.Join(db.path, key)
    err := os.Remove(filePath)
    if err != nil {
        return err
    }
    return nil
}
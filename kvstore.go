package kvstore

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

type KVStore struct {
	data     sync.Map
	filePath string
	done     chan struct{}
}

func NewKVStore(filePath string) (*KVStore, error) {
	store := &KVStore{
		filePath: filePath,
		done:     make(chan struct{}),
	}

	// Wczytaj dane z pliku jeśli istnieje
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		if err := store.load(); err != nil {
			return nil, err
		}
	}

	// Uruchom okresowe zapisywanie w tle
	go store.periodicSave()

	return store, nil
}

func (kv *KVStore) Set(key string, value interface{}) {
	kv.data.Store(key, value)
}

func (kv *KVStore) Get(key string) (interface{}, bool) {
	return kv.data.Load(key)
}

func (kv *KVStore) Delete(key string) {
	kv.data.Delete(key)
}

func (kv *KVStore) periodicSave() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			kv.save()
		case <-kv.done:
			return
		}
	}
}

func (kv *KVStore) load() error {
	file, err := os.ReadFile(kv.filePath)
	if err != nil {
		return err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(file, &data); err != nil {
		return err
	}

	for k, v := range data {
		kv.data.Store(k, v)
	}

	return nil
}

func (kv *KVStore) save() error {
	data := make(map[string]interface{})
	kv.data.Range(func(key, value interface{}) bool {
		if k, ok := key.(string); ok {
			data[k] = value
		}
		return true
	})

	file, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(kv.filePath, file, 0644)
}

func (kv *KVStore) Close() error {
	close(kv.done)
	return kv.save()
}

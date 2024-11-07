package kvstore

import (
	"strconv"
	"testing"
)

func BenchmarkSet(b *testing.B) {
	store, _ := NewKVStore("test.json", false)
	defer store.Close()

	for i := 0; i < b.N; i++ {
		store.Set("key", map[string]interface{}{"name": "John", "age": 30})
	}
}

func BenchmarkGet(b *testing.B) {
	store, _ := NewKVStore("test.json", false)
	defer store.Close()

	store.Set("key", map[string]interface{}{"name": "John", "age": 30})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = store.Get("key")
	}
}

func BenchmarkDelete(b *testing.B) {
	store, _ := NewKVStore("test.json", false)
	defer store.Close()

	store.Set("key", map[string]interface{}{"name": "John", "age": 30})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Delete("key")
	}
}

func BenchmarkKeys(b *testing.B) {
	store, _ := NewKVStore("test.json", false)
	defer store.Close()

	for i := 0; i < 1000; i++ {
		store.Set("key"+strconv.Itoa(i), map[string]interface{}{"name": "John", "age": 30})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		store.Keys()
	}
}

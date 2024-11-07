package kvstore

import (
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

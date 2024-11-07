# KVStore

KVStore is a lightweight, thread-safe key-value store implementation in Go with persistent storage capabilities. It provides a simple interface for storing and retrieving data while offering automatic persistence to disk.

## Features

- Thread-safe operations using `sync.Map`
- Persistent storage to JSON file
- Configurable save strategies:
  - Always save on write operations
  - Periodic automatic saving (default: every 1 minute)
- Automatic data recovery from file on initialization
- Clean shutdown handling

## Installation

```go
import "github.com/Mateusz779/go_kvstorage"
```

## Usage

### Creating a New Store

```go
// Create a new store with always-save enabled
store, err := kvstore.NewKVStore("data.json", true)
if err != nil {
    log.Fatal(err)
}
defer store.Close()
```

### Basic Operations

```go
// Set a value
store.Set("user1", map[string]interface{}{
    "name": "John Doe",
    "age":  30,
})

// Get a value
value, exists := store.Get("user1")
if exists {
    // Use the value
}

// Delete a value
store.Delete("user1")
```

### Configuration Options

The KVStore can be configured with different saving strategies:

- `alwaysSave`: When true, every write operation (Set/Delete) triggers an immediate save to disk
- `periodicSave`: When true (default), the store automatically saves to disk every minute

## Technical Details

### Data Persistence

- Data is stored in JSON format
- File permissions are set to 0644
- The store automatically loads existing data on initialization
- Periodic saves happen in a background goroutine

### Thread Safety

All operations are thread-safe thanks to the use of `sync.Map`, making KVStore suitable for concurrent access in multi-goroutine environments.

### Graceful Shutdown

Always call `Close()` when you're done with the store to ensure:
- All pending saves are completed
- Background goroutines are properly terminated
- Final state is persisted to disk

## Error Handling

The store handles several types of errors:
- File read/write errors
- JSON marshaling/unmarshaling errors
- File system permission errors

## Limitations

- Keys must be strings
- Values must be JSON-serializable
- No built-in data expiration
- No compression of stored data

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
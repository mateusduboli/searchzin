package kvstorage

// KVStorage an interface to define a key value storage
type KVStorage interface {
	// Set the value of a given key
	Set(key string, value interface{}) error
	// Get the value for a given key
	Get(key string) (interface{}, error)
	// Deletes the value of a given key
	Del(key string) error
	// List all values inside the storage
	List() ([]interface{}, error)
}

// MemoryKVStorage a in memory representation of a key value storage
type MemoryKVStorage struct {
	data map[string]interface{}
	KVStorage
}

// NewMemoryKVStorage Creates a MemoryKVStorage instance
func NewMemoryKVStorage() (*MemoryKVStorage, error) {
	return &MemoryKVStorage{
		data: make(map[string]interface{}),
	}, nil
}

// Set the value of a given key in the key value store
func (m *MemoryKVStorage) Set(key string, value interface{}) error {
	m.data[key] = value
	return nil
}

// Get the value of a given key in the key value store
func (m *MemoryKVStorage) Get(key string) (interface{}, error) {
	return m.data[key], nil
}

// Del deletes a key in the key value store
func (m *MemoryKVStorage) Del(key string) error {
	delete(m.data, key)
	return nil
}

// List values inside the key value store
func (m *MemoryKVStorage) List() ([]interface{}, error) {
	values := make([]interface{}, len(m.data))

	for _, v := range m.data {
		values = append(values, v)
	}

	return values, nil
}

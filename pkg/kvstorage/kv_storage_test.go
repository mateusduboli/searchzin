package kvstorage

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

/*
	Tests here should follow the naming convention
	Test_$(Action)_$(Scenario)_Should_$(Expectation)
*/

func Test_New_NoArguments_Should_CreateAKVStorage(t *testing.T) {
	m, err := NewMemoryKVStorage()
	if err != nil {
		t.Fatal(err)
	}
	if m == nil {
		t.Fatalf("Failed to create in-memory kv storage")
	}
}

func Test_GetAnyKey_OnEmpty_Should_ReturnNil(t *testing.T) {
	m, err := NewMemoryKVStorage()
	if err != nil {
		t.Fatal(err)
	}

	key := "ANY_KEY"
	actual, err := m.Get(key)

	if err != nil {
		t.Fatal(err)
	}

	if actual != nil {
		t.Errorf("Empty kv should return nil if empty, actual: %s", actual)
	}
}

func Test_List_OnEmpty_Should_RetrieveEmptyList(t *testing.T) {
	m, err := NewMemoryKVStorage()
	if err != nil {
		t.Fatal(err)
	}

	actual, err := m.List()

	if err != nil {
		t.Fatal(err)
	}

	if len(actual) != 0 {
		t.Errorf("Empty kv should return empty list if empty, actual: %s", actual)
	}
}

func Test_GetValidKey_OnPopulated_Should_ReturnValue(t *testing.T) {
	m, err := NewMemoryKVStorage()
	if err != nil {
		t.Fatal(err)
	}

	key := "1"
	expected := "T3S7D4TA"

	err = m.Set(key, expected)

	if err != nil {
		t.Fatal(err)
	}

	actual, err := m.Get(key)

	if err != nil {
		t.Fatal(err)
	}

	if expected != actual {
		t.Errorf("Failed to store and retrieve a document exp: %s, act: %s", expected, actual)
	}
}

func Test_List_OnSingleElementStorage_Should_SingleElementList(t *testing.T) {
	m, err := NewMemoryKVStorage()
	if err != nil {
		t.Error(err)
	}

	key := "1"
	value := "T3S7D4TA"

	expected := []interface{}{value}

	err = m.Set(key, value)

	if err != nil {
		t.Fatal(err)
	}

	actual, err := m.List()

	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(expected, actual) {
		t.Errorf("Failed to store and retrieve a document exp: %s, act: %s", expected, actual)
	}
}

func Test_List_OnPopulated_Should_ReturnAllValues(t *testing.T) {
	m, err := NewMemoryKVStorage()
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < 100; i++ {
		key := strconv.Itoa(i)
		value := fmt.Sprintf("T3S7D4TA_%d", i)

		err = m.Set(key, value)

		if err != nil {
			t.Errorf("Failed to set %s:%s, err: %s", key, value, err)
		}
	}

	expected := make([]interface{}, 100)

	for i := 0; i < 100; i++ {
		expected = append(expected, fmt.Sprintf("T3S7D4TA_%d", i))
	}

	actual, err := m.List()

	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(expected, actual) {
		t.Errorf("Value mismatch in storage, exp: %s, act: %s", expected, actual)
	}
}

func Test_DeleteAnyKey_OnEmpty_Should_NotFail(t *testing.T) {
	m, err := NewMemoryKVStorage()
	if err != nil {
		t.Fatal(err)
	}

	key := "ANY_KEY"
	err = m.Del(key)

	if err != nil {
		t.Errorf("Deleting on an empty storage should not fail, err: %s", err)
	}
}

func Test_GetAfterDel_OnPopulated_Should_ReturnNil(t *testing.T) {
	m, err := NewMemoryKVStorage()
	if err != nil {
		t.Fatal(err)
	}

	key := "ANY_KEY"
	value := "ANY_VALUE"

	err = m.Set(key, value)
	if err != nil {
		t.Fatal(err)
	}

	err = m.Del(key)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := m.Get(key)
	if err != nil {
		t.Fatal(err)
	}

	if actual != nil {
		t.Errorf("Get after a deletion on the same key should make it nil, actual: %s", actual)
	}
}

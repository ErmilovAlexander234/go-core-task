package main

import "testing"

func TestAddAndGet(t *testing.T) {

	m := NewStringIntMap()

	m.Add("a", 10)

	val, ok := m.Get("a")

	if !ok || val != 10 {
		t.Error("Add or Get failed")
	}
}

func TestRemove(t *testing.T) {

	m := NewStringIntMap()

	m.Add("a", 5)
	m.Remove("a")

	if m.Exists("a") {
		t.Error("Remove failed")
	}
}

func TestExists(t *testing.T) {

	m := NewStringIntMap()

	m.Add("key", 1)

	if !m.Exists("key") {
		t.Error("Exists should return true")
	}
}

func TestCopy(t *testing.T) {

	m := NewStringIntMap()

	m.Add("a", 1)

	copyMap := m.Copy()

	m.Add("b", 2)

	if _, ok := copyMap["b"]; ok {
		t.Error("Copy should not change after original map changes")
	}
}

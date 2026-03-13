package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int)

	for k, v := range m.data {
		copyMap[k] = v
	}

	return copyMap
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	val, ok := m.data[key]
	return val, ok
}

func main() {

	m := NewStringIntMap()

	m.Add("apple", 5)
	m.Add("banana", 10)

	fmt.Println("Map:", m.data)

	val, ok := m.Get("apple")
	fmt.Println("Get apple:", val, ok)

	fmt.Println("Exists banana:", m.Exists("banana"))

	m.Remove("banana")
	fmt.Println("After remove:", m.data)

	copyMap := m.Copy()
	fmt.Println("Copy:", copyMap)
}

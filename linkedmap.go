package linkedmap

import (
	"container/list"
	"fmt"
	"strings"
)

type linkedMapElement struct {
	keyElement *list.Element
	value      interface{}
}

// LinkedMap is a linked hash map.
type LinkedMap struct {
	hashMap map[interface{}]*linkedMapElement
	keys    *list.List
}

// New returns an initialized LinkedMap.
func New() *LinkedMap {
	return new(LinkedMap).Init()
}

func (m *LinkedMap) inited() bool {
	return m.hashMap != nil && m.keys != nil
}

func (m *LinkedMap) lazyInit() {
	if !m.inited() {
		m.Init()
	}
}

// Init initializes or clears LinkedMap m.
func (m *LinkedMap) Init() *LinkedMap {
	m.hashMap = make(map[interface{}]*linkedMapElement)
	m.keys = list.New()
	return m
}

// Set sets value for key to end of the map,
// updates existing keys.
func (m *LinkedMap) Set(key interface{}, val interface{}) (isNewKey bool) {
	if mapVal, exists := m.hashMap[key]; exists {
		mapVal.value = val
		return false
	}
	m.lazyInit()
	mapVal := new(linkedMapElement)
	mapVal.keyElement = m.keys.PushBack(key)
	mapVal.value = val
	m.hashMap[key] = mapVal
	return true
}

// Add adds value for key to end of the map,
// ignores existing keys.
func (m *LinkedMap) Add(key interface{}, val interface{}) (isNewKey bool) {
	if _, exists := m.hashMap[key]; exists {
		return false
	}
	m.lazyInit()
	mapVal := new(linkedMapElement)
	mapVal.keyElement = m.keys.PushBack(key)
	mapVal.value = val
	m.hashMap[key] = mapVal
	return true
}

// MoveToBack moves key to the back of the map.
func (m *LinkedMap) MoveToBack(key interface{}) (ok bool) {
	mapValue, exists := m.hashMap[key]
	if !exists {
		return false
	}
	m.keys.MoveToBack(mapValue.keyElement)
	return true
}

// MoveToBack moves key to the front of the map.
func (m *LinkedMap) MoveToFront(key interface{}) (ok bool) {
	mapValue, exists := m.hashMap[key]
	if !exists {
		return false
	}
	m.keys.MoveToFront(mapValue.keyElement)
	return true
}

// MoveBefore moves key to its new position before mark.
func (m *LinkedMap) MoveBefore(key, mark interface{}) (ok bool) {
	mapValue, exists := m.hashMap[key]
	if !exists {
		return false
	}

	markValue, exists := m.hashMap[mark]
	if !exists {
		return false
	}
	m.keys.MoveBefore(mapValue.keyElement, markValue.keyElement)
	return true
}

// MoveBefore moves key to its new position after mark.
func (m *LinkedMap) MoveAfter(key, mark interface{}) (ok bool) {
	mapValue, exists := m.hashMap[key]
	if !exists {
		return false
	}

	markValue, exists := m.hashMap[mark]
	if !exists {
		return false
	}
	m.keys.MoveAfter(mapValue.keyElement, markValue.keyElement)
	return true
}

// Front returns the first key and value of the map.
func (m *LinkedMap) Front() (key, value interface{}, ok bool) {
	if !m.inited() {
		return nil, nil, false
	}
	frontKey := m.keys.Front()
	if frontKey == nil {
		return nil, nil, false
	}
	key = frontKey.Value
	return key, m.hashMap[key].value, true
}

// Back returns the last key and value of the map.
func (m *LinkedMap) Back() (key, value interface{}, ok bool) {
	if !m.inited() {
		return nil, nil, false
	}
	backKey := m.keys.Back()
	if backKey == nil {
		return nil, nil, false
	}
	key = backKey.Value
	return key, m.hashMap[key].value, true
}

// Has returns whether has the key
func (m *LinkedMap) Has(key interface{}) (has bool) {
	_, has = m.hashMap[key]
	return has
}

// Load returns value in the map for the key
func (m *LinkedMap) Load(key interface{}) (val interface{}, ok bool) {
	mapVal, ok := m.hashMap[key]
	if !ok {
		return nil, false
	}
	return mapVal.value, true
}

// Remove removes key in the map, and returns that value
func (m *LinkedMap) Remove(key interface{}) (val interface{}, ok bool) {
	mapVal, exists := m.hashMap[key]
	if !exists {
		return nil, false
	}
	delete(m.hashMap, key)
	m.keys.Remove(mapVal.keyElement)
	return mapVal.value, true
}

// Range calls f for each key value in the map.
// If f returns false, break the iteration.
func (m *LinkedMap) Range(f func(key, value interface{}) bool) {
	if !m.inited() {
		return
	}
	for elem := m.keys.Front(); elem != nil; elem = elem.Next() {
		if !f(elem.Value, m.hashMap[elem.Value].value) {
			break
		}
	}
}

// String returns string representing the map in the form "{key1: value1, key2: value2}".
func (m *LinkedMap) String() string {
	var builder strings.Builder
	first := true
	builder.WriteString("{")
	m.Range(func(key, value interface{}) bool {
		if first {
			first = false
		} else {
			builder.WriteString(", ")
		}

		builder.WriteString(fmt.Sprint(key))
		builder.WriteString(": ")
		builder.WriteString(fmt.Sprint(value))
		return true
	})

	builder.WriteString("}")
	return builder.String()
}

// Len returns the map current length
func (m *LinkedMap) Len() int {
	return len(m.hashMap)
}

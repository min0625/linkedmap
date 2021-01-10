package linkedmap

import (
	"reflect"
	"testing"
)

func TestLinkedMap(t *testing.T) {
	var (
		key1 = "key1"
		key2 = "key2"
		key3 = "key3"
	)

	var (
		val1 = "val1"
		val2 = "val2"
		val3 = "val3"
	)

	m := New()
	if !checkMap(t, m, []*element{}) {
		return
	}

	m.Add(key1, val1)
	if !checkMap(t, m, []*element{
		{
			key:   key1,
			value: val1,
		},
	}) {
		return
	}

	m.Add(key2, val2)
	if !checkMap(t, m, []*element{
		{
			key:   key1,
			value: val1,
		},
		{
			key:   key2,
			value: val2,
		},
	}) {
		return
	}

	m.Add(key3, val3)
	if !checkMap(t, m, []*element{
		{
			key:   key1,
			value: val1,
		},
		{
			key:   key2,
			value: val2,
		},
		{
			key:   key3,
			value: val3,
		},
	}) {
		return
	}

	m.Remove(key2)
	if !checkMap(t, m, []*element{
		{
			key:   key1,
			value: val1,
		},
		{
			key:   key3,
			value: val3,
		},
	}) {
		return
	}

	m.Set(key1, val2)
	if !checkMap(t, m, []*element{
		{
			key:   key1,
			value: val2,
		},
		{
			key:   key3,
			value: val3,
		},
	}) {
		return
	}

	m.Set(key2, val1)
	if !checkMap(t, m, []*element{
		{
			key:   key1,
			value: val2,
		},
		{
			key:   key3,
			value: val3,
		},
		{
			key:   key2,
			value: val1,
		},
	}) {
		return
	}
}

type element struct {
	key   interface{}
	value interface{}
}

func checkLen(t *testing.T, m *LinkedMap, len int) bool {
	if n := m.Len(); n != len {
		t.Errorf("Len() is %d, want %d", n, len)
		return false
	}
	return true
}

func checkRange(t *testing.T, m *LinkedMap, es []*element) bool {
	for breakIndex := 0; breakIndex < len(es); breakIndex++ {
		index := 0
		esFromMap := make([]*element, 0, len(es))
		m.Range(func(key, value interface{}) bool {
			if index == breakIndex {
				// break iteration
				return false
			}
			if index > breakIndex {
				t.Errorf("Range() should be not continue iteration")
				// break iteration
				return false
			}
			esFromMap = append(esFromMap, &element{
				key:   key,
				value: value,
			})
			index++
			// continue iteration
			return true
		})
		if t.Failed() {
			return false
		}
		if !reflect.DeepEqual(esFromMap, es[:index]) {
			t.Errorf("Range() to elements is unexpected")
			return false
		}
	}
	return true
}

func checkLoad(t *testing.T, m *LinkedMap, es []*element) bool {
	if got, exists := m.Load(new(int)); exists || got != nil {
		t.Errorf("Load() got not exists key in the map")
		return false
	}

	for _, e := range es {
		valFromMap, exists := m.Load(e.key)
		if !exists {
			t.Errorf("Load() not found exists key")
			return false
		}
		if !reflect.DeepEqual(valFromMap, e.value) {
			t.Errorf("Load() value is unexpected")
			return false
		}
	}
	return true
}

func checkFront(t *testing.T, m *LinkedMap, es []*element) bool {
	key, val, ok := m.Front()
	if len(es) == 0 {
		if key == nil && val == nil && !ok {
			return true
		}
		t.Errorf("Front() got not exists key, val, ok")
		return false
	}
	e := es[0]
	if !ok {
		t.Errorf("Front() not found first key, value")
		return false
	}

	if !reflect.DeepEqual(key, e.key) || !reflect.DeepEqual(val, e.value) {
		t.Errorf("Front() got key or value is unexpected")
		return false
	}
	return true
}

func checkBack(t *testing.T, m *LinkedMap, es []*element) bool {
	key, val, ok := m.Back()
	if len(es) == 0 {
		if key == nil && val == nil && !ok {
			return true
		}
		t.Errorf("Back() got not exists key, val, ok")
		return false
	}
	e := es[len(es)-1]
	if !ok {
		t.Errorf("Back() not found last key, value")
		return false
	}

	if !reflect.DeepEqual(key, e.key) || !reflect.DeepEqual(val, e.value) {
		t.Errorf("Back() got key or value is unexpected")
		return false
	}
	return true
}

func checkHas(t *testing.T, m *LinkedMap, es []*element) bool {
	if has := m.Has(new(int)); has {
		t.Errorf("Has() has not exists key in the map")
		return false
	}

	for _, e := range es {
		if has := m.Has(e.key); !has {
			t.Errorf("Has() not has exists key")
			return false
		}
	}
	return true
}

func checkMap(t *testing.T, m *LinkedMap, es []*element) bool {
	if !checkLen(t, m, len(es)) {
		return false
	}

	if !checkRange(t, m, es) {
		return false
	}

	if !checkLoad(t, m, es) {
		return false
	}

	if !checkHas(t, m, es) {
		return false
	}

	if !checkFront(t, m, es) {
		return false
	}

	if !checkBack(t, m, es) {
		return false
	}

	return true
}

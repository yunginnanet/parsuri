package buffer

import "testing"

func queueWithItems[T any](t *testing.T, items ...T) *Queue[T] {
	t.Helper()
	q := NewQueue[T](0)
	for _, item := range items {
		q.Push(item)
	}
	return q
}

func TestPushAddsItemToQueue(t *testing.T) {
	q := NewQueue[int](0)
	q.Push(42)
	if q.Len() != 1 {
		t.Errorf("expected queue length to be 1, got %d", q.Len())
	}
}

func TestNextReturnsLastItemAndRemovesIt(t *testing.T) {
	q := queueWithItems(t, 1, 2, 3)
	item, ok := q.Pop()
	if !ok {
		t.Errorf("expected ok to be true, got false")
	}
	if item != 1 {
		t.Errorf("expected item to be 1, got %d", item)
	}
	if q.Len() != 2 {
		t.Errorf("expected queue length to be 2, got %d", q.Len())
	}
}

func TestNextReturnsFalseWhenQueueIsEmpty(t *testing.T) {
	q := NewQueue[int](0)
	_, ok := q.Pop()
	if ok {
		t.Errorf("expected ok to be false, got true")
	}
}

func TestPopFrontReturnsFirstItemAndRemovesIt(t *testing.T) {
	q := queueWithItems(t, 1, 2, 3)
	item, ok := q.PopFront()
	if !ok {
		t.Errorf("expected ok to be true, got false")
	}
	if item != 3 {
		t.Errorf("expected item to be 3, got %d", item)
	}
	if q.Len() != 2 {
		t.Errorf("expected queue length to be 2, got %d", q.Len())
	}
}

func TestPopFrontReturnsFalseWhenQueueIsEmpty(t *testing.T) {
	q := NewQueue[int](0)
	_, ok := q.PopFront()
	if ok {
		t.Errorf("expected ok to be false, got true")
	}
}

func TestClearEmptiesTheQueue(t *testing.T) {
	q := queueWithItems(t, 1, 2, 3)
	q.Clear()
	if q.Len() != 0 {
		t.Errorf("expected queue length to be 0, got %d", q.Len())
	}
}

func TestLenReturnsCorrectQueueLength(t *testing.T) {
	q := queueWithItems(t, 1, 2, 3)
	if q.Len() != 3 {
		t.Errorf("expected queue length to be 3, got %d", q.Len())
	}
}

func TestStructs(t *testing.T) {
	type TestStruct struct {
		Name string
		Age  int
	}

	q := NewQueue[TestStruct](0)
	q.Push(TestStruct{Name: "Alice", Age: 30})
	q.Push(TestStruct{Name: "Bob", Age: 25})

	item, ok := q.Pop()
	if !ok {
		t.Errorf("expected ok to be true, got false")
	}

	if item.Name != "Alice" || item.Age != 30 {
		t.Errorf("expected item to be {Name: Alice, Age: 30}, got {%s, %d}", item.Name, item.Age)
	}

	item, ok = q.Pop()
	if !ok {
		t.Errorf("expected ok to be true, got false")
	}

	if item.Name != "Bob" || item.Age != 25 {
		t.Errorf("expected item to be {Name: Bob, Age: 25}, got {%s, %d}", item.Name, item.Age)
	}

	if q.Len() != 0 {
		t.Errorf("expected queue length to be 0, got %d", q.Len())
	}

	t.Run("PopFrontAndSkipEmpty", func(t *testing.T) {
		q.Push(TestStruct{})
		q.Push(TestStruct{Name: "Charlie", Age: 35})
		q.Push(TestStruct{Name: "Yeet", Age: 0})
		q.Push(TestStruct{})

		if item, ok = q.PopFront(); !ok {
			t.Errorf("expected ok to be true, got false")
		}
		if item.Name != "Yeet" || item.Age != 0 {
			t.Errorf("expected item to be {Name: Yeet, Age: 0}, got {%s, %d}", item.Name, item.Age)
		}

		if item, ok = q.Pop(); !ok {
			t.Errorf("expected ok to be true, got false")
		}

		if item.Name != "Charlie" || item.Age != 35 {
			t.Errorf("expected item to be {Name: Charlie, Age: 35}, got {%s, %d}", item.Name, item.Age)
		}

		if item, ok = q.Pop(); ok {
			t.Errorf("expected ok to be false, got true")
		}
	})
}

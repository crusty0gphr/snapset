package snapset_test

import (
	"testing"

	"github.com/snapset"
)

// TestInsert checks the Insert method.
func TestInsert(t *testing.T) {
	s := snapset.New[int](snapset.DefaultBucketSize)

	// Insert a new element
	idx := s.Insert(10)
	if idx != 0 {
		t.Errorf("Expected index 0, got %d", idx)
	}

	// Verify that the element exists
	if !s.Exists(10) {
		t.Errorf("Element 10 should exist after insertion")
	}

	// Insert another element
	idx = s.Insert(20)
	if idx != 1 {
		t.Errorf("Expected index 1, got %d", idx)
	}

	// Verify that both elements exist
	if !s.Exists(20) || !s.Exists(10) {
		t.Errorf("Elements 10 and 20 should exist after insertion")
	}
}

// TestDelete checks the Delete method.
func TestDelete(t *testing.T) {
	s := snapset.New[int](snapset.DefaultBucketSize)

	// Attempt to delete from an empty set
	_, ok := s.Delete(10)
	if ok {
		t.Errorf("Should not be able to delete from an empty set")
	}

	// Insert elements
	s.Insert(10)
	s.Insert(20)
	s.Insert(30)

	// Delete an existing element
	idx, ok := s.Delete(20)
	if !ok {
		t.Errorf("Failed to delete existing element 20")
	}

	// Verify index
	if idx != 1 {
		t.Errorf("Expected index 1 for deleted element, got %d", idx)
	}

	// Verify that the element no longer exists
	if s.Exists(20) {
		t.Errorf("Element 20 should not exist after deletion")
	}

	// Verify that other elements still exist
	if !s.Exists(10) || !s.Exists(30) {
		t.Errorf("Elements 10 and 30 should still exist after deletion")
	}

	// Delete a non-existing element
	_, ok = s.Delete(40)
	if ok {
		t.Errorf("Should not be able to delete non-existing element 40")
	}
}

// TestExists checks the Exists method.
func TestExists(t *testing.T) {
	s := snapset.New[string](snapset.DefaultBucketSize)

	// Insert elements
	s.Insert("apple")
	s.Insert("banana")

	// Check existence
	if !s.Exists("apple") {
		t.Errorf("Element 'apple' should exist")
	}

	if s.Exists("cherry") {
		t.Errorf("Element 'cherry' should not exist")
	}
}

// TestGetRandom checks the GetRandom method.
func TestGetRandom(t *testing.T) {
	s := snapset.New[int](snapset.DefaultBucketSize)

	// Insert elements
	s.Insert(1)
	s.Insert(2)
	s.Insert(3)

	// Collect results
	results := make(map[int]bool)
	for i := 0; i < 100; i++ {
		val := s.GetRandom()
		results[val] = true
	}

	// Verify that all inserted elements are returned at least once
	for _, val := range []int{1, 2, 3} {
		if !results[val] {
			t.Errorf("Element %d was not returned by GetRandom", val)
		}
	}

	// Delete an element and ensure it's no longer returned
	s.Delete(2)
	results = make(map[int]bool)
	for i := 0; i < 100; i++ {
		val := s.GetRandom()
		if val == 2 {
			t.Errorf("Deleted element 2 was returned by GetRandom")
		}
		results[val] = true
	}

	// Verify remaining elements
	for _, val := range []int{1, 3} {
		if !results[val] {
			t.Errorf("Element %d was not returned by GetRandom after deletion", val)
		}
	}
}

package snapset

import (
	"math/rand"
	"time"
)

// SnapSet defines the interface for a generic set data structure that supports basic set operations.
type SnapSet[T comparable] interface {
	// Insert adds an element to the set and returns its index.
	Insert(T) int

	// Delete removes the specified element from the set.
	// It returns the index of the deleted element and a boolean indicating if the deletion was successful.
	Delete(T) (int, bool)

	// Exists checks if the specified element is present in the set.
	Exists(T) bool

	// GetRandom returns a random element from the set.
	GetRandom() T
}

// DefaultBucketSize is the default initial size of the internal bucket map.
const DefaultBucketSize = 1 << 5

// Set is a generic set implementation that uses a map and a slice to store elements.
// The map (bucket) maps elements to their indices in the slice (list).
// The slice stores the elements and allows for efficient random access.
type Set[T comparable] struct {
	bucket  map[T]int  // maps elements to their indices in the list
	list    []T        // stores the elements
	currIdx int        // current index (index of the last inserted element)
	rand    *rand.Rand // random number generator for GetRandom
}

// New creates and returns a new instance of Set with the specified initial size.
// It initializes the internal bucket map and random number generator.
func New[T comparable](size int) SnapSet[T] {
	return &Set[T]{
		bucket: make(map[T]int, size),
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Insert adds the specified element to the set.
// It appends the element to the list, updates the bucket map with the new index,
// and updates the current index.
// It returns the index of the inserted element.
func (s *Set[T]) Insert(data T) int {
	s.list = append(s.list, data)
	s.currIdx = len(s.list) - 1
	s.bucket[data] = s.currIdx
	return s.currIdx
}

// Delete removes the specified element from the set.
// If the element exists, it swaps the element with the last element in the list,
// updates the bucket map accordingly, removes the last element from the list,
// and deletes the element from the bucket map.
// It returns the index of the deleted element and true if deletion was successful.
// If the element does not exist, it returns 0 and false.
func (s *Set[T]) Delete(element T) (int, bool) {
	idx, ok := s.bucket[element]
	if !ok {
		return 0, false // Element does not exist
	}

	lastIdx := len(s.list) - 1

	// Swap the element with the last element in the list
	lastElement := s.list[lastIdx]
	s.list[idx], s.list[lastIdx] = s.list[lastIdx], s.list[idx]

	// Update the bucket map with the new index of the last element
	s.bucket[lastElement] = idx

	// Remove the last element from the list
	s.list = s.list[:lastIdx]

	// Delete the element from the bucket map
	delete(s.bucket, element)

	// Update the current index
	s.currIdx = len(s.list) - 1

	return idx, true
}

// Exists checks whether the specified element exists in the set.
// It returns true if the element is found, otherwise false.
func (s *Set[T]) Exists(element T) bool {
	_, ok := s.bucket[element]
	return ok
}

// GetRandom returns a random element from the set.
// It generates a random index within the range of the list and returns the element at that index.
// Note: This method is not safe for concurrent use.
func (s *Set[T]) GetRandom() T {
	// Generate a random index using the random number generator
	rIdx := s.rand.Intn(len(s.list))
	return s.list[rIdx]
}

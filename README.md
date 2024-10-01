# SnapSet

SnapSet is a generic, fast, and efficient set data structure implemented in Go. It provides constant-time (`O(1)`) complexity for core operations such as insertion, deletion, existence checking, and random element retrieval. SnapSet is designed to offer high performance for applications that require frequent and fast set operations.

## Features

- **Generic Implementation**: Supports any comparable type, leveraging Go's type parameters.
- **Constant-Time Operations**: Achieves `O(1)` complexity for insert, delete, exists, and get random operations.
- **Random Element Retrieval**: Efficiently retrieves a random element from the set.
- **Minimalistic API**: Provides a simple and intuitive interface for ease of use.

## Installation

To use SnapSet in your project, you can simply copy the `snapset` package files into your project directory or import it if it's hosted in a repository.

```bash
go get github.com/yourusername/snapset
```

*Note: Replace `github.com/yourusername/snapset` with the actual import path of the package.*

## Usage

Here's how to use SnapSet in your Go project:

### Import the Package

```go
import "github.com/yourusername/snapset"
```

### Create a New Set

```go
// Create a new set of integers
s := snapset.New[int](initialSize)
```

### Insert Elements

```go
s.Insert(10)
s.Insert(20)
s.Insert(30)
```

### Check Existence

```go
exists := s.Exists(20) // returns true
```

### Delete Elements

```go
idx, ok := s.Delete(20)
if ok {
    fmt.Printf("Deleted element at index %d\n", idx)
}
```

### Get a Random Element

```go
randomElement := s.GetRandom()
fmt.Printf("Random Element: %v\n", randomElement)
```

### Full Example

```go
package main

import (
    "fmt"
    "github.com/yourusername/snapset"
)

func main() {
    // Create a new set
    s := snapset.New[string](10)

    // Insert elements
    s.Insert("apple")
    s.Insert("banana")
    s.Insert("cherry")

    // Check existence
    if s.Exists("banana") {
        fmt.Println("Banana exists in the set.")
    }

    // Get a random element
    fmt.Printf("Random fruit: %s\n", s.GetRandom())

    // Delete an element
    _, ok := s.Delete("banana")
    if ok {
        fmt.Println("Banana has been deleted.")
    }

    // Verify deletion
    if !s.Exists("banana") {
        fmt.Println("Banana no longer exists in the set.")
    }
}
```

## API Documentation

### Type Definitions

```go
type SnapSet[T comparable] interface {
    Insert(T) int
    Delete(T) (int, bool)
    Exists(T) bool
    GetRandom() T
}
```

### Functions

- `func New[T comparable](size int) SnapSet[T]`

  Creates and returns a new instance of SnapSet with the specified initial size.

### Methods

- `Insert(data T) int`

  Adds an element to the set. Returns the index of the inserted element.

- `Delete(element T) (int, bool)`

  Removes an element from the set. Returns the index of the deleted element and a boolean indicating success.

- `Exists(element T) bool`

  Checks if an element exists in the set.

- `GetRandom() T`

  Retrieves a random element from the set.

## Performance

SnapSet is designed for high performance with the following characteristics:

- **Insertion**: `O(1)` – Elements are appended to a slice, and their indices are stored in a map.
- **Deletion**: `O(1)` – Elements are swapped with the last element for efficient removal.
- **Existence Check**: `O(1)` – Uses a map to check for the existence of elements.
- **Random Access**: `O(1)` – Retrieves elements using a randomly generated index.

## Concurrency

**Note**: The current implementation of SnapSet is **not safe for concurrent use**. If you need to use SnapSet in a concurrent environment, consider adding synchronization mechanisms like mutexes to protect shared data.

Example with synchronization:

```go
type ThreadSafeSet[T comparable] struct {
    set snapset.SnapSet[T]
    mu  sync.Mutex
}

func (ts *ThreadSafeSet[T]) Insert(data T) int {
    ts.mu.Lock()
    defer ts.mu.Unlock()
    return ts.set.Insert(data)
}

// Implement other methods similarly...
```

## Limitations

- **Comparable Types**: Only types that are comparable can be used with SnapSet due to Go's type parameter constraints.
- **Duplicate Elements**: By default, SnapSet does not prevent the insertion of duplicate elements. If duplicates are undesirable, you should modify the `Insert` method to check for existing elements.
- **Empty Set Random Retrieval**: Calling `GetRandom` on an empty set will cause a runtime panic. Ensure the set is not empty before calling this method.

## Future Improvements

- **Thread Safety**: Implement built-in synchronization to make SnapSet safe for concurrent use.
- **Duplicate Handling**: Add optional duplicate prevention.
- **Error Handling**: Provide better error handling for edge cases like retrieving from an empty set.

## Acknowledgments

- Inspired by the need for efficient set operations in performance-critical applications.
- Utilizes Go's powerful generic types and standard library packages.

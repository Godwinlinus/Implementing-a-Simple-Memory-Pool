package main

import (
	"fmt"
	"sync"
)

// Define a simple struct
type MyObject struct {
	ID int
	Data string
}

// ObjectPool manages a pool of MyObject instances.
type ObjectPool struct {
	pool sync.Pool
}

// NewObjectPool creates a new ObjectPool with a given initial capacity.
func NewObjectPool(initialCapacity int) *ObjectPool {
	return &ObjectPool{
		pool: sync.Pool{
			New: func() interface{} {
				return &MyObject{} // Initialize a new object
			},
		},
	}
}

// Get retrieves an object from the pool.
func (op *ObjectPool) Get() *MyObject {
	return op.pool.Get().(*MyObject) // Type assertion to *MyObject
}

// Put returns an object to the pool.
func (op *ObjectPool) Put(obj *MyObject) {
	// Reset the object's fields before returning it to the pool
	obj.ID = 0
	obj.Data = ""
	op.pool.Put(obj)
}

func main() {
	// Create a new object pool with an initial capacity
	pool := NewObjectPool(10)

	// Get an object from the pool
	obj1 := pool.Get()
	obj1.ID = 1
	obj1.Data = "Object 1"
	fmt.Printf("Object 1: %+v\n", obj1)

	// Get another object from the pool
	obj2 := pool.Get()
	obj2.ID = 2
	obj2.Data = "Object 2"
	fmt.Printf("Object 2: %+v\n", obj2)

	// Return the objects to the pool
	pool.Put(obj1)
	pool.Put(obj2)

	// Get an object again to see if it's reused
	obj3 := pool.Get()
	fmt.Printf("Object 3 (reused): %+v\n", obj3) // Notice the fields are reset
}
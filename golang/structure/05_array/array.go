package _5_array

import "errors"

type Array struct {
	data     []int
	length   uint
	capacity uint
}

func NewArray(capacity uint) (*Array, error) {
	if capacity == 0 {
		return nil, errors.New("the capacity is 0")
	}

	return &Array{
		data:     make([]int, capacity, capacity),
		length:   0,
		capacity: capacity,
	}, nil
}

func (a *Array) len() uint {
	return a.length
}

func (a *Array) cap() uint {
	return a.capacity
}

func (a *Array) isOutOfCap(idx uint) bool {
	return idx >= a.cap()
}

// O(1)
func (a *Array) FindByIdx(idx uint) (int, error) {
	if a.isOutOfCap(idx) {
		return 0, errors.New("out of range")
	}

	return a.data[idx], nil
}

// O(N)
func (a *Array) Insert(idx uint, val int) error {
	if a.len() >= a.cap() {
		return errors.New("array is full")
	}

	if a.isOutOfCap(idx) {
		return errors.New("out of range")
	}

	for i := a.len(); i > idx; i-- {
		a.data[i] = a.data[i-1]
	}

	a.data[idx] = val
	a.length += 1
	return nil
}

func (a *Array) InsertTail(val int) error {
	return a.Insert(a.len(), val)
}

// O(N)
func (a *Array) Delete(idx uint) error {
	if a.isOutOfCap(idx) {
		return errors.New("out of range")
	}

	if idx >= a.len() {
		return nil
	}

	for i := idx; i < a.len(); i++ {
		a.data[i] = a.data[i+1]
	}

	a.length -= 1
	return nil
}
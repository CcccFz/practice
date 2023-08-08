package _5_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := assert.New(t)

	arr, err := NewArray(10)
	a.Nil(err)
	a.NotNil(arr)

	arr, err = NewArray(0)
	a.NotNil(err)
	a.Nil(arr)
}

func TestInsert(t *testing.T) {
	a := assert.New(t)

	capacity := 10
	arr, err := NewArray(uint(capacity))
	a.Nil(err)

	for i := 0; i < capacity - 2; i++ {
		err = arr.Insert(uint(i), i + 1)
		a.Nil(err)
	}
	a.EqualValues(Array{
		data:     []int{1, 2, 3, 4, 5, 6, 7, 8, 0, 0},
		length:   8,
		capacity: 10,
	}, *arr)

	err = arr.Insert(11, 200)
	a.NotNil(err)

	err = arr.InsertTail(9)
	a.Nil(err)
	a.EqualValues(Array{
		data:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		length:   9,
		capacity: 10,
	}, *arr)

	err = arr.Insert(5, 100)
	a.Nil(err)
	a.EqualValues(Array{
		data:     []int{1, 2, 3, 4, 5, 100, 6, 7, 8, 9},
		length:   10,
		capacity: 10,
	}, *arr)

	err = arr.Insert(5, 200)
	a.NotNil(err)

	err = arr.InsertTail(200)
	a.NotNil(err)

	a.EqualValues(Array{
		data:     []int{1, 2, 3, 4, 5, 100, 6, 7, 8, 9},
		length:   10,
		capacity: 10,
	}, *arr)
}

func TestFind(t *testing.T) {
	a := assert.New(t)

	arr := &Array{
		data:     []int{1, 2, 3, 4, 5, 6, 7, 8, 0, 0},
		length:   8,
		capacity: 10,
	}

	ret, err := arr.FindByIdx(5)
	a.Nil(err)
	a.Equal(6, ret)

	ret, err = arr.FindByIdx(11)
	a.NotNil(err)
	a.Equal(0, ret)
}


func TestDelete(t *testing.T) {
	a := assert.New(t)

	arr := &Array{
		data:     []int{1, 2, 3, 4, 5, 6, 7, 8, 0, 0},
		length:   8,
		capacity: 10,
	}

	err := arr.Delete(7)
	a.Nil(err)
	a.EqualValues(Array{
		data:     []int{1, 2, 3, 4, 5, 6, 7, 0, 0, 0},
		length:   7,
		capacity: 10,
	}, *arr)

	err = arr.Delete(0)
	a.Nil(err)
	a.EqualValues(Array{
		data:     []int{2, 3, 4, 5, 6, 7, 0, 0, 0, 0},
		length:   6,
		capacity: 10,
	}, *arr)

	err = arr.Delete(9)
	a.Nil(err)
	a.EqualValues(Array{
		data:     []int{2, 3, 4, 5, 6, 7, 0, 0, 0, 0},
		length:   6,
		capacity: 10,
	}, *arr)
}

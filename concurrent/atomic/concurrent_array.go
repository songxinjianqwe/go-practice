package main

import (
	"errors"
	"fmt"
	"sync/atomic"
)

type ConcurrentArray interface {
	Set(index uint32, element int) error
	Get(index uint32) (int, error)
	Len() uint32
}

type ConcurrentArrayImpl struct {
	length uint32
	val atomic.Value
}

func (this *ConcurrentArrayImpl) Set(index uint32, element int) error {
	if err := this.checkIndex(index); err != nil {
		return err
	}
	if err := this.checkValue(); err != nil {
		return err
	}
	newArray := make([]int, this.length)
	copy(newArray, this.val.Load().([]int))
	newArray[index] = element
	this.val.Store(newArray)
	return nil
}

func (this *ConcurrentArrayImpl) Get(index uint32) (elem int, err error) {
	if err = this.checkIndex(index); err != nil {
		return 0, err
	}
	if err = this.checkValue(); err != nil {
		return 0, err
	}
	elem = this.val.Load().([]int)[index]
	return elem, err
}


func (this *ConcurrentArrayImpl) Len() uint32 {
	return this.length
}

func (this *ConcurrentArrayImpl) checkIndex(index uint32) error {
	if index >= this.length {
		return fmt.Errorf("index out of bounds: %d", index)
	}
	return nil
}

func (this *ConcurrentArrayImpl) checkValue() error {

}



func NewConcurrentArray(length uint32) ConcurrentArray {
	array := ConcurrentArrayImpl{}
	array.length = length
	array.val.Store(make([]int, array.length))
	return array
}


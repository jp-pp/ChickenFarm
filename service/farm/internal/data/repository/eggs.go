package repository

import (
	"math/bits"
	"sync"
)

type fridge struct {
	mutex 	*sync.RWMutex
	eggs 	uint
}

func NewFridge() *fridge {
	return &fridge{mutex: &sync.RWMutex{}}
}

func (f *fridge) PutAnEgg(egg uint) (err error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	if f.eggs != bits.UintSize {
		f.eggs += egg
	}
	return
}

func (f *fridge) TakeEggs(eggs uint) (err error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	if f.eggs >= eggs {
		f.eggs -= eggs
	} else {
		f.eggs = 0
	}
	return
}

func (f *fridge) CountEggs() (eggs uint) {
	f.mutex.RLock()
	defer f.mutex.RUnlock()

	return f.eggs
}

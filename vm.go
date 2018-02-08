package lru_clock

import "sync"

const memSize = 64

type block struct {
	refCount int
	active int
}

type vm struct {
	blocks *[memSize]block
	memLock *sync.Mutex
}

func initMem() *vm {
	emptyMem := new(vm)
	emptyMem.blocks = new([memSize]block)
	emptyMem.memLock = &sync.Mutex{}
	for i := 0; i < memSize; i++ {
		emptyMem.blocks[i].active = -1
	}
	return emptyMem
}

func getBlock(vm *vm) int {
	for i := 0; i < memSize; i++ {
		if vm.blocks[i].active < 0 {
			return i
		}
	}
	return -1
}

func writeBlock(writer *process) int {
	for i := writer.memOffset; i < maxPMem; i++ {
		writer.fs.blocks[i].refCount += 1
		writer.fs.blocks[i].active = writer.pid
	}
	return maxPMem
}

func freeBlock(writer *process) int {
	for i := writer.memOffset; i < maxPMem; i++ {
		writer.fs.blocks[i].active = -1
	}
	return maxPMem
}

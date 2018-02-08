package lru_clock

import "time"

const maxPMem = 16

type process struct {
	pid int
	active bool
	load float32
	memOffset int
	fs *vm
}

func createProcess(id int, vm *vm) *process {
	process := new(process)
	process.load = 0.0
	process.active = false
	process.pid = id
	process.memOffset = -1
	process.fs = vm
	return process
}

func run(process *process)  {
	process.fs.memLock.Lock()
	process.memOffset = getBlock(process.fs)
	writeBlock(process)
	process.fs.memLock.Unlock()

	timer := time.NewTimer(time.Second * 5)
	<-timer.C
	kill(process)
}

func kill(process *process)  {
	process.fs.memLock.Lock()
	freeBlock(process)
	process.fs.memLock.Unlock()
}
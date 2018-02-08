package main

func main() {
	memory := initMem()
	processPool := new([256] *process)

	for i := 0; i < len(processPool); i++ {
		processPool[i] = createProcess(i, memory)
		run(processPool[i])
	}
}

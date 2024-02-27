package queue

type Queue struct {
	Queue []string
}

type CommandsQueue struct {
	Store map[string]*Queue
}

func CreateNewCommandsQueue() *CommandsQueue {
	return &CommandsQueue{
		Store: make(map[string]*Queue),
	}
}

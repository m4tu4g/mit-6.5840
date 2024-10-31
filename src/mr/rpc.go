package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import (
	"os"
	"strconv"
)

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

// Add your RPC definitions here.

const (
	ShutDownAction = "shut down"
	GetTaskAction  = "get task"
)

type GetMapTaskArgs struct {
}

type GetMapTaskReply struct {
	Filename  string
	NReduce   int // nReduce, making in upper for public access
	MapTaskID int
}

type InformMapTaskResultArgs struct {
	TaskID                int
	IntermediateFileNames []string
}

type InformMapTaskResultReply struct {
	Action string
}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/5840-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}

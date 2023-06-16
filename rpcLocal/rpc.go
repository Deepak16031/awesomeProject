package rpcLocal

import (
	"os"
	"strconv"
)

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

func CoordinatorSock() string {
	s := "/var/tmp/824-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}

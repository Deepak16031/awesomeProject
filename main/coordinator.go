package main

import (
	"awesomeProject/rpcLocal"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"
)

func main() {
	m := MakeCoordinator()
	for m.Done() == false {
		time.Sleep(time.Second)
	}

	time.Sleep(time.Second)
}

type Coordinator struct {
	// Your definitions here.

}

func (c *Coordinator) Example(args *rpcLocal.ExampleArgs, reply *rpcLocal.ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

// start a thread that listens for RPCs from worker.go
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := rpcLocal.CoordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

func MakeCoordinator() *Coordinator {
	c := Coordinator{}

	// Your code here.

	c.server()
	return &c
}
func (c *Coordinator) Done() bool {
	ret := false

	// Your code here.
	return ret
}

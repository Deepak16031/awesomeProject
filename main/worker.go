package main

import (
	"awesomeProject/rpcLocal"
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	// Your worker implementation here.

	// uncomment to send the Example RPC to the coordinator.
	CallExample()

}

func CallExample() {

	// declare an argument structure.
	args := rpcLocal.ExampleArgs{}

	// fill in the argument(s).
	args.X = 99

	// declare a reply structure.
	reply := rpcLocal.ExampleReply{}

	// send the RPC request, wait for the reply.
	call("Coordinator.Example", &args, &reply)

	// reply.Y should be 100.
	fmt.Printf("reply.Y %v\n", reply.Y)
}

// send an RPC request to the coordinator, wait for the response.
// usually returns true.
// returns false if something goes wrong.
func call(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpcLocal.DialHTTP("tcp", "127.0.0.1"+":1234")
	sockname := rpcLocal.CoordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}

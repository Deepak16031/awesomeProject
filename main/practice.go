package main

import (
	"awesomeProject/rpcLocal"
	"os"
)

func main() {
	println(rpcLocal.CoordinatorSock())
	print(os.Getuid())
}

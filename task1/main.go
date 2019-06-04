package main

import (
	"awesomeProject/task1/client"
	"awesomeProject/task1/server"
	"fmt"
)
func serv() {
fmt.Println("hi")
server.Start("50051")
}
func main() {
	go serv()

 //time.Sleep(2 * time.Second)
	client.Start("50051")
}

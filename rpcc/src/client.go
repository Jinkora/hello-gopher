package main

import (
	"net/rpc"
	"log"
	"fmt"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func main() {

	c, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Args{7, 8}

	var reply int
	err = c.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	fmt.Println("---------------------------")
	quotient := new(Quotient)
	<-c.Go("Arith.Divide", args, &quotient, nil).Done
	fmt.Println("Quo:", quotient.Quo)
	fmt.Println("Rem:", quotient.Rem)

	fmt.Println("---------------------------")
	//func Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call
	replyToThis := 0
	<-(c.Go("Arith.TwoTimes", 4, &replyToThis, nil)).Done
	fmt.Println("replyToThis=", replyToThis)

}

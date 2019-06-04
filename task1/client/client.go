package client

import (
	"awesomeProject/task1/ayushpb/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func Start(port string) {
	cc, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	ctx := context.Background()
	client := ayushpb.NewCalculatorServiceClient(cc)
	// UNARY
	fmt.Println("Task1")
	var n int32 = 10
	restables, err := client.Tables(ctx, &ayushpb.TableRequest{
		Number: n,
	})

	fmt.Println("Cube of ", n, "is", restables.GetResult())

	// SERVER STREAMING
	fmt.Println()
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")
	fmt.Println()
	fmt.Println("Greeting of the day")
		var str="ayush"
	stream, err := client.Greet(ctx, &ayushpb.GreetingsRequest{
		A:str,
	})
	if err != nil {
		log.Fatal(err)
	}

		resPF:= stream.GetB()
		fmt.Println(resPF)

	// CLIENT STREAMING
	//fmt.Println("Average")
	//
	//streamAvg, err := client.Average(ctx)
	//for _, num := range []int32{1,2,3,4,5,6,7,8} {
	//	streamAvg.Send(&ayushpb.CalculatorRequest{
	//		Number:num,
	//	})
	//}
	//
	//respAvg, err := streamAvg.CloseAndRecv()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Average is ", respAvg.GetResult())

	done := make(chan bool)

	biDiStream, err := client.GCD(ctx)

	go func() {
		for _, num := range []int32{1,2,3,4,5,6,7,8} {
			biDiStream.Send(&ayushpb.TableRequest{
				Number:num,
			})
		}
		biDiStream.CloseSend()
	}()

	go func() {
		fmt.Println()
		fmt.Println("-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")
		fmt.Println()
		fmt.Println("GCD OF THE GIVEN ARRAY")
		for {
			resOE, err := biDiStream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("GCD Result: ", resOE.GetNum())
			}
		}
		done <- true
	}()

	<- done
}
